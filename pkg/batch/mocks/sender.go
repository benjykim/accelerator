// Code generated by MockGen. DO NOT EDIT.
// Source: route/sender.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	tx "github.com/nexledger/accelerator/pkg/batch/tx"
	reflect "reflect"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSender) Send(job *tx.Job) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", job)
}

// Send indicates an expected call of Send
func (mr *MockSenderMockRecorder) Send(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSender)(nil).Send), job)
}
