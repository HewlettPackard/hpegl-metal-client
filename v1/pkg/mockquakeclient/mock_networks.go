// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_networks.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockNetworksAPI is a mock of NetworksAPI interface.
type MockNetworksAPI struct {
	ctrl     *gomock.Controller
	recorder *MockNetworksAPIMockRecorder
}

// MockNetworksAPIMockRecorder is the mock recorder for MockNetworksAPI.
type MockNetworksAPIMockRecorder struct {
	mock *MockNetworksAPI
}

// NewMockNetworksAPI creates a new mock instance.
func NewMockNetworksAPI(ctrl *gomock.Controller) *MockNetworksAPI {
	mock := &MockNetworksAPI{ctrl: ctrl}
	mock.recorder = &MockNetworksAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworksAPI) EXPECT() *MockNetworksAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockNetworksAPI) Add(ctx context.Context, newNetwork client.NewNetwork, localVarOptionals *client.NetworksApiAddOpts) (client.Network, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, newNetwork, localVarOptionals)
	ret0, _ := ret[0].(client.Network)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockNetworksAPIMockRecorder) Add(ctx, newNetwork, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockNetworksAPI)(nil).Add), ctx, newNetwork, localVarOptionals)
}

// Delete mocks base method.
func (m *MockNetworksAPI) Delete(ctx context.Context, networkId string, localVarOptionals *client.NetworksApiDeleteOpts) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, networkId, localVarOptionals)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockNetworksAPIMockRecorder) Delete(ctx, networkId, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetworksAPI)(nil).Delete), ctx, networkId, localVarOptionals)
}

// GetByID mocks base method.
func (m *MockNetworksAPI) GetByID(ctx context.Context, networkId string, localVarOptionals *client.NetworksApiGetByIDOpts) (client.Network, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, networkId, localVarOptionals)
	ret0, _ := ret[0].(client.Network)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockNetworksAPIMockRecorder) GetByID(ctx, networkId, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockNetworksAPI)(nil).GetByID), ctx, networkId, localVarOptionals)
}

// List mocks base method.
func (m *MockNetworksAPI) List(ctx context.Context, localVarOptionals *client.NetworksApiListOpts) ([]client.Network, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, localVarOptionals)
	ret0, _ := ret[0].([]client.Network)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockNetworksAPIMockRecorder) List(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworksAPI)(nil).List), ctx, localVarOptionals)
}

// Update mocks base method.
func (m *MockNetworksAPI) Update(ctx context.Context, networkId string, updateNetwork client.UpdateNetwork, localVarOptionals *client.NetworksApiUpdateOpts) (client.Network, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, networkId, updateNetwork, localVarOptionals)
	ret0, _ := ret[0].(client.Network)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockNetworksAPIMockRecorder) Update(ctx, networkId, updateNetwork, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetworksAPI)(nil).Update), ctx, networkId, updateNetwork, localVarOptionals)
}
