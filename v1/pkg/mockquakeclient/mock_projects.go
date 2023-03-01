// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_projects.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockProjectsAPI is a mock of ProjectsAPI interface.
type MockProjectsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsAPIMockRecorder
}

// MockProjectsAPIMockRecorder is the mock recorder for MockProjectsAPI.
type MockProjectsAPIMockRecorder struct {
	mock *MockProjectsAPI
}

// NewMockProjectsAPI creates a new mock instance.
func NewMockProjectsAPI(ctrl *gomock.Controller) *MockProjectsAPI {
	mock := &MockProjectsAPI{ctrl: ctrl}
	mock.recorder = &MockProjectsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectsAPI) EXPECT() *MockProjectsAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockProjectsAPI) Add(ctx context.Context, newProject client.NewProject, localVarOptionals *client.ProjectsApiAddOpts) (client.Project, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, newProject, localVarOptionals)
	ret0, _ := ret[0].(client.Project)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockProjectsAPIMockRecorder) Add(ctx, newProject, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProjectsAPI)(nil).Add), ctx, newProject, localVarOptionals)
}

// Delete mocks base method.
func (m *MockProjectsAPI) Delete(ctx context.Context, projectId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, projectId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockProjectsAPIMockRecorder) Delete(ctx, projectId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProjectsAPI)(nil).Delete), ctx, projectId)
}

// GetByID mocks base method.
func (m *MockProjectsAPI) GetByID(ctx context.Context, projectId string) (client.Project, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, projectId)
	ret0, _ := ret[0].(client.Project)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProjectsAPIMockRecorder) GetByID(ctx, projectId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProjectsAPI)(nil).GetByID), ctx, projectId)
}

// ImagesGetByID mocks base method.
func (m *MockProjectsAPI) ImagesGetByID(ctx context.Context, projectId, imageId string) ([]client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagesGetByID", ctx, projectId, imageId)
	ret0, _ := ret[0].([]client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImagesGetByID indicates an expected call of ImagesGetByID.
func (mr *MockProjectsAPIMockRecorder) ImagesGetByID(ctx, projectId, imageId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagesGetByID", reflect.TypeOf((*MockProjectsAPI)(nil).ImagesGetByID), ctx, projectId, imageId)
}

// ImagesList mocks base method.
func (m *MockProjectsAPI) ImagesList(ctx context.Context, projectId string) ([]client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagesList", ctx, projectId)
	ret0, _ := ret[0].([]client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImagesList indicates an expected call of ImagesList.
func (mr *MockProjectsAPIMockRecorder) ImagesList(ctx, projectId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagesList", reflect.TypeOf((*MockProjectsAPI)(nil).ImagesList), ctx, projectId)
}

// List mocks base method.
func (m *MockProjectsAPI) List(ctx context.Context, localVarOptionals *client.ProjectsApiListOpts) ([]client.Project, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, localVarOptionals)
	ret0, _ := ret[0].([]client.Project)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockProjectsAPIMockRecorder) List(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectsAPI)(nil).List), ctx, localVarOptionals)
}

// Update mocks base method.
func (m *MockProjectsAPI) Update(ctx context.Context, projectId string, updateProject client.UpdateProject) (client.Project, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, projectId, updateProject)
	ret0, _ := ret[0].(client.Project)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockProjectsAPIMockRecorder) Update(ctx, projectId, updateProject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProjectsAPI)(nil).Update), ctx, projectId, updateProject)
}
