// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	mock "github.com/stretchr/testify/mock"
)

// Traceable is an autogenerated mock type for the Traceable type
type Traceable struct {
	mock.Mock
}

// SetTraceID provides a mock function with given fields: ctx, traceID
func (_m *Traceable) SetTraceID(ctx context.Context, traceID string) context.Context {
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
func (_m *Traceable) TraceID(ctx context.Context) string {
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

// NewTraceable creates a new instance of Traceable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTraceable(t interface {
	mock.TestingT
	Cleanup(func())
}) *Traceable {
	mock := &Traceable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
