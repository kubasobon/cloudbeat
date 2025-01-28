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

package awslib

import (
	aws "github.com/aws/aws-sdk-go-v2/aws"
	cloud "github.com/elastic/cloudbeat/internal/dataprovider/providers/cloud"
	clog "github.com/elastic/cloudbeat/internal/infra/clog"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountProviderAPI is an autogenerated mock type for the AccountProviderAPI type
type MockAccountProviderAPI struct {
	mock.Mock
}

type MockAccountProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountProviderAPI) EXPECT() *MockAccountProviderAPI_Expecter {
	return &MockAccountProviderAPI_Expecter{mock: &_m.Mock}
}

// ListAccounts provides a mock function with given fields: ctx, log, cfg
func (_m *MockAccountProviderAPI) ListAccounts(ctx context.Context, log *clog.Logger, cfg aws.Config) ([]cloud.Identity, error) {
	ret := _m.Called(ctx, log, cfg)

	var r0 []cloud.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *clog.Logger, aws.Config) ([]cloud.Identity, error)); ok {
		return rf(ctx, log, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *clog.Logger, aws.Config) []cloud.Identity); ok {
		r0 = rf(ctx, log, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cloud.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *clog.Logger, aws.Config) error); ok {
		r1 = rf(ctx, log, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountProviderAPI_ListAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccounts'
type MockAccountProviderAPI_ListAccounts_Call struct {
	*mock.Call
}

// ListAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - log *clog.Logger
//   - cfg aws.Config
func (_e *MockAccountProviderAPI_Expecter) ListAccounts(ctx interface{}, log interface{}, cfg interface{}) *MockAccountProviderAPI_ListAccounts_Call {
	return &MockAccountProviderAPI_ListAccounts_Call{Call: _e.mock.On("ListAccounts", ctx, log, cfg)}
}

func (_c *MockAccountProviderAPI_ListAccounts_Call) Run(run func(ctx context.Context, log *clog.Logger, cfg aws.Config)) *MockAccountProviderAPI_ListAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*clog.Logger), args[2].(aws.Config))
	})
	return _c
}

func (_c *MockAccountProviderAPI_ListAccounts_Call) Return(_a0 []cloud.Identity, _a1 error) *MockAccountProviderAPI_ListAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountProviderAPI_ListAccounts_Call) RunAndReturn(run func(context.Context, *clog.Logger, aws.Config) ([]cloud.Identity, error)) *MockAccountProviderAPI_ListAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountProviderAPI creates a new instance of MockAccountProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountProviderAPI {
	mock := &MockAccountProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
