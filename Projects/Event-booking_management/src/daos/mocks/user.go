// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\SivapriyaAnbu\Documents\golang_training\golang_training_sivapriya\Projects\Event-booking_management\src\daos\user.go

// Package mock_daos is a generated GoMock package.
package mock_daos

import (
	models "event-booking/database/models"
	context "event-booking/utils/context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserDAO is a mock of UserDAO interface.
type MockUserDAO struct {
	ctrl     *gomock.Controller
	recorder *MockUserDAOMockRecorder
}

// MockUserDAOMockRecorder is the mock recorder for MockUserDAO.
type MockUserDAOMockRecorder struct {
	mock *MockUserDAO
}

// NewMockUserDAO creates a new mock instance.
func NewMockUserDAO(ctrl *gomock.Controller) *MockUserDAO {
	mock := &MockUserDAO{ctrl: ctrl}
	mock.recorder = &MockUserDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDAO) EXPECT() *MockUserDAOMockRecorder {
	return m.recorder
}

// CheckEmailExists mocks base method.
func (m *MockUserDAO) CheckEmailExists(ctx *context.Context, email, role string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailExists", ctx, email, role)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailExists indicates an expected call of CheckEmailExists.
func (mr *MockUserDAOMockRecorder) CheckEmailExists(ctx, email, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailExists", reflect.TypeOf((*MockUserDAO)(nil).CheckEmailExists), ctx, email, role)
}

// CheckMobileExists mocks base method.
func (m *MockUserDAO) CheckMobileExists(ctx *context.Context, mobile, role string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMobileExists", ctx, mobile, role)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMobileExists indicates an expected call of CheckMobileExists.
func (mr *MockUserDAOMockRecorder) CheckMobileExists(ctx, mobile, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMobileExists", reflect.TypeOf((*MockUserDAO)(nil).CheckMobileExists), ctx, mobile, role)
}

// Create mocks base method.
func (m *MockUserDAO) Create(ctx *context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserDAOMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDAO)(nil).Create), ctx, user)
}

// Delete mocks base method.
func (m *MockUserDAO) Delete(ctx *context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserDAOMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDAO)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockUserDAO) Get(ctx *context.Context, id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserDAOMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserDAO)(nil).Get), ctx, id)
}

// GetAccountForEmail mocks base method.
func (m *MockUserDAO) GetAccountForEmail(ctx *context.Context, email, role string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForEmail", ctx, email, role)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForEmail indicates an expected call of GetAccountForEmail.
func (mr *MockUserDAOMockRecorder) GetAccountForEmail(ctx, email, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForEmail", reflect.TypeOf((*MockUserDAO)(nil).GetAccountForEmail), ctx, email, role)
}

// Upsert mocks base method.
func (m *MockUserDAO) Upsert(ctx *context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockUserDAOMockRecorder) Upsert(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockUserDAO)(nil).Upsert), ctx, user)
}
