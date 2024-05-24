package daos

import (
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"
)

type UserDAO interface {
	Create(ctx *context.Context, user *models.User) error
	Upsert(ctx *context.Context, user *models.User) error
	Get(ctx *context.Context, id string) (*models.User, error)
	Delete(ctx *context.Context, id string) error
	CheckEmailExists(ctx *context.Context, email, role string) (bool, error)
	CheckMobileExists(ctx *context.Context, mobile, role string) (bool, error)
	//GetByEmail(ctx *context.Context, email string) (*models.User, error)
	GetAccountForEmail(ctx *context.Context, email, role string) (*models.User, error)
}

func NewUser() UserDAO {
	return &User{}
}

type User struct {
}

func (u *User) GetAccountForEmail(ctx *context.Context, email, role string) (*models.User, error) {
	user := &models.User{}
	err := ctx.DB.Table("users").Where("email=? AND role = ?", email, role).Scan(user).Error
	if err != nil {
		log.Println("Unable to get account. Err:", err)
		return nil, err
	}
	return user, nil
}

func (u *User) CheckEmailExists(ctx *context.Context, email, role string) (bool, error) {
	var cnt int
	err := ctx.DB.Table("users").Select("count(*)").Where("email=? AND role=?", email, role).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to create user. Err:", err)
		return false, err
	}
	return cnt > 0, nil
}

func (u *User) CheckMobileExists(ctx *context.Context, mobile, role string) (bool, error) {
	var cnt int
	err := ctx.DB.Table("users").Select("count(*)").Where("mobile=? AND role=?", mobile, role).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to create user. Err:", err)
		return false, err
	}
	return cnt > 0, nil
}

func (u *User) Create(ctx *context.Context, user *models.User) error {
	err := ctx.DB.Table("users").Create(user).Error
	if err != nil {
		log.Println("Unable to create user. Err:", err)
		return err
	}
	return nil
}

func (u *User) Upsert(ctx *context.Context, user *models.User) error {
	err := ctx.DB.Table("users").Save(user).Error
	if err != nil {
		log.Println("Unable to upsert user. Err:", err)
		return err
	}
	return nil
}

func (u *User) Get(ctx *context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := ctx.DB.Table("users").First(user, "id = ?", id).Error
	if err != nil {
		log.Println("Unable to get user. Err:", err)
		return nil, err
	}
	return user, nil
}

func (u *User) Delete(ctx *context.Context, id string) error {

	err := ctx.DB.Table("users").Delete(&models.User{
		ID: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete user. Err:", err)
		return err
	}
	return nil
}

// func (u *User) GetByEmail(ctx *context.Context, email string) (*models.User, error) {
// 	user := &models.User{}

// 	err := ctx.DB.Table("users").Where("email = ?", email).First(user).Error
// 	if err != nil {
// 		log.Println("Unable to retrieve account by email:", err)
// 		return nil, err
// 	}
// 	return user, nil
// }
