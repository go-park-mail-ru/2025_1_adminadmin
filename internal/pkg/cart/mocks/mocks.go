// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/cart/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/uuid"
)

// MockCartRepo is a mock of CartRepo interface.
type MockCartRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepoMockRecorder
}

// MockCartRepoMockRecorder is the mock recorder for MockCartRepo.
type MockCartRepoMockRecorder struct {
	mock *MockCartRepo
}

// NewMockCartRepo creates a new mock instance.
func NewMockCartRepo(ctrl *gomock.Controller) *MockCartRepo {
	mock := &MockCartRepo{ctrl: ctrl}
	mock.recorder = &MockCartRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepo) EXPECT() *MockCartRepoMockRecorder {
	return m.recorder
}

// ClearCart mocks base method.
func (m *MockCartRepo) ClearCart(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearCart", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearCart indicates an expected call of ClearCart.
func (mr *MockCartRepoMockRecorder) ClearCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCart", reflect.TypeOf((*MockCartRepo)(nil).ClearCart), ctx, userID)
}

// GetCart mocks base method.
func (m *MockCartRepo) GetCart(ctx context.Context, userID string) (map[string]int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", ctx, userID)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCart indicates an expected call of GetCart.
func (mr *MockCartRepoMockRecorder) GetCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockCartRepo)(nil).GetCart), ctx, userID)
}

// UpdateItemQuantity mocks base method.
func (m *MockCartRepo) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantId string, quantity int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemQuantity", ctx, userID, productID, restaurantId, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemQuantity indicates an expected call of UpdateItemQuantity.
func (mr *MockCartRepoMockRecorder) UpdateItemQuantity(ctx, userID, productID, restaurantId, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemQuantity", reflect.TypeOf((*MockCartRepo)(nil).UpdateItemQuantity), ctx, userID, productID, restaurantId, quantity)
}

// MockCartUsecase is a mock of CartUsecase interface.
type MockCartUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCartUsecaseMockRecorder
}

// MockCartUsecaseMockRecorder is the mock recorder for MockCartUsecase.
type MockCartUsecaseMockRecorder struct {
	mock *MockCartUsecase
}

// NewMockCartUsecase creates a new mock instance.
func NewMockCartUsecase(ctrl *gomock.Controller) *MockCartUsecase {
	mock := &MockCartUsecase{ctrl: ctrl}
	mock.recorder = &MockCartUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUsecase) EXPECT() *MockCartUsecaseMockRecorder {
	return m.recorder
}

// ClearCart mocks base method.
func (m *MockCartUsecase) ClearCart(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearCart", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearCart indicates an expected call of ClearCart.
func (mr *MockCartUsecaseMockRecorder) ClearCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCart", reflect.TypeOf((*MockCartUsecase)(nil).ClearCart), ctx, userID)
}

// CreateOrder mocks base method.
func (m *MockCartUsecase) CreateOrder(ctx context.Context, userID string, details models.OrderInReq, cart models.Cart) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, userID, details, cart)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockCartUsecaseMockRecorder) CreateOrder(ctx, userID, details, cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockCartUsecase)(nil).CreateOrder), ctx, userID, details, cart)
}

// GetCart mocks base method.
func (m *MockCartUsecase) GetCart(ctx context.Context, userID string) (models.Cart, error, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", ctx, userID)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetCart indicates an expected call of GetCart.
func (mr *MockCartUsecaseMockRecorder) GetCart(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockCartUsecase)(nil).GetCart), ctx, userID)
}

// GetOrderById mocks base method.
func (m *MockCartUsecase) GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", ctx, order_id, user_id)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockCartUsecaseMockRecorder) GetOrderById(ctx, order_id, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockCartUsecase)(nil).GetOrderById), ctx, order_id, user_id)
}

// GetOrders mocks base method.
func (m *MockCartUsecase) GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx, user_id, count, offset)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockCartUsecaseMockRecorder) GetOrders(ctx, user_id, count, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockCartUsecase)(nil).GetOrders), ctx, user_id, count, offset)
}

// UpdateItemQuantity mocks base method.
func (m *MockCartUsecase) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantId string, quantity int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemQuantity", ctx, userID, productID, restaurantId, quantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemQuantity indicates an expected call of UpdateItemQuantity.
func (mr *MockCartUsecaseMockRecorder) UpdateItemQuantity(ctx, userID, productID, restaurantId, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemQuantity", reflect.TypeOf((*MockCartUsecase)(nil).UpdateItemQuantity), ctx, userID, productID, restaurantId, quantity)
}

// UpdateOrderStatus mocks base method.
func (m *MockCartUsecase) UpdateOrderStatus(ctx context.Context, order_id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", ctx, order_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockCartUsecaseMockRecorder) UpdateOrderStatus(ctx, order_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockCartUsecase)(nil).UpdateOrderStatus), ctx, order_id)
}

// MockRestaurantRepo is a mock of RestaurantRepo interface.
type MockRestaurantRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantRepoMockRecorder
}

// MockRestaurantRepoMockRecorder is the mock recorder for MockRestaurantRepo.
type MockRestaurantRepoMockRecorder struct {
	mock *MockRestaurantRepo
}

// NewMockRestaurantRepo creates a new mock instance.
func NewMockRestaurantRepo(ctrl *gomock.Controller) *MockRestaurantRepo {
	mock := &MockRestaurantRepo{ctrl: ctrl}
	mock.recorder = &MockRestaurantRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantRepo) EXPECT() *MockRestaurantRepoMockRecorder {
	return m.recorder
}

// GetCartItem mocks base method.
func (m *MockRestaurantRepo) GetCartItem(ctx context.Context, productIDs []string, productAmounts map[string]int, restaurantID string) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItem", ctx, productIDs, productAmounts, restaurantID)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItem indicates an expected call of GetCartItem.
func (mr *MockRestaurantRepoMockRecorder) GetCartItem(ctx, productIDs, productAmounts, restaurantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItem", reflect.TypeOf((*MockRestaurantRepo)(nil).GetCartItem), ctx, productIDs, productAmounts, restaurantID)
}

// GetOrderById mocks base method.
func (m *MockRestaurantRepo) GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", ctx, order_id, user_id)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockRestaurantRepoMockRecorder) GetOrderById(ctx, order_id, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockRestaurantRepo)(nil).GetOrderById), ctx, order_id, user_id)
}

// GetOrders mocks base method.
func (m *MockRestaurantRepo) GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx, user_id, count, offset)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockRestaurantRepoMockRecorder) GetOrders(ctx, user_id, count, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockRestaurantRepo)(nil).GetOrders), ctx, user_id, count, offset)
}

// Save mocks base method.
func (m *MockRestaurantRepo) Save(ctx context.Context, order models.Order, userLogin string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, order, userLogin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRestaurantRepoMockRecorder) Save(ctx, order, userLogin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRestaurantRepo)(nil).Save), ctx, order, userLogin)
}

// ScheduleDeliveryStatusChange mocks base method.
func (m *MockRestaurantRepo) ScheduleDeliveryStatusChange(ctx context.Context, orderID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleDeliveryStatusChange", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleDeliveryStatusChange indicates an expected call of ScheduleDeliveryStatusChange.
func (mr *MockRestaurantRepoMockRecorder) ScheduleDeliveryStatusChange(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleDeliveryStatusChange", reflect.TypeOf((*MockRestaurantRepo)(nil).ScheduleDeliveryStatusChange), ctx, orderID)
}

// UpdateOrderStatus mocks base method.
func (m *MockRestaurantRepo) UpdateOrderStatus(ctx context.Context, order_id uuid.UUID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", ctx, order_id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockRestaurantRepoMockRecorder) UpdateOrderStatus(ctx, order_id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockRestaurantRepo)(nil).UpdateOrderStatus), ctx, order_id, status)
}
