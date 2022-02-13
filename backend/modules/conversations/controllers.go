package conversations

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
	DB      *gorm.DB
	Config  *config.Config
	Service *Service
}

func (ctrl *Controller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := GetAllDTO{}
		if ok := validation.ValidateReq(&dto, c); !ok {
			return
		}
		user := c.MustGet("user").(*models.UserModel)

		page, data, err := ctrl.Service.GetAll(&dto, user)

		if err == pagination.InvalidCursorErr {
			lib.HttpResponse(400).Message("Invalid cursor").Send(c)
			return
		}

		lib.HttpResponse(200).Data(data).Page(page).Send(c)
	}
}

func (ctrl *Controller) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		conversation_id := c.Param("conversation_id")
		user := c.MustGet("user").(*models.UserModel)
		data, err := ctrl.Service.GetOne(conversation_id, user)

		if err != nil {
			HttpErrors[err.Code].Send(c)
			return
		}

		lib.HttpResponse(200).Data(data).Send(c)
	}
}
