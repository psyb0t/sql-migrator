// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/psyb0t/sql-migrator/internal/pkg/migrate (interfaces: Migrator)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/migrate.gen.go -package=mocks . Migrator
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMigrator is a mock of Migrator interface.
type MockMigrator struct {
	ctrl     *gomock.Controller
	recorder *MockMigratorMockRecorder
}

// MockMigratorMockRecorder is the mock recorder for MockMigrator.
type MockMigratorMockRecorder struct {
	mock *MockMigrator
}

// NewMockMigrator creates a new mock instance.
func NewMockMigrator(ctrl *gomock.Controller) *MockMigrator {
	mock := &MockMigrator{ctrl: ctrl}
	mock.recorder = &MockMigratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrator) EXPECT() *MockMigratorMockRecorder {
	return m.recorder
}

// Down mocks base method.
func (m *MockMigrator) Down() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down")
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down.
func (mr *MockMigratorMockRecorder) Down() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockMigrator)(nil).Down))
}

// Up mocks base method.
func (m *MockMigrator) Up() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up.
func (mr *MockMigratorMockRecorder) Up() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockMigrator)(nil).Up))
}