// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.53.3. DO NOT EDIT.

package uniqueness

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// IsLeader provides a mock function with no fields
func (_m *MockManager) IsLeader() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsLeader")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockManager_IsLeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsLeader'
type MockManager_IsLeader_Call struct {
	*mock.Call
}

// IsLeader is a helper method to define mock.On call
func (_e *MockManager_Expecter) IsLeader() *MockManager_IsLeader_Call {
	return &MockManager_IsLeader_Call{Call: _e.mock.On("IsLeader")}
}

func (_c *MockManager_IsLeader_Call) Run(run func()) *MockManager_IsLeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_IsLeader_Call) Return(_a0 bool) *MockManager_IsLeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_IsLeader_Call) RunAndReturn(run func() bool) *MockManager_IsLeader_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockManager) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockManager_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockManager_Expecter) Run(ctx interface{}) *MockManager_Run_Call {
	return &MockManager_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockManager_Run_Call) Run(run func(ctx context.Context)) *MockManager_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockManager_Run_Call) Return(_a0 error) *MockManager_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_Run_Call) RunAndReturn(run func(context.Context) error) *MockManager_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *MockManager) Stop() {
	_m.Called()
}

// MockManager_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockManager_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockManager_Expecter) Stop() *MockManager_Stop_Call {
	return &MockManager_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockManager_Stop_Call) Run(run func()) *MockManager_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_Stop_Call) Return() *MockManager_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_Stop_Call) RunAndReturn(run func()) *MockManager_Stop_Call {
	_c.Run(run)
	return _c
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
