package messages

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/models"
	"mchat.com/api/validation"
)

type Controller struct {
	DB      *gorm.DB
	Config  *config.Config
	Service *Service
}

func (ctrl *Controller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := &GetAllDTO{}
		conv_id := c.Param("conversation_id")
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}

		user := c.MustGet("user").(*models.UserModel)

		page, data, err := ctrl.Service.GetAll(conv_id, dto, user)

		if err == lib.InvalidCursorErr {
			lib.HttpResponse(400).Message("Invalid cursor").Send(c)
			return
		}

		lib.HttpResponse(200).Data(data).Page(page).Send(c)
	}
}
