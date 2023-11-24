// Code generated by MockGen. DO NOT EDIT.
// Source: order/internal/infrastructure/interfaces/order.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	model "github.com/Go-Marketplace/backend/order/internal/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockOrderRepo is a mock of OrderRepo interface.
type MockOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoMockRecorder
}

// MockOrderRepoMockRecorder is the mock recorder for MockOrderRepo.
type MockOrderRepoMockRecorder struct {
	mock *MockOrderRepo
}

// NewMockOrderRepo creates a new mock instance.
func NewMockOrderRepo(ctrl *gomock.Controller) *MockOrderRepo {
	mock := &MockOrderRepo{ctrl: ctrl}
	mock.recorder = &MockOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepo) EXPECT() *MockOrderRepoMockRecorder {
	return m.recorder
}

// CancelOrder mocks base method.
func (m *MockOrderRepo) CancelOrder(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockOrderRepoMockRecorder) CancelOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockOrderRepo)(nil).CancelOrder), ctx, id)
}

// CreateOrder mocks base method.
func (m *MockOrderRepo) CreateOrder(ctx context.Context, order model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepoMockRecorder) CreateOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepo)(nil).CreateOrder), ctx, order)
}

// GetAllUserOrders mocks base method.
func (m *MockOrderRepo) GetAllUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserOrders", ctx, userID)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserOrders indicates an expected call of GetAllUserOrders.
func (mr *MockOrderRepoMockRecorder) GetAllUserOrders(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserOrders", reflect.TypeOf((*MockOrderRepo)(nil).GetAllUserOrders), ctx, userID)
}

// GetOrder mocks base method.
func (m *MockOrderRepo) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderRepoMockRecorder) GetOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderRepo)(nil).GetOrder), ctx, id)
}