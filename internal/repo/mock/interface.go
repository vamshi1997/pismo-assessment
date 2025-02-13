// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vamshi1997/pismo-assessment/internal/model"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockIRepository) CreateAccount(account model.Account) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockIRepositoryMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIRepository)(nil).CreateAccount), account)
}

// CreateTransaction mocks base method.
func (m *MockIRepository) CreateTransaction(transaction model.Transaction) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", transaction)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockIRepositoryMockRecorder) CreateTransaction(transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockIRepository)(nil).CreateTransaction), transaction)
}

// GetAccount mocks base method.
func (m *MockIRepository) GetAccount(accountId uint) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", accountId)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockIRepositoryMockRecorder) GetAccount(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockIRepository)(nil).GetAccount), accountId)
}

// GetPreviousTransactions mocks base method.
func (m *MockIRepository) GetPreviousTransactions() ([]model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviousTransactions")
	ret0, _ := ret[0].([]model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreviousTransactions indicates an expected call of GetPreviousTransactions.
func (mr *MockIRepositoryMockRecorder) GetPreviousTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreviousTransactions", reflect.TypeOf((*MockIRepository)(nil).GetPreviousTransactions))
}

// UpdateTransactionBalance mocks base method.
func (m *MockIRepository) UpdateTransactionBalance(balance float64, transactionId uint) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionBalance", balance, transactionId)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransactionBalance indicates an expected call of UpdateTransactionBalance.
func (mr *MockIRepositoryMockRecorder) UpdateTransactionBalance(balance, transactionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionBalance", reflect.TypeOf((*MockIRepository)(nil).UpdateTransactionBalance), balance, transactionId)
}
