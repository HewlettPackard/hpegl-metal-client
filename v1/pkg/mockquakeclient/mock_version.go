// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_version.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockVersionAPI is a mock of VersionAPI interface.
type MockVersionAPI struct {
	ctrl     *gomock.Controller
	recorder *MockVersionAPIMockRecorder
}

// MockVersionAPIMockRecorder is the mock recorder for MockVersionAPI.
type MockVersionAPIMockRecorder struct {
	mock *MockVersionAPI
}

// NewMockVersionAPI creates a new mock instance.
func NewMockVersionAPI(ctrl *gomock.Controller) *MockVersionAPI {
	mock := &MockVersionAPI{ctrl: ctrl}
	mock.recorder = &MockVersionAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionAPI) EXPECT() *MockVersionAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVersionAPI) Get(ctx context.Context) (client.Version, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(client.Version)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockVersionAPIMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVersionAPI)(nil).Get), ctx)
}
