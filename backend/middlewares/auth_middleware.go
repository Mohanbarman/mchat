package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

var unauthorizedErr = gin.H{
	"success": false,
	"message": "Unauthorized",
	"code":    401,
}

type AuthMiddleware struct {
	Jwt *lib.Jwt
	DB  *gorm.DB
}

func (a *AuthMiddleware) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedErr)
			return
		}

		_token := strings.Split(authHeader, " ")
		token := ""

		if len(_token) == 2 {
			token = _token[1]
		} else {
			lib.HttpResponse(401).Send(c)
			return
		}

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		sub, err := a.Jwt.ParseToken(token, lib.AccessToken)

		if err != nil {
			c.JSON(401, gin.H{
				"message": "Token is expired",
				"code":    401,
			})
			c.Abort()
			return
		}

		u := models.UserModel{}
		result := a.DB.Find(&u, &models.UserModel{UUID: sub})

		if result.RowsAffected <= 0 {
			c.JSON(401, gin.H{
				"message": "User doesn't exists please signup again",
				"code":    401,
			})
			c.Abort()
			return
		}

		c.Set("user", &u)
		c.Next()
	}
}
