// Code generated by MockGen. DO NOT EDIT.
// Source: proto/api/v1beta1/api.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/socketfunc/colony/proto/api/v1beta1"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockApiServiceClient is a mock of ApiServiceClient interface
type MockApiServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceClientMockRecorder
}

// MockApiServiceClientMockRecorder is the mock recorder for MockApiServiceClient
type MockApiServiceClientMockRecorder struct {
	mock *MockApiServiceClient
}

// NewMockApiServiceClient creates a new mock instance
func NewMockApiServiceClient(ctrl *gomock.Controller) *MockApiServiceClient {
	mock := &MockApiServiceClient{ctrl: ctrl}
	mock.recorder = &MockApiServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceClient) EXPECT() *MockApiServiceClientMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockApiServiceClient) GetConfig(ctx context.Context, in *v1beta1.GetConfigRequest, opts ...grpc.CallOption) (*v1beta1.GetConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].(*v1beta1.GetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockApiServiceClientMockRecorder) GetConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockApiServiceClient)(nil).GetConfig), varargs...)
}

// SaveConfig mocks base method
func (m *MockApiServiceClient) SaveConfig(ctx context.Context, in *v1beta1.SaveConfigRequest, opts ...grpc.CallOption) (*v1beta1.SaveConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveConfig", varargs...)
	ret0, _ := ret[0].(*v1beta1.SaveConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveConfig indicates an expected call of SaveConfig
func (mr *MockApiServiceClientMockRecorder) SaveConfig(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveConfig", reflect.TypeOf((*MockApiServiceClient)(nil).SaveConfig), varargs...)
}

// MockApiServiceServer is a mock of ApiServiceServer interface
type MockApiServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceServerMockRecorder
}

// MockApiServiceServerMockRecorder is the mock recorder for MockApiServiceServer
type MockApiServiceServerMockRecorder struct {
	mock *MockApiServiceServer
}

// NewMockApiServiceServer creates a new mock instance
func NewMockApiServiceServer(ctrl *gomock.Controller) *MockApiServiceServer {
	mock := &MockApiServiceServer{ctrl: ctrl}
	mock.recorder = &MockApiServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceServer) EXPECT() *MockApiServiceServerMockRecorder {
	return m.recorder
}

// GetConfig mocks base method
func (m *MockApiServiceServer) GetConfig(arg0 context.Context, arg1 *v1beta1.GetConfigRequest) (*v1beta1.GetConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.GetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockApiServiceServerMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockApiServiceServer)(nil).GetConfig), arg0, arg1)
}

// SaveConfig mocks base method
func (m *MockApiServiceServer) SaveConfig(arg0 context.Context, arg1 *v1beta1.SaveConfigRequest) (*v1beta1.SaveConfigResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveConfig", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.SaveConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveConfig indicates an expected call of SaveConfig
func (mr *MockApiServiceServerMockRecorder) SaveConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveConfig", reflect.TypeOf((*MockApiServiceServer)(nil).SaveConfig), arg0, arg1)
}
