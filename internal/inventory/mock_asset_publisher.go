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

package inventory

import (
	beat "github.com/elastic/beats/v7/libbeat/beat"
	mock "github.com/stretchr/testify/mock"
)

// MockAssetPublisher is an autogenerated mock type for the AssetPublisher type
type MockAssetPublisher struct {
	mock.Mock
}

type MockAssetPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAssetPublisher) EXPECT() *MockAssetPublisher_Expecter {
	return &MockAssetPublisher_Expecter{mock: &_m.Mock}
}

// PublishAll provides a mock function with given fields: _a0
func (_m *MockAssetPublisher) PublishAll(_a0 []beat.Event) {
	_m.Called(_a0)
}

// MockAssetPublisher_PublishAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PublishAll'
type MockAssetPublisher_PublishAll_Call struct {
	*mock.Call
}

// PublishAll is a helper method to define mock.On call
//   - _a0 []beat.Event
func (_e *MockAssetPublisher_Expecter) PublishAll(_a0 interface{}) *MockAssetPublisher_PublishAll_Call {
	return &MockAssetPublisher_PublishAll_Call{Call: _e.mock.On("PublishAll", _a0)}
}

func (_c *MockAssetPublisher_PublishAll_Call) Run(run func(_a0 []beat.Event)) *MockAssetPublisher_PublishAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]beat.Event))
	})
	return _c
}

func (_c *MockAssetPublisher_PublishAll_Call) Return() *MockAssetPublisher_PublishAll_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockAssetPublisher_PublishAll_Call) RunAndReturn(run func([]beat.Event)) *MockAssetPublisher_PublishAll_Call {
	_c.Run(run)
	return _c
}

// NewMockAssetPublisher creates a new instance of MockAssetPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAssetPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAssetPublisher {
	mock := &MockAssetPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
