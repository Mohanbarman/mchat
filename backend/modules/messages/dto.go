package messages

import "mchat.com/api/lib/pagination"

type GetAllDTO struct {
	pagination.CursorPaginationDTO
}
