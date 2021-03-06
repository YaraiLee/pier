// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package mock_router is a generated GoMock package.
package mock_router

import (
	gomock "github.com/golang/mock/gomock"
	pb "github.com/meshplus/bitxhub-model/pb"
	rpcx "github.com/meshplus/go-bitxhub-client"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockRouter) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockRouterMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRouter)(nil).Start))
}

// Stop mocks base method
func (m *MockRouter) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockRouterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRouter)(nil).Stop))
}

// Broadcast mocks base method
func (m *MockRouter) Broadcast(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockRouterMockRecorder) Broadcast(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockRouter)(nil).Broadcast), ids)
}

// Route mocks base method
func (m *MockRouter) Route(ibtp *pb.IBTP) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", ibtp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Route indicates an expected call of Route
func (mr *MockRouterMockRecorder) Route(ibtp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), ibtp)
}

// ExistAppchain mocks base method
func (m *MockRouter) ExistAppchain(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistAppchain", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistAppchain indicates an expected call of ExistAppchain
func (mr *MockRouterMockRecorder) ExistAppchain(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistAppchain", reflect.TypeOf((*MockRouter)(nil).ExistAppchain), id)
}

// AddAppchains mocks base method
func (m *MockRouter) AddAppchains(appchains []*rpcx.Appchain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppchains", appchains)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppchains indicates an expected call of AddAppchains
func (mr *MockRouterMockRecorder) AddAppchains(appchains interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppchains", reflect.TypeOf((*MockRouter)(nil).AddAppchains), appchains)
}
