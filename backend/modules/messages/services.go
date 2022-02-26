package messages

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Service struct {
	DB *gorm.DB
	WS *lib.WsStore
}

func (s *Service) GetAll(conversationId string, dto *GetAllDTO, user *models.UserModel) (page lib.H, data []map[string]interface{}, err int) {
	conv := models.ConversationModel{}
	if r := s.DB.First(&conv, "uuid = ?", conversationId); r.RowsAffected < 1 {
		return
	}

	records := []models.MessageModel{}
	page = map[string]interface{}{}

	var pageErr int
	scope := lib.CursorPaginate("messages", &pageErr, &dto.CursorPaginationDTO, &page, true)
	s.DB.Scopes(scope).Preload(clause.Associations).Where("conversation_id = ?", conv.ID).Order("created_at desc").Find(&records)

	if pageErr != 0 {
		err = lib.InvalidCursorErr
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

	return
}
