package conversation

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm/clause"
	"mchat.com/api/models"
	"mchat.com/api/modules/ws/connection"
)

type Controller struct {
	Manager *connection.ConnStore
}

// Send a new message to a user
// payload : data send from the user
// ctx : connection context
func (c *Controller) Send(payload interface{}, ctx *connection.Context) {
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
	}

	// finding conversation of the user
	conversation := models.ConversationModel{}
	records = ctx.DB.Find(
		&conversation,
		"(to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)",
		toUser.ID, ctx.User.ID, ctx.User.ID, toUser.ID,
	)

	// creating the conversation if doesn't exists
	if records.RowsAffected <= 0 {
		conversation := models.ConversationModel{
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
	}

	if con, err := c.Manager.Get(uint(otherUserID)); err == nil {
		con.Send("conversation/new_message", message.Transform())
		message.Status = models.MessageStatusDelivered
		ctx.DB.Save(&message)
		ctx.Send("conversation/delivered", message.Transform())
	}
}

func (c *Controller) Read(payload interface{}, ctx *connection.Context) {
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
	}

	otherUser := &models.UserModel{}
	ctx.DB.Find(otherUser, otherUserID)

	records = ctx.DB.Debug().Model(&models.MessageModel{}).Where("conversation_id = ? AND ?=?", conversation.ID, otherUserColumn, otherUserID).Update("status", models.MessageStatusSeen)
	fmt.Println(records)

	if con, err := c.Manager.Get(otherUser.ID); err == nil {
		con.Send("conversation/seen", conversation.Transform())
	}
}
