// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pseudomuto/pseudocms/pkg/api/v1 (interfaces: AdminService_ListDefinitionsServer)

// Package server_test is a generated GoMock package.
package server_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockAdminService_ListDefinitionsServer is a mock of AdminService_ListDefinitionsServer interface
type MockAdminService_ListDefinitionsServer struct {
	ctrl     *gomock.Controller
	recorder *MockAdminService_ListDefinitionsServerMockRecorder
}

// MockAdminService_ListDefinitionsServerMockRecorder is the mock recorder for MockAdminService_ListDefinitionsServer
type MockAdminService_ListDefinitionsServerMockRecorder struct {
	mock *MockAdminService_ListDefinitionsServer
}

// NewMockAdminService_ListDefinitionsServer creates a new mock instance
func NewMockAdminService_ListDefinitionsServer(ctrl *gomock.Controller) *MockAdminService_ListDefinitionsServer {
	mock := &MockAdminService_ListDefinitionsServer{ctrl: ctrl}
	mock.recorder = &MockAdminService_ListDefinitionsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminService_ListDefinitionsServer) EXPECT() *MockAdminService_ListDefinitionsServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockAdminService_ListDefinitionsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockAdminService_ListDefinitionsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockAdminService_ListDefinitionsServer) Send(arg0 *v1.Definition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockAdminService_ListDefinitionsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockAdminService_ListDefinitionsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockAdminService_ListDefinitionsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockAdminService_ListDefinitionsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockAdminService_ListDefinitionsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAdminService_ListDefinitionsServer)(nil).SetTrailer), arg0)
}