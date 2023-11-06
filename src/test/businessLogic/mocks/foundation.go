// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/internal/business/services/foundation_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// MockIFoundationRepository is a mock of IFoundationRepository interface.
type MockIFoundationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFoundationRepositoryMockRecorder
}

// MockIFoundationRepositoryMockRecorder is the mock recorder for MockIFoundationRepository.
type MockIFoundationRepositoryMockRecorder struct {
	mock *MockIFoundationRepository
}

// NewMockIFoundationRepository creates a new mock instance.
func NewMockIFoundationRepository(ctrl *gomock.Controller) *MockIFoundationRepository {
	mock := &MockIFoundationRepository{ctrl: ctrl}
	mock.recorder = &MockIFoundationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFoundationRepository) EXPECT() *MockIFoundationRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIFoundationRepository) Delete(F entities.Foundation) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", F)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIFoundationRepositoryMockRecorder) Delete(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIFoundationRepository)(nil).Delete), F)
}

// Insert mocks base method.
func (m *MockIFoundationRepository) Insert(F entities.Foundation) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", F)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockIFoundationRepositoryMockRecorder) Insert(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIFoundationRepository)(nil).Insert), F)
}

// Select mocks base method.
func (m *MockIFoundationRepository) Select() ([]entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select")
	ret0, _ := ret[0].([]entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockIFoundationRepositoryMockRecorder) Select() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIFoundationRepository)(nil).Select))
}

// SelectByCountry mocks base method.
func (m *MockIFoundationRepository) SelectByCountry(country string) ([]entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByCountry", country)
	ret0, _ := ret[0].([]entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByCountry indicates an expected call of SelectByCountry.
func (mr *MockIFoundationRepositoryMockRecorder) SelectByCountry(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByCountry", reflect.TypeOf((*MockIFoundationRepository)(nil).SelectByCountry), country)
}

// SelectById mocks base method.
func (m *MockIFoundationRepository) SelectById(id uint64) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockIFoundationRepositoryMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockIFoundationRepository)(nil).SelectById), id)
}

// SelectByLogin mocks base method.
func (m *MockIFoundationRepository) SelectByLogin(name string) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByLogin", name)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByLogin indicates an expected call of SelectByLogin.
func (mr *MockIFoundationRepositoryMockRecorder) SelectByLogin(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByLogin", reflect.TypeOf((*MockIFoundationRepository)(nil).SelectByLogin), name)
}

// SelectByName mocks base method.
func (m *MockIFoundationRepository) SelectByName(name string) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByName", name)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByName indicates an expected call of SelectByName.
func (mr *MockIFoundationRepositoryMockRecorder) SelectByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByName", reflect.TypeOf((*MockIFoundationRepository)(nil).SelectByName), name)
}

// Update mocks base method.
func (m *MockIFoundationRepository) Update(F entities.Foundation) (entities.Foundation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", F)
	ret0, _ := ret[0].(entities.Foundation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIFoundationRepositoryMockRecorder) Update(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIFoundationRepository)(nil).Update), F)
}
