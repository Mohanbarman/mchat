package conversation

import (
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

func (c *Controller) Typing(payload interface{}, ctx *lib.WsContext) {
	dto := TypingEventDTO{}
	if err := lib.ValidatePayload(&payload, &dto); err != nil {
		return
	}

	conversation := models.ConversationModel{}

	if records := ctx.DB.First(&conversation, "uuid = ?", dto.ConversationID); records.RowsAffected < 1 {
		ctx.SendErr("errors/conversation", "conversation not found")
		return
	}

	otherUserID := conversation.GetOtherUser(ctx.User.ID)

	if con, err := c.Store.Get(otherUserID); err != nil {
		con.Send("conversation/typing", conversation.UUID)
	}
}
