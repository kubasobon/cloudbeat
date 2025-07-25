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

package registry

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUpdaterFunc is an autogenerated mock type for the UpdaterFunc type
type MockUpdaterFunc struct {
	mock.Mock
}

type MockUpdaterFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdaterFunc) EXPECT() *MockUpdaterFunc_Expecter {
	return &MockUpdaterFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx
func (_m *MockUpdaterFunc) Execute(ctx context.Context) (FetchersMap, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 FetchersMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (FetchersMap, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) FetchersMap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(FetchersMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdaterFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockUpdaterFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUpdaterFunc_Expecter) Execute(ctx interface{}) *MockUpdaterFunc_Execute_Call {
	return &MockUpdaterFunc_Execute_Call{Call: _e.mock.On("Execute", ctx)}
}

func (_c *MockUpdaterFunc_Execute_Call) Run(run func(ctx context.Context)) *MockUpdaterFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUpdaterFunc_Execute_Call) Return(_a0 FetchersMap, _a1 error) *MockUpdaterFunc_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdaterFunc_Execute_Call) RunAndReturn(run func(context.Context) (FetchersMap, error)) *MockUpdaterFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdaterFunc creates a new instance of MockUpdaterFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdaterFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdaterFunc {
	mock := &MockUpdaterFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
