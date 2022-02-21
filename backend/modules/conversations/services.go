package conversations

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetAll(dto *GetAllDTO, user *models.UserModel) (page lib.H, data []map[string]interface{}, err int) {
	records := []models.ConversationModel{}
	page = lib.H{}
	data = []map[string]interface{}{}

	pageErr := 0
	scope := lib.CursorPaginate((&models.ConversationModel{}).TableName(), &pageErr, &dto.CursorPaginationDTO, &page, true)
	s.DB.Scopes(models.FindAllUserConversations(user.ID)).Scopes(scope).Find(&records)

	if pageErr == lib.InvalidCursorErr {
		err = pageErr
		return
	}

	for i := range records {
		r := records[i].Transform()
		if records[i].FromUserID == user.ID {
			r["user"] = records[i].ToUser.Transform()
		} else {
			r["user"] = records[i].FromUser.Transform()
		}
		data = append(data, r)
	}

	return
}

func (s *Service) GetOne(conversationID string, user *models.UserModel) (data lib.H, err *lib.ServiceError) {
	conversation := models.ConversationModel{}

	if r := s.DB.Preload(clause.Associations).First(&conversation, "uuid = ?", conversationID); r.RowsAffected < 1 {
		err = &lib.ServiceError{Code: NotFoundErr}
		return
	}

	data = conversation.Transform()
	if conversation.FromUserID == user.ID {
		data["user"] = conversation.ToUser.Transform()
	} else {
		data["user"] = conversation.FromUser.Transform()
	}

	return
}
