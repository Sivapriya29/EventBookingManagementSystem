package tests

import (
	"errors"
	"event-booking/constants"
	mock_daos "event-booking/daos/mocks"
	"event-booking/dtos"
	"event-booking/services/users"
	"event-booking/utils/context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accMock := mock_daos.NewMockUserDAO(ctrl)
	req := &dtos.RegisterReq{
		FirstName: "Test1",
		LastName:  "testing",
		Email:     "test1@gmail.com",
		Mobile:    "test2",
		Password:  "test3",
		Role:      "admin",
	}

	accMock.EXPECT().CheckEmailExists(gomock.Any(), req.Email, req.Role).Return(false, nil)
	accMock.EXPECT().CheckMobileExists(gomock.Any(), req.Mobile, req.Role).Return(false, nil)
	accMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	svc := users.New()
	svc.SetDAOs(accMock, nil, nil)

	err := svc.Register(&context.Context{}, req)
	if err != nil {
		t.Error("Unable to register. Err:", err)
		return
	}
}

func TestRegisterEmailAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accMock := mock_daos.NewMockUserDAO(ctrl)
	req := &dtos.RegisterReq{
		FirstName: "Test1",
		LastName:  "testing",
		Email:     "test1@gmail.com",
		Mobile:    "test2",
		Password:  "test3",
		Role:      "admin",
	}

	accMock.EXPECT().CheckEmailExists(gomock.Any(), req.Email, req.Role).Return(true, nil)

	svc := users.New()
	svc.SetDAOs(accMock, nil, nil)

	err := svc.Register(&context.Context{}, req)
	if err != constants.ErrEmailTaken {
		t.Error("Unable to register. Err:", err)
		return
	}
}

func TestRegisterMobileTaken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accMock := mock_daos.NewMockUserDAO(ctrl)
	req := &dtos.RegisterReq{
		FirstName: "Test1",
		LastName:  "testing",
		Email:     "test1@gmail.com",
		Mobile:    "test2",
		Password:  "test3",
		Role:      "admin",
	}

	accMock.EXPECT().CheckEmailExists(gomock.Any(), req.Email, req.Role).Return(false, nil)
	accMock.EXPECT().CheckMobileExists(gomock.Any(), req.Mobile, req.Role).Return(true, nil)

	svc := users.New()
	svc.SetDAOs(accMock, nil, nil)

	err := svc.Register(&context.Context{}, req)
	if err != constants.ErrMobileTaken {
		t.Error("Unable to register. Err:", err)
		return
	}
}

func TestRegisterDatabaseIssue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accMock := mock_daos.NewMockUserDAO(ctrl)
	req := &dtos.RegisterReq{
		FirstName: "Test1",
		LastName:  "testing",
		Email:     "test1@gmail.com",
		Mobile:    "test2",
		Password:  "test3",
		Role:      "admin",
	}

	var dbIssue = errors.New("database issue")

	accMock.EXPECT().CheckEmailExists(gomock.Any(), req.Email, req.Role).Return(false, nil)
	accMock.EXPECT().CheckMobileExists(gomock.Any(), req.Mobile, req.Role).Return(false, nil)
	accMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(dbIssue)

	svc := users.New()
	svc.SetDAOs(accMock, nil, nil)

	err := svc.Register(&context.Context{}, req)
	if err != dbIssue {
		t.Error("Unable to register. Err:", err)
		return
	}
}
