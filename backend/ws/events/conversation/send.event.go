package conversation

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

// Send a new message to a user
// payload : data send from the user
// ctx : websocket context
func (c *Controller) Send(payload map[string]interface{}, ctx *lib.WsContext) {
	dto := SendMessageDTO{}
	if err := lib.ValidatePayload(&payload, &dto); err != nil {
		return
	}

	// finding the user by to send id
	toUser := models.UserModel{}
	records := ctx.DB.First(&toUser, &models.UserModel{UUID: dto.UserID})
	if records.RowsAffected <= 0 {
		ctx.SendJSON(gin.H{
			"event": "conversation/error",
			"code":  "USER_NOT_FOUND",
		})
		return
	}

	// finding conversation of the user
	conversation := models.ConversationModel{}
	records = ctx.DB.Scopes(
		models.FindUserConversation(toUser.ID, ctx.User.ID),
	).Preload(clause.Associations).Find(&conversation)

	// creating the conversation if doesn't exists
	if records.RowsAffected <= 0 {
		conversation = models.ConversationModel{
			FromUserID: ctx.User.ID,
			ToUserID:   toUser.ID,
		}
		ctx.DB.Create(&conversation)

		otherUserID := conversation.GetOtherUser(ctx.User.ID)
		if con, err := c.Store.Get(uint(otherUserID)); err == nil {
			con.Send("conversation/new", conversation.TransformForUser(ctx.User.ID))
		}
	}

	// creating message
	message := models.MessageModel{
		Text:           dto.Text,
		ConversationID: conversation.ID,
		Status:         models.MessageStatusSent,
		FromUserID:     int(ctx.User.ID),
		ToUserID:       int(toUser.ID),
	}
	ctx.DB.Create(&message)
	ctx.DB.Preload("Conversation").Find(&message, message.ID)

	otherUserID := conversation.GetOtherUser(ctx.User.ID)

	// incrementing unread count of other user
	if conversation.FromUserID == ctx.User.ID {
		conversation.ToUserUnreadCount += 1
	} else {
		conversation.FromUserUnreadCount += 1
	}

	// setting current message as last sent message
	conversation.LastMessageText = dto.Text
	conversation.LastMessageTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	// updating the row in table
	ctx.DB.Save(&conversation)

	// sending events to other user if user is connected
	if con, err := c.Store.Get(uint(otherUserID)); err == nil {
		con.Send("conversation/new_message", message.Transform())

		message.Status = models.MessageStatusDelivered
		ctx.DB.Save(&message)

		ctx.Send("conversation/delivered", lib.H{
			"conversation": message.Conversation.UUID,
			"message":      message.UUID,
		})

		con.Send("conversation/update", conversation.TransformForUser(otherUserID))
	}
	ctx.Send("conversation/update", conversation.TransformForUser(ctx.User.ID))
}
