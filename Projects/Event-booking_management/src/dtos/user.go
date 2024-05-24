package dtos

type User struct {
	ID         string `json:"id" binding:"required"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Mobile     string `json:"mobile" binding:"required"`
	Role       string `json:"role" binding:"required"`
}
