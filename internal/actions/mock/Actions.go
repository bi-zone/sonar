// Code generated by mockery v2.6.0. DO NOT EDIT.

package actions_mock

import (
	context "context"

	actions "github.com/bi-zone/sonar/internal/actions"

	errors "github.com/bi-zone/sonar/internal/utils/errors"

	mock "github.com/stretchr/testify/mock"
)

// Actions is an autogenerated mock type for the Actions type
type Actions struct {
	mock.Mock
}

// DNSRecordsCreate provides a mock function with given fields: _a0, _a1
func (_m *Actions) DNSRecordsCreate(_a0 context.Context, _a1 actions.DNSRecordsCreateParams) (actions.DNSRecordsCreateResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.DNSRecordsCreateResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.DNSRecordsCreateParams) actions.DNSRecordsCreateResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.DNSRecordsCreateResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.DNSRecordsCreateParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// DNSRecordsDelete provides a mock function with given fields: _a0, _a1
func (_m *Actions) DNSRecordsDelete(_a0 context.Context, _a1 actions.DNSRecordsDeleteParams) (actions.DNSRecordsDeleteResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.DNSRecordsDeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.DNSRecordsDeleteParams) actions.DNSRecordsDeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.DNSRecordsDeleteResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.DNSRecordsDeleteParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// DNSRecordsList provides a mock function with given fields: _a0, _a1
func (_m *Actions) DNSRecordsList(_a0 context.Context, _a1 actions.DNSRecordsListParams) (actions.DNSRecordsListResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.DNSRecordsListResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.DNSRecordsListParams) actions.DNSRecordsListResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.DNSRecordsListResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.DNSRecordsListParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// EventsGet provides a mock function with given fields: _a0, _a1
func (_m *Actions) EventsGet(_a0 context.Context, _a1 actions.EventsGetParams) (actions.EventsGetResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.EventsGetResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.EventsGetParams) actions.EventsGetResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.EventsGetResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.EventsGetParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// EventsList provides a mock function with given fields: _a0, _a1
func (_m *Actions) EventsList(_a0 context.Context, _a1 actions.EventsListParams) (actions.EventsListResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.EventsListResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.EventsListParams) actions.EventsListResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.EventsListResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.EventsListParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// HTTPRoutesCreate provides a mock function with given fields: _a0, _a1
func (_m *Actions) HTTPRoutesCreate(_a0 context.Context, _a1 actions.HTTPRoutesCreateParams) (actions.HTTPRoutesCreateResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.HTTPRoutesCreateResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.HTTPRoutesCreateParams) actions.HTTPRoutesCreateResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.HTTPRoutesCreateResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.HTTPRoutesCreateParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// HTTPRoutesDelete provides a mock function with given fields: _a0, _a1
func (_m *Actions) HTTPRoutesDelete(_a0 context.Context, _a1 actions.HTTPRoutesDeleteParams) (actions.HTTPRoutesDeleteResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.HTTPRoutesDeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.HTTPRoutesDeleteParams) actions.HTTPRoutesDeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.HTTPRoutesDeleteResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.HTTPRoutesDeleteParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// HTTPRoutesList provides a mock function with given fields: _a0, _a1
func (_m *Actions) HTTPRoutesList(_a0 context.Context, _a1 actions.HTTPRoutesListParams) (actions.HTTPRoutesListResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.HTTPRoutesListResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.HTTPRoutesListParams) actions.HTTPRoutesListResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.HTTPRoutesListResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.HTTPRoutesListParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// PayloadsCreate provides a mock function with given fields: _a0, _a1
func (_m *Actions) PayloadsCreate(_a0 context.Context, _a1 actions.PayloadsCreateParams) (actions.PayloadsCreateResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.PayloadsCreateResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.PayloadsCreateParams) actions.PayloadsCreateResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.PayloadsCreateResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.PayloadsCreateParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// PayloadsDelete provides a mock function with given fields: _a0, _a1
func (_m *Actions) PayloadsDelete(_a0 context.Context, _a1 actions.PayloadsDeleteParams) (actions.PayloadsDeleteResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.PayloadsDeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.PayloadsDeleteParams) actions.PayloadsDeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.PayloadsDeleteResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.PayloadsDeleteParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// PayloadsList provides a mock function with given fields: _a0, _a1
func (_m *Actions) PayloadsList(_a0 context.Context, _a1 actions.PayloadsListParams) (actions.PayloadsListResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.PayloadsListResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.PayloadsListParams) actions.PayloadsListResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.PayloadsListResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.PayloadsListParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// PayloadsUpdate provides a mock function with given fields: _a0, _a1
func (_m *Actions) PayloadsUpdate(_a0 context.Context, _a1 actions.PayloadsUpdateParams) (actions.PayloadsUpdateResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.PayloadsUpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.PayloadsUpdateParams) actions.PayloadsUpdateResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.PayloadsUpdateResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.PayloadsUpdateParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// UserCurrent provides a mock function with given fields: _a0
func (_m *Actions) UserCurrent(_a0 context.Context) (actions.UserCurrentResult, errors.Error) {
	ret := _m.Called(_a0)

	var r0 actions.UserCurrentResult
	if rf, ok := ret.Get(0).(func(context.Context) actions.UserCurrentResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.UserCurrentResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context) errors.Error); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// UsersCreate provides a mock function with given fields: _a0, _a1
func (_m *Actions) UsersCreate(_a0 context.Context, _a1 actions.UsersCreateParams) (actions.UsersCreateResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.UsersCreateResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.UsersCreateParams) actions.UsersCreateResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.UsersCreateResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.UsersCreateParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}

// UsersDelete provides a mock function with given fields: _a0, _a1
func (_m *Actions) UsersDelete(_a0 context.Context, _a1 actions.UsersDeleteParams) (actions.UsersDeleteResult, errors.Error) {
	ret := _m.Called(_a0, _a1)

	var r0 actions.UsersDeleteResult
	if rf, ok := ret.Get(0).(func(context.Context, actions.UsersDeleteParams) actions.UsersDeleteResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actions.UsersDeleteResult)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, actions.UsersDeleteParams) errors.Error); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}
