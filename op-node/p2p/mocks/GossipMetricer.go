// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GossipMetricer is an autogenerated mock type for the GossipMetricer type
type GossipMetricer struct {
	mock.Mock
}

// RecordGossipEvent provides a mock function with given fields: evType
func (_m *GossipMetricer) RecordGossipEvent(evType int32) {
	_m.Called(evType)
}

// SetPeerScores provides a mock function with given fields: _a0
func (_m *GossipMetricer) SetPeerScores(_a0 map[string]float64) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewGossipMetricer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipMetricer creates a new instance of GossipMetricer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipMetricer(t mockConstructorTestingTNewGossipMetricer) *GossipMetricer {
	mock := &GossipMetricer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
