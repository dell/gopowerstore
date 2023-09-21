/*
 *
 * Copyright © 2021-2022 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	api "github.com/dell/gopowerstore/api"
	context "context"
	mock "github.com/stretchr/testify/mock"
)

// TimeoutSemaphoreInterface is an autogenerated mock type for the TimeoutSemaphoreInterface type
type TimeoutSemaphoreInterface struct {
	mock.Mock
}

// Acquire provides a mock function with given fields: ctx
func (_m *TimeoutSemaphoreInterface) Acquire(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Release provides a mock function with given fields: ctx
func (_m *TimeoutSemaphoreInterface) Release(ctx context.Context) {
	_m.Called(ctx)
}

// SetLogger provides a mock function with given fields: logger
func (_m *TimeoutSemaphoreInterface) SetLogger(logger api.Logger) api.TimeoutSemaphoreInterface {
	ret := _m.Called(logger)

	var r0 api.TimeoutSemaphoreInterface
	if rf, ok := ret.Get(0).(func(api.Logger) api.TimeoutSemaphoreInterface); ok {
		r0 = rf(logger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.TimeoutSemaphoreInterface)
		}
	}

	return r0
}
