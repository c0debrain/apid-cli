// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/getapid/apid-cli/common/result (interfaces: Writer)

// Package mock is a generated GoMock package.
package mock

import (
	result "github.com/getapid/apid-cli/common/result"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWriter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriter)(nil).Close))
}

// Write mocks base method
func (m *MockWriter) Write(arg0 result.TransactionResult) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0)
}

// Write indicates an expected call of Write
func (mr *MockWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}
