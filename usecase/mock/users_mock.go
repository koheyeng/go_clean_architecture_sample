// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/users.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/koheyeng/go_clean_architecture_sample/domain"
	usecase "github.com/koheyeng/go_clean_architecture_sample/usecase"
	reflect "reflect"
)

// MockUsersInputPort is a mock of UsersInputPort interface
type MockUsersInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockUsersInputPortMockRecorder
}

// MockUsersInputPortMockRecorder is the mock recorder for MockUsersInputPort
type MockUsersInputPortMockRecorder struct {
	mock *MockUsersInputPort
}

// NewMockUsersInputPort creates a new mock instance
func NewMockUsersInputPort(ctrl *gomock.Controller) *MockUsersInputPort {
	mock := &MockUsersInputPort{ctrl: ctrl}
	mock.recorder = &MockUsersInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersInputPort) EXPECT() *MockUsersInputPortMockRecorder {
	return m.recorder
}

// GetUserAge mocks base method
func (m *MockUsersInputPort) GetUserAge(arg0 usecase.UsersDto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAge", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserAge indicates an expected call of GetUserAge
func (mr *MockUsersInputPortMockRecorder) GetUserAge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAge", reflect.TypeOf((*MockUsersInputPort)(nil).GetUserAge), arg0)
}

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// GetUser mocks base method
func (m *MockUsersRepository) GetUser(arg0 usecase.UsersDto) (*domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUsersRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsersRepository)(nil).GetUser), arg0)
}

// MockOuterSystemGateway is a mock of OuterSystemGateway interface
type MockOuterSystemGateway struct {
	ctrl     *gomock.Controller
	recorder *MockOuterSystemGatewayMockRecorder
}

// MockOuterSystemGatewayMockRecorder is the mock recorder for MockOuterSystemGateway
type MockOuterSystemGatewayMockRecorder struct {
	mock *MockOuterSystemGateway
}

// NewMockOuterSystemGateway creates a new mock instance
func NewMockOuterSystemGateway(ctrl *gomock.Controller) *MockOuterSystemGateway {
	mock := &MockOuterSystemGateway{ctrl: ctrl}
	mock.recorder = &MockOuterSystemGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOuterSystemGateway) EXPECT() *MockOuterSystemGatewayMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockOuterSystemGateway) Send(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockOuterSystemGatewayMockRecorder) Send(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOuterSystemGateway)(nil).Send), userID)
}

// MockUsersPresenter is a mock of UsersPresenter interface
type MockUsersPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockUsersPresenterMockRecorder
}

// MockUsersPresenterMockRecorder is the mock recorder for MockUsersPresenter
type MockUsersPresenterMockRecorder struct {
	mock *MockUsersPresenter
}

// NewMockUsersPresenter creates a new mock instance
func NewMockUsersPresenter(ctrl *gomock.Controller) *MockUsersPresenter {
	mock := &MockUsersPresenter{ctrl: ctrl}
	mock.recorder = &MockUsersPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsersPresenter) EXPECT() *MockUsersPresenterMockRecorder {
	return m.recorder
}

// RespUserAge mocks base method
func (m *MockUsersPresenter) RespUserAge(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespUserAge", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespUserAge indicates an expected call of RespUserAge
func (mr *MockUsersPresenterMockRecorder) RespUserAge(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespUserAge", reflect.TypeOf((*MockUsersPresenter)(nil).RespUserAge), arg0)
}
