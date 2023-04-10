// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux_bpf
// +build linux_bpf

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/network/netlink/conntrack.go

// Package netlink is a generated GoMock package.
package netlink

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConntrack is a mock of Conntrack interface
type MockConntrack struct {
	ctrl     *gomock.Controller
	recorder *MockConntrackMockRecorder
}

// MockConntrackMockRecorder is the mock recorder for MockConntrack
type MockConntrackMockRecorder struct {
	mock *MockConntrack
}

// NewMockConntrack creates a new mock instance
func NewMockConntrack(ctrl *gomock.Controller) *MockConntrack {
	mock := &MockConntrack{ctrl: ctrl}
	mock.recorder = &MockConntrackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConntrack) EXPECT() *MockConntrackMockRecorder {
	return m.recorder
}

// Exists mocks base method
func (m *MockConntrack) Exists(conn *Con) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", conn)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockConntrackMockRecorder) Exists(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockConntrack)(nil).Exists), conn)
}

// Dump mocks base method
func (m *MockConntrack) Dump() ([]Con, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump")
	ret0, _ := ret[0].([]Con)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dump indicates an expected call of Dump
func (mr *MockConntrackMockRecorder) Dump() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockConntrack)(nil).Dump))
}

// Get mocks base method
func (m *MockConntrack) Get(conn *Con) (Con, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", conn)
	ret0, _ := ret[0].(Con)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockConntrackMockRecorder) Get(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConntrack)(nil).Get), conn)
}

// Close mocks base method
func (m *MockConntrack) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConntrackMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConntrack)(nil).Close))
}
