// Code generated by MockGen. DO NOT EDIT.
// Source: grpcserver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGRPCServer is a mock of GRPCServer interface.
type MockGRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCServerMockRecorder
}

// MockGRPCServerMockRecorder is the mock recorder for MockGRPCServer.
type MockGRPCServerMockRecorder struct {
	mock *MockGRPCServer
}

// NewMockGRPCServer creates a new mock instance.
func NewMockGRPCServer(ctrl *gomock.Controller) *MockGRPCServer {
	mock := &MockGRPCServer{ctrl: ctrl}
	mock.recorder = &MockGRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCServer) EXPECT() *MockGRPCServerMockRecorder {
	return m.recorder
}

// GracefulStop mocks base method.
func (m *MockGRPCServer) GracefulStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GracefulStop")
}

// GracefulStop indicates an expected call of GracefulStop.
func (mr *MockGRPCServerMockRecorder) GracefulStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GracefulStop", reflect.TypeOf((*MockGRPCServer)(nil).GracefulStop))
}

// Serve mocks base method.
func (m *MockGRPCServer) Serve(listener net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", listener)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockGRPCServerMockRecorder) Serve(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockGRPCServer)(nil).Serve), listener)
}
