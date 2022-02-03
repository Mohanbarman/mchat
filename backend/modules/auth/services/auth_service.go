package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/models"
	dto "mchat.com/api/modules/auth/dto"
	"mchat.com/api/utils"
)

type AuthService struct {
	Config *config.Config
	Db     *gorm.DB
}

func (service *AuthService) Register(c *gin.Context) {
	registerDto := c.MustGet("data").(*dto.RegisterDto)

	user := models.UserModel{
		Email:  registerDto.Email,
		Name:   registerDto.Name,
		Status: registerDto.Name,
	}

	err := user.SetPassword(registerDto.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": "Interal server error",
		})
		return
	}

	if created := service.Db.Create(&user); created.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"email": []string{"Email already exists"}})
		return
	}

	c.JSON(http.StatusOK, user.Transform())
}

func (service AuthService) Login(c *gin.Context, jwtService *JwtService) {
	loginDto := c.MustGet("data").(*dto.LoginDto)

	user := models.UserModel{}

	fmt.Println(loginDto, user)

	service.Db.Find(&user, &models.UserModel{Email: loginDto.Email})

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"email": []string{"email not found"}})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	}

	accessToken, aerr := jwtService.SignToken(user.UUID, AccessToken)
	refreshToken, rerr := jwtService.SignToken(user.UUID, RefreshToken)

	if aerr != nil || rerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	result := user.Transform()
	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    200,
		"data":    result,
	})
}

func (service *AuthService) GetMe(c *gin.Context) {
	u := c.MustGet("user").(*models.UserModel)
	c.JSON(200, u.Transform())
}

func (service *AuthService) SendResetPasswordMail(c *gin.Context, s *utils.MailClient, rdb *utils.RedisClient) {
	data := c.MustGet("data").(*dto.ResetPasswordDTO)

	user := models.UserModel{}
	userRecord := service.Db.First(&user, &models.UserModel{Email: data.Email})

	if userRecord.RowsAffected > 0 {
		hash := utils.GenRandomString([]byte(user.UUID))
		rdb.Set(hash, user.UUID, time.Minute*15)
		fmt.Println(hash)
		go s.SendMail([]string{data.Email}, "Test mail", fmt.Sprintf("localhost:8000/reset-password?id=%s", hash))
	}

	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"message": "You will receive an email with reset password link if your email exists",
	})
}

func (service *AuthService) ResetPassword(c *gin.Context, rdb *utils.RedisClient) {
	data := c.MustGet("data").(*dto.ResetPasswordChangeDTO)

	uuid, err := rdb.Get(data.Secret)

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"message": "Reset password link expired please try again",
		})
		return
	}

	rdb.Remove(data.Secret)

	user := models.UserModel{}

	result := service.Db.First(&user, &models.UserModel{UUID: uuid.(string)})

	if result.RowsAffected <= 0 {
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"message": "Reset password link expired please try again",
		})
	}

	err = user.SetPassword(data.Password)
	service.Db.Save(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"message": "Failed to set password",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    http.StatusOK,
		"message": "Password changed successfully",
	})
}
