// Code generated by mockery v1.0.0. DO NOT EDIT.

package actions_mock

import (
	actions "github.com/bi-zone/sonar/internal/actions"
	errors "github.com/bi-zone/sonar/internal/utils/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/bi-zone/sonar/internal/models"
)

// Actions is an autogenerated mock type for the Actions type
type Actions struct {
	mock.Mock
}

// CreatePayload provides a mock function with given fields: _a0, _a1
func (_m *Actions) CreatePayload(_a0 *models.User, _a1 actions.CreatePayloadParams) (*models.Payload, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Payload
	if rf, ok := ret.Get(0).(func(*models.User, actions.CreatePayloadParams) *models.Payload); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Payload)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(*models.User, actions.CreatePayloadParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *Actions) CreateUser(_a0 *models.User, _a1 actions.CreateUserParams) (*models.User, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(*models.User, actions.CreateUserParams) *models.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(*models.User, actions.CreateUserParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// DeletePayload provides a mock function with given fields: _a0, _a1
func (_m *Actions) DeletePayload(_a0 *models.User, _a1 actions.DeletePayloadParams) (*actions.MessageResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 *actions.MessageResult
	if rf, ok := ret.Get(0).(func(*models.User, actions.DeletePayloadParams) *actions.MessageResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.MessageResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(*models.User, actions.DeletePayloadParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: _a0, _a1
func (_m *Actions) DeleteUser(_a0 *models.User, _a1 actions.DeleteUserParams) (*actions.MessageResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 *actions.MessageResult
	if rf, ok := ret.Get(0).(func(*models.User, actions.DeleteUserParams) *actions.MessageResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*actions.MessageResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(*models.User, actions.DeleteUserParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// ListPayloads provides a mock function with given fields: _a0, _a1
func (_m *Actions) ListPayloads(_a0 *models.User, _a1 actions.ListPayloadsParams) ([]*models.Payload, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.Payload
	if rf, ok := ret.Get(0).(func(*models.User, actions.ListPayloadsParams) []*models.Payload); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Payload)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(*models.User, actions.ListPayloadsParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}
