package users

import (
	"event-booking/config"
	"event-booking/constants"
	"event-booking/daos"
	"event-booking/database/models"
	"event-booking/dtos"
	"event-booking/utils/context"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	user         daos.UserDAO
	accessToken  daos.AccessTokenDAO
	refreshToken daos.RefreshTokenDAO
}

func New() *User {
	return &User{
		user:         daos.NewUser(),
		accessToken:  daos.NewAccessToken(),
		refreshToken: daos.NewRefreshToken(),
	}
}

func (u *User) SetDAOs(user daos.UserDAO, accessToken daos.AccessTokenDAO, refreshToken daos.RefreshTokenDAO) {
	u.user = user
	u.accessToken = accessToken
	u.refreshToken = refreshToken
}

func (u *User) userFromRegisterReq(req *dtos.RegisterReq) *models.User {
	return &models.User{
		ID:         uuid.New().String(),
		First_name: req.FirstName,
		Last_name:  req.LastName,
		Email:      req.Email,
		Password:   req.Password,
		Mobile:     req.Mobile,
		Role:       req.Role,
		Created_at: time.Now(),
	}
}

func (u *User) userModelToDTO(req *models.User) *dtos.User {
	return &dtos.User{
		ID:         req.ID,
		First_name: req.First_name,
		Last_name:  req.Last_name,
		Email:      req.Email,
		Mobile:     req.Mobile,
		Role:       req.Role,
	}
}

func (u *User) GetAccountWithToken(ctx *context.Context, token string) (*dtos.User, error) {
	at, err := u.accessToken.Get(ctx, token)
	if err != nil {
		log.Println("Unable to get access token.Err:", err)
		return nil, err
	}

	if at.ExpiresAt.Before(time.Now()) {
		return nil, constants.ErrAccessTokenExpired
	}
	user, err := u.user.Get(ctx, at.UserId)
	if err != nil {
		log.Println("Unable to get account. Err: ", err)
		return nil, err
	}

	return u.userModelToDTO(user), nil
}

func (u *User) GetAccessTokenFromRefreshToken(ctx *context.Context, token string) (string, error) {
	rt, err := u.refreshToken.Get(ctx, token)
	if err != nil {
		log.Println("Unable to get access token.Err:", err)
		return "", err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		return "", constants.ErrAccessTokenExpired
	}

	accessToken, _ := GetAccessAndRefreshToken(config.Conf.TokenSize)
	err = u.accessToken.Create(ctx, &models.AccessToken{
		Token:        accessToken,
		RefreshToken: rt.Token,
		UserId:       rt.UserId,
		ExpiresAt:    time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
	})
	if err != nil {
		log.Println("Unable to save access tokens. Err:", err)
		return "", err
	}
	return accessToken, nil
}

func (u *User) Register(ctx *context.Context, req *dtos.RegisterReq) error {
	user := u.userFromRegisterReq(req)
	// account, err := a.accountFromRegisterReq(req)
	// if err != nil {
	// 	return err
	// }

	if ok, _ := u.user.CheckEmailExists(ctx, user.Email, user.Role); ok {
		return constants.ErrEmailTaken
	}

	if ok, _ := u.user.CheckMobileExists(ctx, user.Mobile, user.Role); ok {
		return constants.ErrMobileTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("Unable to create password hash.Err:", err)
		return err
	}

	user.Password = string(hash)

	return u.user.Create(ctx, user)
}

func (u *User) Login(ctx *context.Context, req *dtos.LoginReq) (*dtos.LoginRes, error) {
	user, err := u.user.GetAccountForEmail(ctx, req.Email, req.Role)
	if err == gorm.ErrRecordNotFound {
		log.Println("No record found. Err:", err)
		return nil, constants.ErrInvalidEmailOrPassword
	}

	if err != nil {
		log.Println("Error retrieving account. Err:", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("Invalid credentials.Err:", err)
		return nil, constants.ErrInvalidEmailOrPasswordOrRole
	}

	accessToken, refreshToken := GetAccessAndRefreshToken(config.Conf.TokenSize)
	err = u.refreshToken.Create(ctx, &models.RefreshToken{
		Token:     refreshToken,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(time.Duration(config.Conf.RefreshTokenExpiry) * time.Hour),
	})
	if err != nil {
		log.Println("Unable to save refresh tokens. Err:", err)
		return nil, err
	}

	err = u.accessToken.Create(ctx, &models.AccessToken{
		Token:        accessToken,
		RefreshToken: refreshToken,
		UserId:       user.ID,
		ExpiresAt:    time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
	})
	if err != nil {
		log.Println("Unable to save access tokens. Err:", err)
		return nil, err
	}

	// Login successful
	return &dtos.LoginRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
