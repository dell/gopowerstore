// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/dell/gopowerstore/api"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// APIClient is an autogenerated mock type for the APIClient type
type APIClient struct {
	mock.Mock
}

// GetCustomHTTPHeaders provides a mock function with given fields:
func (_m *APIClient) GetCustomHTTPHeaders() http.Header {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCustomHTTPHeaders")
	}

	var r0 http.Header
	if rf, ok := ret.Get(0).(func() http.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Header)
		}
	}

	return r0
}

// Query provides a mock function with given fields: ctx, cfg, resp
func (_m *APIClient) Query(ctx context.Context, cfg api.RequestConfigRenderer, resp interface{}) (api.RespMeta, error) {
	ret := _m.Called(ctx, cfg, resp)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 api.RespMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, api.RequestConfigRenderer, interface{}) (api.RespMeta, error)); ok {
		return rf(ctx, cfg, resp)
	}
	if rf, ok := ret.Get(0).(func(context.Context, api.RequestConfigRenderer, interface{}) api.RespMeta); ok {
		r0 = rf(ctx, cfg, resp)
	} else {
		r0 = ret.Get(0).(api.RespMeta)
	}

	if rf, ok := ret.Get(1).(func(context.Context, api.RequestConfigRenderer, interface{}) error); ok {
		r1 = rf(ctx, cfg, resp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryParams provides a mock function with given fields:
func (_m *APIClient) QueryParams() api.QueryParamsEncoder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for QueryParams")
	}

	var r0 api.QueryParamsEncoder
	if rf, ok := ret.Get(0).(func() api.QueryParamsEncoder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.QueryParamsEncoder)
		}
	}

	return r0
}

// QueryParamsWithFields provides a mock function with given fields: provider
func (_m *APIClient) QueryParamsWithFields(provider api.FieldProvider) api.QueryParamsEncoder {
	ret := _m.Called(provider)

	if len(ret) == 0 {
		panic("no return value specified for QueryParamsWithFields")
	}

	var r0 api.QueryParamsEncoder
	if rf, ok := ret.Get(0).(func(api.FieldProvider) api.QueryParamsEncoder); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.QueryParamsEncoder)
		}
	}

	return r0
}

// SetCustomHTTPHeaders provides a mock function with given fields: headers
func (_m *APIClient) SetCustomHTTPHeaders(headers http.Header) {
	_m.Called(headers)
}

// SetLogger provides a mock function with given fields: logger
func (_m *APIClient) SetLogger(logger api.Logger) {
	_m.Called(logger)
}

// SetTraceID provides a mock function with given fields: ctx, traceID
func (_m *APIClient) SetTraceID(ctx context.Context, traceID string) context.Context {
	ret := _m.Called(ctx, traceID)

	if len(ret) == 0 {
		panic("no return value specified for SetTraceID")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, traceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// TraceID provides a mock function with given fields: ctx
func (_m *APIClient) TraceID(ctx context.Context) string {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for TraceID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewAPIClient creates a new instance of APIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *APIClient {
	mock := &APIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
