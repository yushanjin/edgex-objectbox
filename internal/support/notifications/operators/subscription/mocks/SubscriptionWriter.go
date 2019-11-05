// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// SubscriptionWriter is an autogenerated mock type for the SubscriptionWriter type
type SubscriptionWriter struct {
	mock.Mock
}

// AddSubscription provides a mock function with given fields: s
func (_m *SubscriptionWriter) AddSubscription(s models.Subscription) (string, error) {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Subscription) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Subscription) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
