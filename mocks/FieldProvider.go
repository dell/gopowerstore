// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FieldProvider is an autogenerated mock type for the FieldProvider type
type FieldProvider struct {
	mock.Mock
}

// Fields provides a mock function with given fields:
func (_m *FieldProvider) Fields() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
