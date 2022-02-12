package messages

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/lib/pagination"
	"mchat.com/api/models"
	"mchat.com/api/validation"
)

type Controller struct {
	DB     *gorm.DB
	Config *config.Config
}

func (ctrl *Controller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &pagination.CursorPaginationDTO{}
		conv_id := c.Param("conversation_id")
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}

		conv := models.ConversationModel{}
		if r := ctrl.DB.First(&conv, "uuid = ?", conv_id); r.RowsAffected < 1 {
			lib.HttpResponse(404).Message("Conversation not found").Send(c)
			return
		}

		data := []models.MessageModel{}
		// dto := c.MustGet("data").(*pagination.CursorPaginationDTO)
		pageData := map[string]interface{}{}
		var e int
		ctrl.DB.Scopes(pagination.CursorPaginate("messages", &e, dto, &pageData)).Preload("Conversation").Where("conversation_id = ?", conv.ID).Find(&data)
		mapData := []map[string]interface{}{}
		for d := range data {
			mapData = append(mapData, data[d].Transform())
		}

		lib.HttpResponse(200).Data(mapData).Page(pageData).Send(c)
	}
}
