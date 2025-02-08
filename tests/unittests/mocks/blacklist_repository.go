// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/contracts/repository/blacklist_interface.go

// Package mock_repositoty is a generated GoMock package.
package mock_repositoty

import (
	reflect "reflect"

	entity "github.com/GeovanniGomes/blacklist/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIBlackListRepository is a mock of IBlackListRepository interface.
type MockIBlackListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBlackListRepositoryMockRecorder
}

// MockIBlackListRepositoryMockRecorder is the mock recorder for MockIBlackListRepository.
type MockIBlackListRepositoryMockRecorder struct {
	mock *MockIBlackListRepository
}

// NewMockIBlackListRepository creates a new mock instance.
func NewMockIBlackListRepository(ctrl *gomock.Controller) *MockIBlackListRepository {
	mock := &MockIBlackListRepository{ctrl: ctrl}
	mock.recorder = &MockIBlackListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBlackListRepository) EXPECT() *MockIBlackListRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIBlackListRepository) Add(blacklist *entity.BlackList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", blacklist)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIBlackListRepositoryMockRecorder) Add(blacklist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIBlackListRepository)(nil).Add), blacklist)
}

// Check mocks base method.
func (m *MockIBlackListRepository) Check(userIndentifier int, evendId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", userIndentifier, evendId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockIBlackListRepositoryMockRecorder) Check(userIndentifier, evendId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIBlackListRepository)(nil).Check), userIndentifier, evendId)
}

// Remove mocks base method.
func (m *MockIBlackListRepository) Remove(userIndentifier int, eventId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", userIndentifier, eventId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockIBlackListRepositoryMockRecorder) Remove(userIndentifier, eventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIBlackListRepository)(nil).Remove), userIndentifier, eventId)
}
