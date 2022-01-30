package auth

type ResetPasswordDTO struct {
	Email string `json:"email" form:"email" binding:"required,email,max=100"`
}

type ResetPasswordChangeDTO struct {
	Secret   string `json:"secret" binding:"required,min=30,max=100"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}
