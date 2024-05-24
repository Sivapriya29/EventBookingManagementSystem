package daos

import (
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"
)

type AccessTokenDAO interface {
	Create(ctx *context.Context, accessToken *models.AccessToken) error
	Upsert(ctx *context.Context, accessToken *models.AccessToken) error
	Get(ctx *context.Context, token string) (*models.AccessToken, error)
	Update(ctx *context.Context, id string, updatedSchedule *models.AccessToken) error
	Delete(ctx *context.Context, token string) error
}

func NewAccessToken() AccessTokenDAO {
	return &AccessToken{}
}

type AccessToken struct {
}

func (t *AccessToken) Create(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_tokens").Create(accessToken).Error
	if err != nil {
		log.Println("Unable to create accessToken. Err:", err)
		return err
	}
	return nil
}

func (t *AccessToken) Upsert(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_tokens").Save(accessToken).Error
	if err != nil {
		log.Println("Unable to upsert accessToken. Err:", err)
		return err
	}
	return nil
}

func (t *AccessToken) Get(ctx *context.Context, token string) (*models.AccessToken, error) {
	accessToken := &models.AccessToken{}
	err := ctx.DB.Table("access_tokens").First(accessToken, "token = ?", token).Error
	if err != nil {
		log.Println("Unable to get AccessToken. Err:", err)
		return nil, err
	}
	return accessToken, nil
}

func (t *AccessToken) Update(ctx *context.Context, id string, updatedSchedule *models.AccessToken) error {
	err := ctx.DB.Table("access_tokens").Where("id = ?", id).Updates(updatedSchedule).Error
	if err != nil {
		log.Println("Unable to update accessToken.Err:", err)
		return err
	}
	return nil
}

func (t *AccessToken) Delete(ctx *context.Context, token string) error {

	err := ctx.DB.Table("access_tokens").Delete(&models.AccessToken{
		Token: token,
	}).Error
	if err != nil {
		log.Println("Unable to delete accessToken. Err:", err)
		return err
	}
	return nil
}
