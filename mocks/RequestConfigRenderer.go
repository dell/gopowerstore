// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	api "github.com/dell/gopowerstore/api"
	mock "github.com/stretchr/testify/mock"
)

// RequestConfigRenderer is an autogenerated mock type for the RequestConfigRenderer type
type RequestConfigRenderer struct {
	mock.Mock
}

// RenderRequestConfig provides a mock function with given fields:
func (_m *RequestConfigRenderer) RenderRequestConfig() api.RequestConfig {
	ret := _m.Called()

	var r0 api.RequestConfig
	if rf, ok := ret.Get(0).(func() api.RequestConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.RequestConfig)
	}

	return r0
}
