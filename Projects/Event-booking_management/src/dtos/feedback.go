package dtos

type FeedbackReq struct {
	User_id  string  `json:"user_id" binding:"required"`
	Event_id string  `json:"event_id" binding:"required"`
	Rating   float64 `json:"rating" binding:"required"`
	Comments string  `json:"comments" binding:"required"`
}

type FeedbackRes struct {
	ID       string  `json:"id"`
	User_id  string  `json:"user_id"`
	Event_id string  `json:"event_id"`
	Rating   float64 `json:"rating"`
	Comments string  `json:"comments"`
}
