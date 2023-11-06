// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/internal/business/services/foundrising_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sabrs0/bmstu-web/src/internal/business/entities"
)

// MockIFoundrisingRepository is a mock of IFoundrisingRepository interface.
type MockIFoundrisingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFoundrisingRepositoryMockRecorder
}

// MockIFoundrisingRepositoryMockRecorder is the mock recorder for MockIFoundrisingRepository.
type MockIFoundrisingRepositoryMockRecorder struct {
	mock *MockIFoundrisingRepository
}

// NewMockIFoundrisingRepository creates a new mock instance.
func NewMockIFoundrisingRepository(ctrl *gomock.Controller) *MockIFoundrisingRepository {
	mock := &MockIFoundrisingRepository{ctrl: ctrl}
	mock.recorder = &MockIFoundrisingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFoundrisingRepository) EXPECT() *MockIFoundrisingRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockIFoundrisingRepository) Delete(F entities.Foundrising) (entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", F)
	ret0, _ := ret[0].(entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIFoundrisingRepositoryMockRecorder) Delete(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIFoundrisingRepository)(nil).Delete), F)
}

// Insert mocks base method.
func (m *MockIFoundrisingRepository) Insert(F entities.Foundrising) (entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", F)
	ret0, _ := ret[0].(entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockIFoundrisingRepositoryMockRecorder) Insert(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIFoundrisingRepository)(nil).Insert), F)
}

// Select mocks base method.
func (m *MockIFoundrisingRepository) Select() ([]entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select")
	ret0, _ := ret[0].([]entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockIFoundrisingRepositoryMockRecorder) Select() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockIFoundrisingRepository)(nil).Select))
}

// SelectByCloseDate mocks base method.
func (m *MockIFoundrisingRepository) SelectByCloseDate(date string) ([]entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByCloseDate", date)
	ret0, _ := ret[0].([]entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByCloseDate indicates an expected call of SelectByCloseDate.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectByCloseDate(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByCloseDate", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectByCloseDate), date)
}

// SelectByCreateDate mocks base method.
func (m *MockIFoundrisingRepository) SelectByCreateDate(date string) ([]entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByCreateDate", date)
	ret0, _ := ret[0].([]entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByCreateDate indicates an expected call of SelectByCreateDate.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectByCreateDate(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByCreateDate", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectByCreateDate), date)
}

// SelectByFoundId mocks base method.
func (m *MockIFoundrisingRepository) SelectByFoundId(id uint64) ([]entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByFoundId", id)
	ret0, _ := ret[0].([]entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByFoundId indicates an expected call of SelectByFoundId.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectByFoundId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByFoundId", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectByFoundId), id)
}

// SelectByFoundIdActive mocks base method.
func (m *MockIFoundrisingRepository) SelectByFoundIdActive(id uint64) ([]entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByFoundIdActive", id)
	ret0, _ := ret[0].([]entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByFoundIdActive indicates an expected call of SelectByFoundIdActive.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectByFoundIdActive(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByFoundIdActive", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectByFoundIdActive), id)
}

// SelectById mocks base method.
func (m *MockIFoundrisingRepository) SelectById(id uint64) (entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectById), id)
}

// SelectByIdAndFoundId mocks base method.
func (m *MockIFoundrisingRepository) SelectByIdAndFoundId(id_, found_id_ uint64) (entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByIdAndFoundId", id_, found_id_)
	ret0, _ := ret[0].(entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByIdAndFoundId indicates an expected call of SelectByIdAndFoundId.
func (mr *MockIFoundrisingRepositoryMockRecorder) SelectByIdAndFoundId(id_, found_id_ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByIdAndFoundId", reflect.TypeOf((*MockIFoundrisingRepository)(nil).SelectByIdAndFoundId), id_, found_id_)
}

// Update mocks base method.
func (m *MockIFoundrisingRepository) Update(F entities.Foundrising) (entities.Foundrising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", F)
	ret0, _ := ret[0].(entities.Foundrising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIFoundrisingRepositoryMockRecorder) Update(F interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIFoundrisingRepository)(nil).Update), F)
}
