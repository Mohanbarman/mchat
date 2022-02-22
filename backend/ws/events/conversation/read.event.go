package conversation

import (
	"gorm.io/gorm/clause"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

// Mark conversation as read
func (c *Controller) Read(payload interface{}, ctx *lib.WsContext) {
	dto := ReadConversationDTO{}
	if err := lib.ValidatePayload(&payload, &dto); err != nil {
		return
	}

	conversation := &models.ConversationModel{}
	records := ctx.DB.Preload(clause.Associations).Find(conversation, &models.ConversationModel{UUID: dto.ConversationID})

	if records.RowsAffected <= 0 {
		return
	}

	otherUserID := conversation.FromUserID
	otherUserColumn := "from_user_id"

	if ctx.User.ID == conversation.FromUserID {
		otherUserID = conversation.ToUserID
		otherUserColumn = "to_user_id"
		conversation.FromUserUnreadCount = 0
	} else {
		conversation.ToUserUnreadCount = 0
	}

	otherUser := &models.UserModel{}
	ctx.DB.Find(otherUser, otherUserID)

	ctx.DB.Scopes(models.SeenConversationMessageStatus(conversation.ID, otherUserColumn, otherUserID)).Model(&models.MessageModel{})
	ctx.DB.Save(&conversation)

	if con, err := c.Store.Get(otherUser.ID); err == nil {
		con.Send("conversation/seen", conversation.TransformForUser(otherUserID))
	}
	ctx.Send("conversation/update", conversation.TransformForUser(ctx.User.ID))
}
