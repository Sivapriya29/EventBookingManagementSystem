package daos

import (
	"errors"
	"event-booking/database/models"
	"event-booking/utils/context"
	"log"

	"github.com/google/uuid"
)

type FeedbackDAO interface {
	Create(ctx *context.Context, feedback *models.Feedback) error
	Upsert(ctx *context.Context, feedback *models.Feedback) error
	Get(ctx *context.Context, id string) (*models.Feedback, error)
	Update(ctx *context.Context, id string, updatedBooking *models.Feedback) error
	Delete(ctx *context.Context, id string) error
	GetAllFeedbacks(ctx *context.Context) ([]*models.Feedback, error)
	CheckEventExists(ctx *context.Context, eventID string) (bool, error)
	CheckUserExists(ctx *context.Context, userID string) (bool, error)
	GetByEventID(ctx *context.Context, eventID string) ([]*models.Feedback, error)
	GetByUserID(ctx *context.Context, userID string) ([]*models.Feedback, error)
}

func NewFeedback() FeedbackDAO {
	return &Feedback{}
}

type Feedback struct {
}

func (f *Feedback) Create(ctx *context.Context, feedback *models.Feedback) error {
	// Validate UUIDs
	if _, err := uuid.Parse(feedback.Event_id); err != nil {
		return errors.New("invalid event_id format")
	}
	if _, err := uuid.Parse(feedback.User_id); err != nil {
		return errors.New("invalid user_id format")
	}

	// Check if event_id exists
	eventExists, err := f.CheckEventExists(ctx, feedback.Event_id)
	if err != nil {
		return err
	}
	if !eventExists {
		return errors.New("event_id not found")
	}

	// Check if user_id exists
	userExists, err := f.CheckUserExists(ctx, feedback.User_id)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user_id not found")
	}

	err = ctx.DB.Table("feedbacks").Create(feedback).Error
	if err != nil {
		log.Println("Unable to create feedback. Err:", err)
		return err
	}
	return nil
}

func (f *Feedback) Upsert(ctx *context.Context, feedback *models.Feedback) error {
	err := ctx.DB.Table("feedbacks").Save(feedback).Error
	if err != nil {
		log.Println("Unable to upsert feedback. Err:", err)
		return err
	}
	return nil
}

func (f *Feedback) Get(ctx *context.Context, id string) (*models.Feedback, error) {
	feedback := &models.Feedback{}
	err := ctx.DB.Table("feedbacks").First(feedback, "id = ?", id).Error
	if err != nil {
		log.Println("Unable to get feedback. Err:", err)
		return nil, err
	}
	return feedback, nil
}

func (f *Feedback) Update(ctx *context.Context, id string, updatedBooking *models.Feedback) error {
	err := ctx.DB.Table("feedbacks").Where("id = ?", id).Updates(updatedBooking).Error
	if err != nil {
		log.Println("Unable to update feedback.Err:", err)
		return err
	}
	return nil
}

func (f *Feedback) Delete(ctx *context.Context, id string) error {

	err := ctx.DB.Table("feedbacks").Delete(&models.Feedback{
		ID: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete feedback. Err:", err)
		return err
	}
	return nil
}

func (f *Feedback) GetAllFeedbacks(ctx *context.Context) ([]*models.Feedback, error) {
	var feedbacks []*models.Feedback
	err := ctx.DB.Table("feedbacks").Find(&feedbacks).Error

	if err != nil {
		log.Println("Unable to get all feedbacks. Err: ", err)
		return nil, err
	}

	return feedbacks, nil
}

func (f *Feedback) CheckEventExists(ctx *context.Context, eventID string) (bool, error) {
	var count int64
	if err := ctx.DB.Model(&models.Event{}).Where("id = ?", eventID).Count(&count).Error; err != nil {
		log.Println("Unable to check event existence. Err:", err)
		return false, err
	}
	return count > 0, nil
}

func (f *Feedback) CheckUserExists(ctx *context.Context, userID string) (bool, error) {
	var count int64
	if err := ctx.DB.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		log.Println("Unable to check user existence. Err:", err)
		return false, err
	}
	return count > 0, nil
}

func (f *Feedback) GetByEventID(ctx *context.Context, eventID string) ([]*models.Feedback, error) {
	if _, err := uuid.Parse(eventID); err != nil {
		log.Println("invalid event_id format or event_id not found. Err:", err)
		return nil, err
	}

	eventExists, err := f.CheckEventExists(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if !eventExists {
		return nil, errors.New("event_id not found")
	}

	var feedbacks []*models.Feedback
	if err := ctx.DB.Table("feedbacks").Where("event_id = ?", eventID).Find(&feedbacks).Error; err != nil {
		log.Println("Unable to get feedbacks by event ID. Err:", err)
		return nil, err
	}
	return feedbacks, nil
}

func (f *Feedback) GetByUserID(ctx *context.Context, userID string) ([]*models.Feedback, error) {
	if _, err := uuid.Parse(userID); err != nil {
		log.Println("invalid user_id format or user_id not found. Err:", err)
		return nil, err
	}

	userExists, err := f.CheckUserExists(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, errors.New("user_id not found")
	}

	var feedbacks []*models.Feedback
	if err := ctx.DB.Where("user_id = ?", userID).Find(&feedbacks).Error; err != nil {
		log.Println("Unable to get feedbacks by user ID. Err:", err)
		return nil, err
	}
	return feedbacks, nil
}
