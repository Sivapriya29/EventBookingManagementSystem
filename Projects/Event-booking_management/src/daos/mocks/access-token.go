// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\SivapriyaAnbu\Documents\golang_training\golang_training_sivapriya\Projects\Event-booking_management\src\daos\access-token.go

// Package mock_daos is a generated GoMock package.
package mock_daos

import (
	models "event-booking/database/models"
	context "event-booking/utils/context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccessTokenDAO is a mock of AccessTokenDAO interface.
type MockAccessTokenDAO struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenDAOMockRecorder
}

// MockAccessTokenDAOMockRecorder is the mock recorder for MockAccessTokenDAO.
type MockAccessTokenDAOMockRecorder struct {
	mock *MockAccessTokenDAO
}

// NewMockAccessTokenDAO creates a new mock instance.
func NewMockAccessTokenDAO(ctrl *gomock.Controller) *MockAccessTokenDAO {
	mock := &MockAccessTokenDAO{ctrl: ctrl}
	mock.recorder = &MockAccessTokenDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenDAO) EXPECT() *MockAccessTokenDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccessTokenDAO) Create(ctx *context.Context, accessToken *models.AccessToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccessTokenDAOMockRecorder) Create(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccessTokenDAO)(nil).Create), ctx, accessToken)
}

// Delete mocks base method.
func (m *MockAccessTokenDAO) Delete(ctx *context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccessTokenDAOMockRecorder) Delete(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccessTokenDAO)(nil).Delete), ctx, token)
}

// Get mocks base method.
func (m *MockAccessTokenDAO) Get(ctx *context.Context, token string) (*models.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, token)
	ret0, _ := ret[0].(*models.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccessTokenDAOMockRecorder) Get(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccessTokenDAO)(nil).Get), ctx, token)
}

// Update mocks base method.
func (m *MockAccessTokenDAO) Update(ctx *context.Context, id string, updatedSchedule *models.AccessToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, updatedSchedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAccessTokenDAOMockRecorder) Update(ctx, id, updatedSchedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccessTokenDAO)(nil).Update), ctx, id, updatedSchedule)
}

// Upsert mocks base method.
func (m *MockAccessTokenDAO) Upsert(ctx *context.Context, accessToken *models.AccessToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockAccessTokenDAOMockRecorder) Upsert(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockAccessTokenDAO)(nil).Upsert), ctx, accessToken)
}