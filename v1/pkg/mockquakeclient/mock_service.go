// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_service.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockServiceAPI is a mock of ServiceAPI interface.
type MockServiceAPI struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAPIMockRecorder
}

// MockServiceAPIMockRecorder is the mock recorder for MockServiceAPI.
type MockServiceAPIMockRecorder struct {
	mock *MockServiceAPI
}

// NewMockServiceAPI creates a new mock instance.
func NewMockServiceAPI(ctrl *gomock.Controller) *MockServiceAPI {
	mock := &MockServiceAPI{ctrl: ctrl}
	mock.recorder = &MockServiceAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAPI) EXPECT() *MockServiceAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockServiceAPI) Add(ctx context.Context, fileName *os.File, localVarOptionals *client.ServiceApiAddOpts) (client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, fileName, localVarOptionals)
	ret0, _ := ret[0].(client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockServiceAPIMockRecorder) Add(ctx, fileName, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockServiceAPI)(nil).Add), ctx, fileName, localVarOptionals)
}