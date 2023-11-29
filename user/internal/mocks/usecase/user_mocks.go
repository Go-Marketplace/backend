// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/usecase/user.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	model "github.com/Go-Marketplace/backend/user/internal/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// ChangeUserRole mocks base method.
func (m *MockIUserUsecase) ChangeUserRole(ctx context.Context, user model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserRole", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUserRole indicates an expected call of ChangeUserRole.
func (mr *MockIUserUsecaseMockRecorder) ChangeUserRole(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserRole", reflect.TypeOf((*MockIUserUsecase)(nil).ChangeUserRole), ctx, user)
}

// CreateUser mocks base method.
func (m *MockIUserUsecase) CreateUser(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserUsecaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserUsecase)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockIUserUsecase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserUsecaseMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserUsecase)(nil).DeleteUser), ctx, id)
}

// GetAllUsers mocks base method.
func (m *MockIUserUsecase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserUsecaseMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserUsecase)(nil).GetAllUsers), ctx)
}

// GetUser mocks base method.
func (m *MockIUserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserUsecaseMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserUsecase)(nil).GetUser), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockIUserUsecase) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockIUserUsecaseMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockIUserUsecase)(nil).GetUserByEmail), ctx, email)
}

// UpdateUser mocks base method.
func (m *MockIUserUsecase) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserUsecaseMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserUsecase)(nil).UpdateUser), ctx, user)
}
