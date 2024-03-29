// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pseudomuto/pseudocms/pkg/server (interfaces: DefinitionsRepo,FieldsRepo)
// Package server_test is a generated GoMock package.
package server_test

import (
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	models "github.com/pseudomuto/pseudocms/pkg/models"
	"github.com/pseudomuto/pseudocms/pkg/repo"
)

// MockDefinitionsRepo is a mock of DefinitionsRepo interface
type MockDefinitionsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDefinitionsRepoMockRecorder
}

// MockDefinitionsRepoMockRecorder is the mock recorder for MockDefinitionsRepo
type MockDefinitionsRepoMockRecorder struct {
	mock *MockDefinitionsRepo
}

// NewMockDefinitionsRepo creates a new mock instance
func NewMockDefinitionsRepo(ctrl *gomock.Controller) *MockDefinitionsRepo {
	mock := &MockDefinitionsRepo{ctrl: ctrl}
	mock.recorder = &MockDefinitionsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDefinitionsRepo) EXPECT() *MockDefinitionsRepoMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockDefinitionsRepo) Create(arg0 *models.Definition, arg1 repo.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockDefinitionsRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDefinitionsRepo)(nil).Create), arg0, arg1)
}

// Find mocks base method
func (m *MockDefinitionsRepo) Find(arg0 uuid.UUID, arg1 repo.FindOptions) (*models.Definition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*models.Definition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockDefinitionsRepoMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockDefinitionsRepo)(nil).Find), arg0, arg1)
}

// List mocks base method
func (m *MockDefinitionsRepo) List(arg0 ...repo.ListOption) (*repo.ListResult[models.Definition], error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*repo.ListResult[models.Definition])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDefinitionsRepoMockRecorder) List(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDefinitionsRepo)(nil).List), arg0...)
}

// Update mocks base method
func (m *MockDefinitionsRepo) Update(arg0 *models.Definition, arg1 repo.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDefinitionsRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDefinitionsRepo)(nil).Update), arg0, arg1)
}

// MockFieldsRepo is a mock of FieldsRepo interface
type MockFieldsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockFieldsRepoMockRecorder
}

// MockFieldsRepoMockRecorder is the mock recorder for MockFieldsRepo
type MockFieldsRepoMockRecorder struct {
	mock *MockFieldsRepo
}

// NewMockFieldsRepo creates a new mock instance
func NewMockFieldsRepo(ctrl *gomock.Controller) *MockFieldsRepo {
	mock := &MockFieldsRepo{ctrl: ctrl}
	mock.recorder = &MockFieldsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFieldsRepo) EXPECT() *MockFieldsRepoMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockFieldsRepo) Create(arg0 *models.Field, arg1 repo.CreateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockFieldsRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFieldsRepo)(nil).Create), arg0, arg1)
}

// Find mocks base method
func (m *MockFieldsRepo) Find(arg0 uuid.UUID, arg1 repo.FindOptions) (*models.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*models.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFieldsRepoMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFieldsRepo)(nil).Find), arg0, arg1)
}

// List mocks base method
func (m *MockFieldsRepo) List(arg0 ...repo.ListOption) (*repo.ListResult[models.Field], error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*repo.ListResult[models.Field])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockFieldsRepoMockRecorder) List(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFieldsRepo)(nil).List), arg0...)
}

// Update mocks base method
func (m *MockFieldsRepo) Update(arg0 *models.Field, arg1 repo.UpdateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockFieldsRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFieldsRepo)(nil).Update), arg0, arg1)
}
