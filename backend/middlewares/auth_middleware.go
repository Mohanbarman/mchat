package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	auth "mchat.com/api/modules/auth/services"
	user "mchat.com/api/modules/users/models"
)

var unauthorizedErr = gin.H{
	"success": false,
	"message": "Unauthorized",
	"code":    401,
}

func AuthMiddleware(jwtService *auth.JwtService, tokenType auth.TokenType, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		token := strings.Split(authHeader, " ")[1]

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		sub, err := jwtService.ParseToken(token, tokenType)

		if err != nil {
			c.JSON(401, gin.H{
				"message": "Token is expired",
				"code":    401,
			})
			c.Abort()
			return
		}

		u := user.UserModel{}
		result := db.Find(&u, &user.UserModel{UUID: sub})

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
