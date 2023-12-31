// Code generated by MockGen. DO NOT EDIT.
// Source: cart/internal/usecase/cart.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	model "github.com/Go-Marketplace/backend/cart/internal/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockICartUsecase is a mock of ICartUsecase interface.
type MockICartUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockICartUsecaseMockRecorder
}

// MockICartUsecaseMockRecorder is the mock recorder for MockICartUsecase.
type MockICartUsecaseMockRecorder struct {
	mock *MockICartUsecase
}

// NewMockICartUsecase creates a new mock instance.
func NewMockICartUsecase(ctrl *gomock.Controller) *MockICartUsecase {
	mock := &MockICartUsecase{ctrl: ctrl}
	mock.recorder = &MockICartUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICartUsecase) EXPECT() *MockICartUsecaseMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockICartUsecase) CreateCart(ctx context.Context, cart model.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", ctx, cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockICartUsecaseMockRecorder) CreateCart(ctx, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockICartUsecase)(nil).CreateCart), ctx, cart)
}

// CreateCartline mocks base method.
func (m *MockICartUsecase) CreateCartline(ctx context.Context, cartline *model.CartLine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartline", ctx, cartline)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCartline indicates an expected call of CreateCartline.
func (mr *MockICartUsecaseMockRecorder) CreateCartline(ctx, cartline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartline", reflect.TypeOf((*MockICartUsecase)(nil).CreateCartline), ctx, cartline)
}

// CreateCartlines mocks base method.
func (m *MockICartUsecase) CreateCartlines(ctx context.Context, cartlines []*model.CartLine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartlines", ctx, cartlines)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCartlines indicates an expected call of CreateCartlines.
func (mr *MockICartUsecaseMockRecorder) CreateCartlines(ctx, cartlines interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartlines", reflect.TypeOf((*MockICartUsecase)(nil).CreateCartlines), ctx, cartlines)
}

// DeleteCart mocks base method.
func (m *MockICartUsecase) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockICartUsecaseMockRecorder) DeleteCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockICartUsecase)(nil).DeleteCart), ctx, userID)
}

// DeleteCartCartlines mocks base method.
func (m *MockICartUsecase) DeleteCartCartlines(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartCartlines", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartCartlines indicates an expected call of DeleteCartCartlines.
func (mr *MockICartUsecaseMockRecorder) DeleteCartCartlines(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartCartlines", reflect.TypeOf((*MockICartUsecase)(nil).DeleteCartCartlines), ctx, userID)
}

// DeleteCartline mocks base method.
func (m *MockICartUsecase) DeleteCartline(ctx context.Context, userID, productID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartline", ctx, userID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartline indicates an expected call of DeleteCartline.
func (mr *MockICartUsecaseMockRecorder) DeleteCartline(ctx, userID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartline", reflect.TypeOf((*MockICartUsecase)(nil).DeleteCartline), ctx, userID, productID)
}

// DeleteProductCartlines mocks base method.
func (m *MockICartUsecase) DeleteProductCartlines(ctx context.Context, productID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductCartlines", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductCartlines indicates an expected call of DeleteProductCartlines.
func (mr *MockICartUsecaseMockRecorder) DeleteProductCartlines(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductCartlines", reflect.TypeOf((*MockICartUsecase)(nil).DeleteProductCartlines), ctx, productID)
}

// GetCartline mocks base method.
func (m *MockICartUsecase) GetCartline(ctx context.Context, userID, productID uuid.UUID) (*model.CartLine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartline", ctx, userID, productID)
	ret0, _ := ret[0].(*model.CartLine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartline indicates an expected call of GetCartline.
func (mr *MockICartUsecaseMockRecorder) GetCartline(ctx, userID, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartline", reflect.TypeOf((*MockICartUsecase)(nil).GetCartline), ctx, userID, productID)
}

// GetUserCart mocks base method.
func (m *MockICartUsecase) GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCart", ctx, userID)
	ret0, _ := ret[0].(*model.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCart indicates an expected call of GetUserCart.
func (mr *MockICartUsecaseMockRecorder) GetUserCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCart", reflect.TypeOf((*MockICartUsecase)(nil).GetUserCart), ctx, userID)
}

// UpdateCartline mocks base method.
func (m *MockICartUsecase) UpdateCartline(ctx context.Context, cartline model.CartLine) (*model.CartLine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartline", ctx, cartline)
	ret0, _ := ret[0].(*model.CartLine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartline indicates an expected call of UpdateCartline.
func (mr *MockICartUsecaseMockRecorder) UpdateCartline(ctx, cartline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartline", reflect.TypeOf((*MockICartUsecase)(nil).UpdateCartline), ctx, cartline)
}
