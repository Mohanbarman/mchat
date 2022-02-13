package messages

import (
	"gorm.io/gorm"
	"mchat.com/api/lib"
	"mchat.com/api/lib/pagination"
	"mchat.com/api/models"
	"mchat.com/api/modules/ws/connection"
)

type Service struct {
	DB *gorm.DB
	WS *connection.ConnStore
}

func (s *Service) GetAll(conversationId string, dto *GetAllDTO, user *models.UserModel) (page lib.H, data []map[string]interface{}, err int) {
	conv := models.ConversationModel{}
	if r := s.DB.First(&conv, "uuid = ?", conversationId); r.RowsAffected < 1 {
		return
	}

	records := []models.MessageModel{}
	page = map[string]interface{}{}

	var pageErr int
	scope := pagination.CursorPaginate("messages", &pageErr, &dto.CursorPaginationDTO, &page, true)
	s.DB.Scopes(scope).Preload("FromUser").Preload("ToUser").Where("conversation_id = ?", conv.ID).Order("created_at desc").Find(&records)

	if pageErr != 0 {
		err = pagination.InvalidCursorErr
		return
	}

	data = []map[string]interface{}{}

	for i := range records {
		d := records[i].Transform()
		d["is_me"] = false
		if records[i].FromUserID == int(user.ID) {
			d["is_me"] = true
		}
		data = append(data, d)
	}

	otherUserID := conv.FromUserID
	if conv.FromUserID == user.ID {
		otherUserID = conv.ToUserID
	}

	if conn, err := s.WS.Get(otherUserID); err == nil {
		conn.Send("conversation/seen", conv.Transform())
	}

	return
}
