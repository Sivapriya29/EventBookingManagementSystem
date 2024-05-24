package daos

import (
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"
)

type RefreshTokenDAO interface {
	Create(ctx *context.Context, refreshToken *models.RefreshToken) error
	Upsert(ctx *context.Context, refreshToken *models.RefreshToken) error
	Get(ctx *context.Context, token string) (*models.RefreshToken, error)
	Update(ctx *context.Context, id string, updatedSchedule *models.RefreshToken) error
	Delete(ctx *context.Context, token string) error
}

func NewRefreshToken() RefreshTokenDAO {
	return &RefreshToken{}
}

type RefreshToken struct {
}

func (t *RefreshToken) Create(ctx *context.Context, refreshToken *models.RefreshToken) error {
	err := ctx.DB.Table("refresh_tokens").Create(refreshToken).Error
	if err != nil {
		log.Println("Unable to create refreshToken. Err:", err)
		return err
	}
	return nil
}

func (t *RefreshToken) Upsert(ctx *context.Context, refreshToken *models.RefreshToken) error {
	err := ctx.DB.Table("refresh_tokens").Save(refreshToken).Error
	if err != nil {
		log.Println("Unable to upsert refreshToken. Err:", err)
		return err
	}
	return nil
}

func (t *RefreshToken) Get(ctx *context.Context, token string) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{}
	err := ctx.DB.Table("refresh_tokens").First(refreshToken, "token = ?", token).Error
	if err != nil {
		log.Println("Unable to get RefreshToken. Err:", err)
		return nil, err
	}
	return refreshToken, nil
}

func (t *RefreshToken) Update(ctx *context.Context, id string, updatedSchedule *models.RefreshToken) error {
	err := ctx.DB.Table("refresh_tokens").Where("id = ?", id).Updates(updatedSchedule).Error
	if err != nil {
		log.Println("Unable to update refreshToken.Err:", err)
		return err
	}
	return nil
}

func (t *RefreshToken) Delete(ctx *context.Context, token string) error {

	err := ctx.DB.Table("refresh_tokens").Delete(&models.RefreshToken{
		Token: token,
	}).Error
	if err != nil {
		log.Println("Unable to delete refreshToken. Err:", err)
		return err
	}
	return nil
}
