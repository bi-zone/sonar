package main

import (
	"database/sql"
	"encoding/json"
	"net"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/bi-zone/sonar/internal/database"
	"github.com/bi-zone/sonar/internal/database/dbactions"
	"github.com/bi-zone/sonar/internal/handlers"
	"github.com/bi-zone/sonar/internal/handlers/dnsx"
	"github.com/bi-zone/sonar/internal/models"
	"github.com/bi-zone/sonar/internal/modules"
	"github.com/bi-zone/sonar/internal/tls"
	"github.com/bi-zone/sonar/pkg/server/http"
	"github.com/bi-zone/sonar/pkg/server/smtp"
)

var (
	log *logrus.Logger
)

func main() {

	//
	// Logger
	//

	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	//
	// Config
	//

	var cfg Config

	if err := envconfig.Process("sonar", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	prettyCfg, _ := json.MarshalIndent(&cfg, "", "  ")
	log.Infof("Config:\n%+v", string(prettyCfg))

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	//
	// DB
	//

	db, err := database.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	// Create admin user
	admin := &models.User{
		Name:      "admin",
		IsAdmin:   true,
		CreatedBy: nil,
		Params: models.UserParams{
			TelegramID: cfg.Modules.Telegram.Admin,
			APIToken:   cfg.Modules.API.Admin,
		},
	}

	if u, err := db.UsersGetByName("admin"); err == sql.ErrNoRows {
		// There is no admin yet - create one
		if err := db.UsersCreate(admin); err != nil {
			log.Fatal(err)
		}
	} else if err == nil {
		// Admin user exists - update
		admin.ID = u.ID
		if err := db.UsersUpdate(admin); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

	//
	// Actions
	//

	actions := dbactions.New(db, log, cfg.Domain)

	//
	// Events
	//

	events := make(chan models.Event, 100)
	defer close(events)

	//
	// DNS
	//

	var dnsStarted sync.WaitGroup

	dnsStarted.Add(1)

	defaultRecords, err := dnsx.DefaultRecords(cfg.Domain, net.ParseIP(cfg.IP))
	if err != nil {
		log.Fatal(err)
	}

	dns := dnsx.New(":53", cfg.Domain,
		[]dnsx.Finder{
			dnsx.NewDatabaseFinder(db, cfg.Domain),
			defaultRecords,
		},
		dnsx.NotifyRequestFunc(handlers.NotifyRequestFunc(AddProtoEvent("DNS", events))),
		dnsx.NotifyStartedFunc(func() {
			dnsStarted.Done()
		}),
	)

	if err != nil {
		log.Fatal("Failed to create DNS handler: %w", err)
	}

	go func() {
		if err := dns.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start DNS handler: %v", err.Error())
		}
	}()

	//
	// TLS
	//

	// Wait for DNS server to start, because we need to
	// use it as DNS challenge provider for Let's Encrypt
	dnsStarted.Wait()

	tls, err := tls.New(&cfg.TLS, log, cfg.Domain, defaultRecords)
	if err != nil {
		log.Fatalf("Failed to init TLS: %v", err)
	}

	go func() {
		err := tls.Start()
		if err != nil {
			log.Fatalf("Failed to start TLS: %v", err)
		}
	}()

	tls.Wait()

	tlsConfig, err := tls.GetConfig()
	if err != nil {
		log.Fatalf("Failed to get TLS config: %v", err)
	}

	//
	// HTTP
	//

	go func() {
		srv := http.New(":80", http.NotifyRequestFunc(AddProtoEvent("HTTP", events)))

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed start HTTP handler: %s", err.Error())
		}

	}()

	//
	// HTTPS
	//

	go func() {
		srv := http.New(":443",
			http.NotifyRequestFunc(AddProtoEvent("HTTPS", events)),
			http.TLSConfig(tlsConfig))

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed start HTTPS handler: %s", err.Error())
		}
	}()

	//
	// SMTP
	//

	go func() {
		// Pass TLS config to be able to handle "STARTTLS" command.
		srv := smtp.New(":25", cfg.Domain,
			smtp.NotifyRequestFunc(AddProtoEvent("SMTP", events)),
			smtp.TLSConfig(tlsConfig),
			smtp.StartTLS(true))

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed start SMTP handler: %s", err.Error())
		}
	}()

	//
	// Modules
	//

	controllers, notifiers, err := modules.Init(&cfg.Modules, db, log, tlsConfig, actions, cfg.Domain)
	if err != nil {
		log.Fatal(err)
	}

	// Start controllers
	for _, c := range controllers {
		go func(c modules.Controller) {
			if err := c.Start(); err != nil {
				log.Fatal(err)
			}
		}(c)
	}

	// Process events
	go ProcessEvents(events, db, notifiers)

	// Wait forever
	select {}
}
