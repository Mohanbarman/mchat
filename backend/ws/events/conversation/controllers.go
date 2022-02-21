package conversation

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm/clause"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Controller struct {
	Store *lib.WsStore
}

// Send a new message to a user
// payload : data send from the user
// ctx : websocket context
func (c *Controller) Send(payload interface{}, ctx *lib.WsContext) {
	dto := SendMessageDTO{}
	mapstructure.Decode(payload, &dto)

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
	).Find(&conversation)

	// creating the conversation if doesn't exists
	if records.RowsAffected <= 0 {
		conversation = models.ConversationModel{
			FromUserID: ctx.User.ID,
			ToUserID:   toUser.ID,
		}
		ctx.DB.Create(&conversation)
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

	otherUserID := conversation.FromUserID

	if conversation.FromUserID == ctx.User.ID {
		otherUserID = conversation.ToUserID
		conversation.ToUserUnreadCount += 1
	} else {
		conversation.FromUserUnreadCount += 1
	}

	conversation.LastMessageText = dto.Text
	conversation.LastMessageTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	ctx.DB.Save(&conversation)

	if con, err := c.Store.Get(uint(otherUserID)); err == nil {
		con.Send("conversation/new_message", message.Transform())
		message.Status = models.MessageStatusDelivered
		ctx.DB.Save(&message)
		ctx.Send("conversation/delivered", message.Transform())
	}
}

func (c *Controller) Read(payload interface{}, ctx *lib.WsContext) {
	dto := ReadConversationDTO{}
	mapstructure.Decode(payload, &dto)

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
		con.Send("conversation/seen", conversation.Transform())
	}
}
