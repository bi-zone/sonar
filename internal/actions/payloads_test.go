package actions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bi-zone/sonar/internal/actions"
	"github.com/bi-zone/sonar/internal/utils/errors"
)

func TestCreatePayload_Succes(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.CreatePayloadParams{
		Name: "test",
	}

	r, err := acts.CreatePayload(u, p)
	assert.NoError(t, err)

	assert.NotNil(t, r)
	assert.Equal(t, "test", r.Name)
	assert.Equal(t, u.ID, r.UserID)
}

func TestCreatePayload_Validation(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.CreatePayloadParams{
		Name: "",
	}

	_, err = acts.CreatePayload(u, p)
	assert.Error(t, err)
	assert.IsType(t, &errors.ValidationError{}, err)
}

func TestCreatePayload_Conflict(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.CreatePayloadParams{
		Name: "payload1",
	}

	_, err = acts.CreatePayload(u, p)
	assert.Error(t, err)
	assert.IsType(t, &errors.ConflictError{}, err)
}

func TestDeletePayload_Succes(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.DeletePayloadParams{
		Name: "payload1",
	}

	_, err = acts.DeletePayload(u, p)
	assert.NoError(t, err)
}

func TestDeletePayload_Validation(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.DeletePayloadParams{
		Name: "",
	}

	_, err = acts.DeletePayload(u, p)
	assert.Error(t, err)
	assert.IsType(t, &errors.ValidationError{}, err)
}

func TestDeletePayload_NotFound(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(1)
	require.NoError(t, err)

	p := actions.DeletePayloadParams{
		Name: "not-exist",
	}

	_, err = acts.DeletePayload(u, p)
	assert.Error(t, err)
	assert.IsType(t, &errors.NotFoundError{}, err)
}

func TestListPayloads_Succes(t *testing.T) {
	setup(t)
	defer teardown(t)

	u, err := db.UsersGetByID(2)
	require.NoError(t, err)

	p := actions.ListPayloadsParams{
		Name: "",
	}

	r, err := acts.ListPayloads(u, p)
	assert.NoError(t, err)
	assert.Len(t, r, 2)

	p = actions.ListPayloadsParams{
		Name: "payload2",
	}

	r, err = acts.ListPayloads(u, p)
	assert.NoError(t, err)
	assert.Len(t, r, 1)
}
