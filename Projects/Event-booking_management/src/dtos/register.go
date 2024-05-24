package dtos

type RegisterReq struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	Role      string `json:"role" binding:"required"`
}
