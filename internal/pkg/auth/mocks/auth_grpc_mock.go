// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen (interfaces: AuthServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gen "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockAuthServiceClient is a mock of AuthServiceClient interface.
type MockAuthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceClientMockRecorder
}

// MockAuthServiceClientMockRecorder is the mock recorder for MockAuthServiceClient.
type MockAuthServiceClientMockRecorder struct {
	mock *MockAuthServiceClient
}

// NewMockAuthServiceClient creates a new mock instance.
func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
	mock := &MockAuthServiceClient{ctrl: ctrl}
	mock.recorder = &MockAuthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockAuthServiceClient) AddAddress(arg0 context.Context, arg1 *gen.Address, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAddress", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockAuthServiceClientMockRecorder) AddAddress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockAuthServiceClient)(nil).AddAddress), varargs...)
}

// Check mocks base method.
func (m *MockAuthServiceClient) Check(arg0 context.Context, arg1 *gen.CheckRequest, arg2 ...grpc.CallOption) (*gen.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Check", varargs...)
	ret0, _ := ret[0].(*gen.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockAuthServiceClientMockRecorder) Check(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthServiceClient)(nil).Check), varargs...)
}

// DeleteAddress mocks base method.
func (m *MockAuthServiceClient) DeleteAddress(arg0 context.Context, arg1 *gen.DeleteAddressRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAddress", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockAuthServiceClientMockRecorder) DeleteAddress(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockAuthServiceClient)(nil).DeleteAddress), varargs...)
}

// GetUserAddresses mocks base method.
func (m *MockAuthServiceClient) GetUserAddresses(arg0 context.Context, arg1 *gen.AddressRequest, arg2 ...grpc.CallOption) (*gen.AddressListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserAddresses", varargs...)
	ret0, _ := ret[0].(*gen.AddressListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAddresses indicates an expected call of GetUserAddresses.
func (mr *MockAuthServiceClientMockRecorder) GetUserAddresses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAddresses", reflect.TypeOf((*MockAuthServiceClient)(nil).GetUserAddresses), varargs...)
}

// SignIn mocks base method.
func (m *MockAuthServiceClient) SignIn(arg0 context.Context, arg1 *gen.SignInRequest, arg2 ...grpc.CallOption) (*gen.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignIn", varargs...)
	ret0, _ := ret[0].(*gen.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthServiceClientMockRecorder) SignIn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthServiceClient)(nil).SignIn), varargs...)
}

// SignUp mocks base method.
func (m *MockAuthServiceClient) SignUp(arg0 context.Context, arg1 *gen.SignUpRequest, arg2 ...grpc.CallOption) (*gen.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignUp", varargs...)
	ret0, _ := ret[0].(*gen.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthServiceClientMockRecorder) SignUp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthServiceClient)(nil).SignUp), varargs...)
}

// UpdateUser mocks base method.
func (m *MockAuthServiceClient) UpdateUser(arg0 context.Context, arg1 *gen.UpdateUserRequest, arg2 ...grpc.CallOption) (*gen.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUser", varargs...)
	ret0, _ := ret[0].(*gen.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAuthServiceClientMockRecorder) UpdateUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateUser), varargs...)
}

// UpdateUserPic mocks base method.
func (m *MockAuthServiceClient) UpdateUserPic(arg0 context.Context, arg1 *gen.UpdateUserPicRequest, arg2 ...grpc.CallOption) (*gen.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserPic", varargs...)
	ret0, _ := ret[0].(*gen.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPic indicates an expected call of UpdateUserPic.
func (mr *MockAuthServiceClientMockRecorder) UpdateUserPic(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPic", reflect.TypeOf((*MockAuthServiceClient)(nil).UpdateUserPic), varargs...)
}
