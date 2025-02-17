// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022 Datadog, Inc.
// Code generated by mockery v2.14.0. DO NOT EDIT.

package slack

import (
	slack_goslack "github.com/slack-go/slack"
	mock "github.com/stretchr/testify/mock"
)

// mockSlackNotifier is an autogenerated mock type for the slackNotifier type
type mockSlackNotifier struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *mockSlackNotifier) GetUserByEmail(email string) (*slack_goslack.User, error) {
	ret := _m.Called(email)

	var r0 *slack_goslack.User
	if rf, ok := ret.Get(0).(func(string) *slack_goslack.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack_goslack.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostMessage provides a mock function with given fields: channelID, options
func (_m *mockSlackNotifier) PostMessage(channelID string, options ...slack_goslack.MsgOption) (string, string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, channelID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...slack_goslack.MsgOption) string); ok {
		r0 = rf(channelID, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, ...slack_goslack.MsgOption) string); ok {
		r1 = rf(channelID, options...)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, ...slack_goslack.MsgOption) error); ok {
		r2 = rf(channelID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTnewMockSlackNotifier interface {
	mock.TestingT
	Cleanup(func())
}

// newMockSlackNotifier creates a new instance of mockSlackNotifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockSlackNotifier(t mockConstructorTestingTnewMockSlackNotifier) *mockSlackNotifier {
	mock := &mockSlackNotifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
