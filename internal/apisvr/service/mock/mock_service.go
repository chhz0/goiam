// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	v1 "github.com/chhz0/goiam/internal/apisvr/service/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Policies mocks base method.
func (m *MockService) Policies() v1.PolicySrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policies")
	ret0, _ := ret[0].(v1.PolicySrv)
	return ret0
}

// Policies indicates an expected call of Policies.
func (mr *MockServiceMockRecorder) Policies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policies", reflect.TypeOf((*MockService)(nil).Policies))
}

// Secrets mocks base method.
func (m *MockService) Secrets() v1.SecretSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Secrets")
	ret0, _ := ret[0].(v1.SecretSrv)
	return ret0
}

// Secrets indicates an expected call of Secrets.
func (mr *MockServiceMockRecorder) Secrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secrets", reflect.TypeOf((*MockService)(nil).Secrets))
}

// Users mocks base method.
func (m *MockService) Users() v1.UserSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(v1.UserSrv)
	return ret0
}

// Users indicates an expected call of Users.
func (mr *MockServiceMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockService)(nil).Users))
}
