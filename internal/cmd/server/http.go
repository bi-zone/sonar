package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bi-zone/sonar/internal/models"
	"github.com/bi-zone/sonar/internal/utils"
	"github.com/bi-zone/sonar/pkg/httpx"
)

const (
	httpHandlerTimeout = time.Second * 10
	httpMaxBodyBytes   = 1 << 20
)

func HTTPDefault(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	rnd, _ := utils.GenerateRandomString(8)
	w.Write([]byte(fmt.Sprintf("<html><body>%s</body></html>", rnd)))
}

func HTTPHandler(notify func(*httpx.Event)) http.Handler {
	return http.TimeoutHandler(
		httpx.MultipartHandler(
			httpx.MaxBytesHandler(
				httpx.NotifyHandler(notify, http.HandlerFunc(HTTPDefault)),
				httpMaxBodyBytes,
			),
			httpMaxBodyBytes,
		),
		httpHandlerTimeout,
		"timeout",
	)
}

func HTTPEvent(e *httpx.Event) *models.Event {
	var proto models.Proto

	if e.Secure {
		proto = models.ProtoHTTPS
	} else {
		proto = models.ProtoHTTP
	}

	return &models.Event{
		Protocol: proto,
		Log:      append(e.RawRequest[:], e.RawResponse...),
		Meta: map[string]interface{}{
			"tls": e.Secure,
		},
		RemoteAddr: e.RemoteAddr,
		ReceivedAt: e.ReceivedAt,
	}
}
