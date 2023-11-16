// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/internal/business/services/transaction_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/sabrs0/bmstu-web/src/internal/business/entities"
	services "github.com/sabrs0/bmstu-web/src/internal/business/services"
)

// MockITransactionRepository is a mock of ITransactionRepository interface.
type MockITransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepositoryMockRecorder
}

// MockITransactionRepositoryMockRecorder is the mock recorder for MockITransactionRepository.
type MockITransactionRepositoryMockRecorder struct {
	mock *MockITransactionRepository
}

// NewMockITransactionRepository creates a new mock instance.
func NewMockITransactionRepository(ctrl *gomock.Controller) *MockITransactionRepository {
	mock := &MockITransactionRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepository) EXPECT() *MockITransactionRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockITransactionRepository) Delete(T entities.Transaction) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", T)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockITransactionRepositoryMockRecorder) Delete(T interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITransactionRepository)(nil).Delete), T)
}

// Insert mocks base method.
func (m *MockITransactionRepository) Insert(T entities.Transaction) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", T)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockITransactionRepositoryMockRecorder) Insert(T interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockITransactionRepository)(nil).Insert), T)
}

// Select mocks base method.
func (m *MockITransactionRepository) Select() ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select")
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockITransactionRepositoryMockRecorder) Select() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockITransactionRepository)(nil).Select))
}

// SelectById mocks base method.
func (m *MockITransactionRepository) SelectById(id uint64) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockITransactionRepositoryMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockITransactionRepository)(nil).SelectById), id)
}

// SelectFoundrisingPhilantropIds mocks base method.
func (m *MockITransactionRepository) SelectFoundrisingPhilantropIds(foundrising_id uint64) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFoundrisingPhilantropIds", foundrising_id)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFoundrisingPhilantropIds indicates an expected call of SelectFoundrisingPhilantropIds.
func (mr *MockITransactionRepositoryMockRecorder) SelectFoundrisingPhilantropIds(foundrising_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFoundrisingPhilantropIds", reflect.TypeOf((*MockITransactionRepository)(nil).SelectFoundrisingPhilantropIds), foundrising_id)
}

// SelectFrom mocks base method.
func (m *MockITransactionRepository) SelectFrom(type_ bool, id uint64) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFrom", type_, id)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFrom indicates an expected call of SelectFrom.
func (mr *MockITransactionRepositoryMockRecorder) SelectFrom(type_, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFrom", reflect.TypeOf((*MockITransactionRepository)(nil).SelectFrom), type_, id)
}

// SelectTo mocks base method.
func (m *MockITransactionRepository) SelectTo(type_ bool, id uint64) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTo", type_, id)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTo indicates an expected call of SelectTo.
func (mr *MockITransactionRepositoryMockRecorder) SelectTo(type_, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTo", reflect.TypeOf((*MockITransactionRepository)(nil).SelectTo), type_, id)
}

// MockExistByIdChecker is a mock of ExistByIdChecker interface.
type MockExistByIdChecker struct {
	ctrl     *gomock.Controller
	recorder *MockExistByIdCheckerMockRecorder
}

// MockExistByIdCheckerMockRecorder is the mock recorder for MockExistByIdChecker.
type MockExistByIdCheckerMockRecorder struct {
	mock *MockExistByIdChecker
}

// NewMockExistByIdChecker creates a new mock instance.
func NewMockExistByIdChecker(ctrl *gomock.Controller) *MockExistByIdChecker {
	mock := &MockExistByIdChecker{ctrl: ctrl}
	mock.recorder = &MockExistByIdCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExistByIdChecker) EXPECT() *MockExistByIdCheckerMockRecorder {
	return m.recorder
}

// ExistsById mocks base method.
func (m *MockExistByIdChecker) ExistsById(id uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsById", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsById indicates an expected call of ExistsById.
func (mr *MockExistByIdCheckerMockRecorder) ExistsById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsById", reflect.TypeOf((*MockExistByIdChecker)(nil).ExistsById), id)
}

// MockITransactionService is a mock of ITransactionService interface.
type MockITransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionServiceMockRecorder
}

// MockITransactionServiceMockRecorder is the mock recorder for MockITransactionService.
type MockITransactionServiceMockRecorder struct {
	mock *MockITransactionService
}

// NewMockITransactionService creates a new mock instance.
func NewMockITransactionService(ctrl *gomock.Controller) *MockITransactionService {
	mock := &MockITransactionService{ctrl: ctrl}
	mock.recorder = &MockITransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionService) EXPECT() *MockITransactionServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockITransactionService) Add(TPars entities.TransactionAdd) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", TPars)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockITransactionServiceMockRecorder) Add(TPars interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockITransactionService)(nil).Add), TPars)
}

// Delete mocks base method.
func (m *MockITransactionService) Delete(id_ string) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id_)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockITransactionServiceMockRecorder) Delete(id_ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITransactionService)(nil).Delete), id_)
}

// ExistsById mocks base method.
func (m *MockITransactionService) ExistsById(id uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsById", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsById indicates an expected call of ExistsById.
func (mr *MockITransactionServiceMockRecorder) ExistsById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsById", reflect.TypeOf((*MockITransactionService)(nil).ExistsById), id)
}

// GetAll mocks base method.
func (m *MockITransactionService) GetAll() ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockITransactionServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockITransactionService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockITransactionService) GetById(id_ string) (entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id_)
	ret0, _ := ret[0].(entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockITransactionServiceMockRecorder) GetById(id_ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockITransactionService)(nil).GetById), id_)
}

// GetFoundrisingPhilantropIds mocks base method.
func (m *MockITransactionService) GetFoundrisingPhilantropIds(id_ string, FndgS services.ExistByIdChecker) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoundrisingPhilantropIds", id_, FndgS)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoundrisingPhilantropIds indicates an expected call of GetFoundrisingPhilantropIds.
func (mr *MockITransactionServiceMockRecorder) GetFoundrisingPhilantropIds(id_, FndgS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoundrisingPhilantropIds", reflect.TypeOf((*MockITransactionService)(nil).GetFoundrisingPhilantropIds), id_, FndgS)
}

// GetFromId mocks base method.
func (m *MockITransactionService) GetFromId(type_ bool, id_ string, FndS, US services.ExistByIdChecker) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromId", type_, id_, FndS, US)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromId indicates an expected call of GetFromId.
func (mr *MockITransactionServiceMockRecorder) GetFromId(type_, id_, FndS, US interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromId", reflect.TypeOf((*MockITransactionService)(nil).GetFromId), type_, id_, FndS, US)
}

// GetRepo mocks base method.
func (m *MockITransactionService) GetRepo() services.ITransactionRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepo")
	ret0, _ := ret[0].(services.ITransactionRepository)
	return ret0
}

// GetRepo indicates an expected call of GetRepo.
func (mr *MockITransactionServiceMockRecorder) GetRepo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepo", reflect.TypeOf((*MockITransactionService)(nil).GetRepo))
}

// GetToId mocks base method.
func (m *MockITransactionService) GetToId(type_ bool, id_ string, FndS, FndgS services.ExistByIdChecker) ([]entities.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToId", type_, id_, FndS, FndgS)
	ret0, _ := ret[0].([]entities.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToId indicates an expected call of GetToId.
func (mr *MockITransactionServiceMockRecorder) GetToId(type_, id_, FndS, FndgS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToId", reflect.TypeOf((*MockITransactionService)(nil).GetToId), type_, id_, FndS, FndgS)
}