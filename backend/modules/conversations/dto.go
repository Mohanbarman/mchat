package conversations

import "mchat.com/api/lib"

type GetAllDTO struct {
	lib.CursorPaginationDTO
}

type CreateDTO struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}
