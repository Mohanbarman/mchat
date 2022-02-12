package pagination

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
	"mchat.com/api/lib"
)

type CursorPaginationDTO struct {
	Cursor string `json:"cursor" form:"cursor"`
	Limit  string `json:"limit" form:"limit" binding:"number"`
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
		limit := lib.MustGetInt(dto.Limit)

		if e != nil {
			(*err) = InvalidCursorErr
			return db
		}

		(*meta)["next"] = nil

		if len(dto.Cursor) > 0 {
			icursor, e := strconv.ParseInt(cursor, 10, 64)
			if e != nil {
				(*err) = InvalidCursorErr
				return db
			}
			t := time.UnixMicro(icursor).UTC()
			db.Session(&gorm.Session{}).Table(tableName).Where("created_at >= ?", t).Limit(limit + 1).Scan(&rows)
			if len(rows) > limit {
				(*meta)["next"] = encodeCursor(rows[limit].CreatedAt.UTC().UnixMicro())
			}
			return db.Where("created_at >= ?", t).Limit(limit)
		}

		db.Session(&gorm.Session{}).Table(tableName).Limit(limit + 1).Scan(&rows)
		if len(rows) > limit {
			(*meta)["next"] = encodeCursor(rows[limit].CreatedAt.UTC().UnixMicro())
		}

		return db.Limit(limit)
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
