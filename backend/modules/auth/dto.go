package auth

type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterDto struct {
	Email    string `json:"email" form:"email" binding:"required,email,max=100"`
	Password string `json:"password" form:"password" binding:"required,max=100,min=8"`
	Name     string `json:"name" form:"name" binding:"required,max=100"`
	Status   string `json:"status" form:"status" binding:"required,max=150"`
}

type ResetPasswordChangeDTO struct {
	Secret   string `json:"secret" binding:"required,min=30,max=100"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type ResetPasswordDTO struct {
	Email string `json:"email" form:"email" binding:"required,email,max=100"`
}

type RefreshTokenDTO struct {
	Token string `json:"token" form:"token" binding:"required,jwt"`
}
