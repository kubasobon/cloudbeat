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

// Code generated by mockery v2.37.1. DO NOT EDIT.

package preset

import (
	context "context"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	cloud "github.com/elastic/cloudbeat/internal/dataprovider/providers/cloud"

	fetching "github.com/elastic/cloudbeat/internal/resources/fetching"

	logp "github.com/elastic/elastic-agent-libs/logp"

	mock "github.com/stretchr/testify/mock"

	registry "github.com/elastic/cloudbeat/internal/resources/fetching/registry"
)

// mockAwsFactory is an autogenerated mock type for the awsFactory type
type mockAwsFactory struct {
	mock.Mock
}

type mockAwsFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAwsFactory) EXPECT() *mockAwsFactory_Expecter {
	return &mockAwsFactory_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *mockAwsFactory) Execute(_a0 context.Context, _a1 *logp.Logger, _a2 aws.Config, _a3 chan fetching.ResourceInfo, _a4 *cloud.Identity) registry.FetchersMap {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 registry.FetchersMap
	if rf, ok := ret.Get(0).(func(context.Context, *logp.Logger, aws.Config, chan fetching.ResourceInfo, *cloud.Identity) registry.FetchersMap); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.FetchersMap)
		}
	}

	return r0
}

// mockAwsFactory_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type mockAwsFactory_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *logp.Logger
//   - _a2 aws.Config
//   - _a3 chan fetching.ResourceInfo
//   - _a4 *cloud.Identity
func (_e *mockAwsFactory_Expecter) Execute(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}, _a4 interface{}) *mockAwsFactory_Execute_Call {
	return &mockAwsFactory_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1, _a2, _a3, _a4)}
}

func (_c *mockAwsFactory_Execute_Call) Run(run func(_a0 context.Context, _a1 *logp.Logger, _a2 aws.Config, _a3 chan fetching.ResourceInfo, _a4 *cloud.Identity)) *mockAwsFactory_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*logp.Logger), args[2].(aws.Config), args[3].(chan fetching.ResourceInfo), args[4].(*cloud.Identity))
	})
	return _c
}

func (_c *mockAwsFactory_Execute_Call) Return(_a0 registry.FetchersMap) *mockAwsFactory_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAwsFactory_Execute_Call) RunAndReturn(run func(context.Context, *logp.Logger, aws.Config, chan fetching.ResourceInfo, *cloud.Identity) registry.FetchersMap) *mockAwsFactory_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// newMockAwsFactory creates a new instance of mockAwsFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAwsFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAwsFactory {
	mock := &mockAwsFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
