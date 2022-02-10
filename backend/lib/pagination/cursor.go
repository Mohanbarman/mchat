package pagination

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type CursorPaginationDTO struct {
	Cursor string `json:"cursor" form:"cursor"`
	Limit  int    `json:"limit" form:"limit" binding:"required"`
}

// used to store records required for pagination
type records struct {
	CreatedAt time.Time
}

// errors
const (
	InvalidCursorErr = iota + 1
)

func CursorPaginate(tableName string, err *int, dto *CursorPaginationDTO, meta *map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rows := []records{}
		cursor, e := decodeCursor(dto.Cursor)

		if e != nil {
			(*err) = InvalidCursorErr
			return db
		}

		(*meta)["next"] = nil
		(*meta)["limit"] = dto.Limit

		if len(dto.Cursor) > 0 {
			icursor, e := strconv.ParseInt(cursor, 10, 64)
			if e != nil {
				(*err) = InvalidCursorErr
				return db
			}
			t := time.UnixMicro(icursor).UTC()
			db.Session(&gorm.Session{}).Table(tableName).Where("created_at >= ?", t).Limit(dto.Limit + 1).Scan(&rows)
			if len(rows) > dto.Limit {
				(*meta)["next"] = encodeCursor(rows[dto.Limit].CreatedAt.UTC().UnixMicro())
			}
			return db.Where("created_at >= ?", t).Limit(dto.Limit)
		}

		db.Session(&gorm.Session{}).Table(tableName).Limit(dto.Limit + 1).Scan(&rows)
		if len(rows) > dto.Limit {
			(*meta)["next"] = encodeCursor(rows[dto.Limit].CreatedAt.UTC().UnixMicro())
		}

		return db.Limit(dto.Limit)
	}
}

// encoding cursor to base64
func encodeCursor(cursor int64) string {
	e := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", cursor)))
	return e
}

// decoding base64 cursor
func decodeCursor(cursor string) (dcoded string, err error) {
	d, err := base64.StdEncoding.DecodeString(cursor)
	dcoded = string(d)
	return
}
