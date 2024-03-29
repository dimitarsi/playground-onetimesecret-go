// Code generated by MockGen. DO NOT EDIT.
// Source: secret.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSecretRepository is a mock of SecretRepository interface.
type MockSecretRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSecretRepositoryMockRecorder
}

// MockSecretRepositoryMockRecorder is the mock recorder for MockSecretRepository.
type MockSecretRepositoryMockRecorder struct {
	mock *MockSecretRepository
}

// NewMockSecretRepository creates a new mock instance.
func NewMockSecretRepository(ctrl *gomock.Controller) *MockSecretRepository {
	mock := &MockSecretRepository{ctrl: ctrl}
	mock.recorder = &MockSecretRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretRepository) EXPECT() *MockSecretRepositoryMockRecorder {
	return m.recorder
}

// GetDel mocks base method.
func (m *MockSecretRepository) GetDel(key string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", key)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDel indicates an expected call of GetDel.
func (mr *MockSecretRepositoryMockRecorder) GetDel(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockSecretRepository)(nil).GetDel), key)
}

// Set mocks base method.
func (m *MockSecretRepository) Set(key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSecretRepositoryMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSecretRepository)(nil).Set), key, value)
}
