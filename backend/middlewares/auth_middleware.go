package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mchat.com/api/models"
	auth "mchat.com/api/modules/auth/services"
)

var unauthorizedErr = gin.H{
	"success": false,
	"message": "Unauthorized",
	"code":    401,
}

type AuthMiddleware struct {
	Jwt *auth.JwtService
	DB  *gorm.DB
}

func (a *AuthMiddleware) Validate(tokenType auth.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		token := strings.Split(authHeader, " ")[1]

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		sub, err := a.Jwt.ParseToken(token, tokenType)

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
