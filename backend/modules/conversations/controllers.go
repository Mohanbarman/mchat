package conversations

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Controller struct {
	DB     *gorm.DB
	Config *config.Config
}

func (ctrl *Controller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := []models.ConversationModel{}
		pageData := lib.H{}
		mapData := []map[string]interface{}{}
		for d := range data {
			mapData = append(mapData, data[d].Transform())
		}
		lib.HttpResponse(200).Data(mapData).Page(pageData).Send(c)
	}
}
