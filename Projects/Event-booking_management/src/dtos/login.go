package dtos

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginRes struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
