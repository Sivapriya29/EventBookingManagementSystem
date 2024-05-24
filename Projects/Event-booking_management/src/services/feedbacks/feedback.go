package feedbacks

import (
	"event-booking/daos"
	"event-booking/database/models"
	"event-booking/dtos"
	"event-booking/utils/context"
	"log"

	"github.com/google/uuid"
)

type Feedback struct {
	feedback daos.FeedbackDAO
}

func New() *Feedback {
	return &Feedback{
		feedback: daos.NewFeedback(),
	}
}

func (f *Feedback) CreateFeedback(ctx *context.Context, req *dtos.FeedbackReq) error {
	feedback := &models.Feedback{
		ID:       uuid.New().String(),
		User_id:  req.User_id,
		Event_id: req.Event_id,
		Rating:   req.Rating,
		Comments: req.Comments,
	}
	return f.feedback.Create(ctx, feedback)
}

func (f *Feedback) GetFeedback(ctx *context.Context, id string) (*dtos.FeedbackRes, error) {
	feedback, err := f.feedback.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dtos.FeedbackRes{
		ID:       feedback.ID,
		User_id:  feedback.User_id,
		Event_id: feedback.Event_id,
		Rating:   feedback.Rating,
		Comments: feedback.Comments,
	}, nil
}

func (f *Feedback) DeleteFeedback(ctx *context.Context, id string) error {
	return f.feedback.Delete(ctx, id)
}

func (f *Feedback) GetAllFeedbacks(ctx *context.Context) ([]*dtos.FeedbackRes, error) {
	feedbacks, err := f.feedback.GetAllFeedbacks(ctx)
	if err != nil {
		log.Println("Unable to fetch all feedbacks. Err: ", err)
		return nil, err
	}

	var result []*dtos.FeedbackRes
	for _, feedback := range feedbacks {
		result = append(result, &dtos.FeedbackRes{
			ID:       feedback.ID,
			User_id:  feedback.User_id,
			Event_id: feedback.Event_id,
			Rating:   feedback.Rating,
			Comments: feedback.Comments,
		})
	}
	return result, nil
}

func (f *Feedback) GetFeedbacksByEventID(ctx *context.Context, eventID string) ([]*dtos.FeedbackRes, error) {
	feedbacks, err := f.feedback.GetByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	var result []*dtos.FeedbackRes
	for _, feedback := range feedbacks {
		result = append(result, &dtos.FeedbackRes{
			ID:       feedback.ID,
			User_id:  feedback.User_id,
			Event_id: feedback.Event_id,
			Rating:   feedback.Rating,
			Comments: feedback.Comments,
		})
	}
	return result, nil
}

func (f *Feedback) GetFeedbacksByUserID(ctx *context.Context, userID string) ([]*dtos.FeedbackRes, error) {
	feedbacks, err := f.feedback.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []*dtos.FeedbackRes
	for _, feedback := range feedbacks {
		result = append(result, &dtos.FeedbackRes{
			ID:       feedback.ID,
			User_id:  feedback.User_id,
			Event_id: feedback.Event_id,
			Rating:   feedback.Rating,
			Comments: feedback.Comments,
		})
	}
	return result, nil
}
