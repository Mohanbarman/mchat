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
		r := records[i].TransformForUser(user.ID)
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

	data = conversation.TransformForUser(user.ID)

	return
}

func (s *Service) Create(dto *CreateDTO, user *models.UserModel) (data lib.H, err *lib.ServiceError) {
	otherUser := models.UserModel{}

	if records := s.DB.Find(&otherUser, &models.UserModel{Email: dto.Email}); records.RowsAffected < 1 {
		err = &lib.ServiceError{Code: UserNotFoundErr}
		return
	}

	if record := s.DB.Scopes(models.FindUserConversation(user.ID, otherUser.ID)).Find(&models.ConversationModel{}); record.RowsAffected > 0 {
		err = &lib.ServiceError{Code: AlreadyExistsErr}
		return
	}

	conversation := models.ConversationModel{
		FromUserID: user.ID,
		ToUserID:   otherUser.ID,
	}

	s.DB.Save(&conversation)

	newConversation := &models.ConversationModel{}

	s.DB.Preload(clause.Associations).Find(&newConversation, conversation.ID)

	data = newConversation.TransformForUser(user.ID)
	return
}
