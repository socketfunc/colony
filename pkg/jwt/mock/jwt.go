// Code generated by MockGen. DO NOT EDIT.
// Source: jwt.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	jwt "github.com/socketfunc/colony/pkg/jwt"
	reflect "reflect"
)

// MockJWT is a mock of JWT interface
type MockJWT struct {
	ctrl     *gomock.Controller
	recorder *MockJWTMockRecorder
}

// MockJWTMockRecorder is the mock recorder for MockJWT
type MockJWTMockRecorder struct {
	mock *MockJWT
}

// NewMockJWT creates a new mock instance
func NewMockJWT(ctrl *gomock.Controller) *MockJWT {
	mock := &MockJWT{ctrl: ctrl}
	mock.recorder = &MockJWTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJWT) EXPECT() *MockJWTMockRecorder {
	return m.recorder
}

// Sign mocks base method
func (m *MockJWT) Sign(payload *jwt.Payload) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockJWTMockRecorder) Sign(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockJWT)(nil).Sign), payload)
}

// Verify mocks base method
func (m *MockJWT) Verify(token string) (*jwt.Payload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", token)
	ret0, _ := ret[0].(*jwt.Payload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockJWTMockRecorder) Verify(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockJWT)(nil).Verify), token)
}
