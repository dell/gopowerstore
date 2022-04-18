// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import api "github.com/dell/gopowerstore/api"
import context "context"
import http "net/http"
import mock "github.com/stretchr/testify/mock"

// ApiClient is an autogenerated mock type for the ApiClient type
type ApiClient struct {
	mock.Mock
}

// Query provides a mock function with given fields: ctx, cfg, resp
func (_m *ApiClient) Query(ctx context.Context, cfg api.RequestConfigRenderer, resp interface{}) (api.RespMeta, error) {
	ret := _m.Called(ctx, cfg, resp)

	var r0 api.RespMeta
	if rf, ok := ret.Get(0).(func(context.Context, api.RequestConfigRenderer, interface{}) api.RespMeta); ok {
		r0 = rf(ctx, cfg, resp)
	} else {
		r0 = ret.Get(0).(api.RespMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, api.RequestConfigRenderer, interface{}) error); ok {
		r1 = rf(ctx, cfg, resp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryParams provides a mock function with given fields:
func (_m *ApiClient) QueryParams() api.QueryParamsEncoder {
	ret := _m.Called()

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
func (_m *ApiClient) QueryParamsWithFields(provider api.FieldProvider) api.QueryParamsEncoder {
	ret := _m.Called(provider)

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
func (_m *ApiClient) SetCustomHTTPHeaders(headers http.Header) {
	_m.Called(headers)
}

// SetLogger provides a mock function with given fields: logger
func (_m *ApiClient) SetLogger(logger api.Logger) {
	_m.Called(logger)
}

// SetTraceID provides a mock function with given fields: ctx, traceID
func (_m *ApiClient) SetTraceID(ctx context.Context, traceID string) context.Context {
	ret := _m.Called(ctx, traceID)

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
func (_m *ApiClient) TraceID(ctx context.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
