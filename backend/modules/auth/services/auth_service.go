package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mchat.com/api/config"
	dto "mchat.com/api/modules/auth/dto"
	users "mchat.com/api/modules/users/models"
)

type AuthService struct {
	Config *config.Config
	Db     *gorm.DB
}

func (service *AuthService) Register(c *gin.Context) {
	registerDto := c.MustGet("data").(*dto.RegisterDto)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	user := users.UserModel{
		Email:    registerDto.Email,
		Password: string(hashedPasswordBytes),
		Name:     registerDto.Name,
		Status:   registerDto.Name,
	}

	if created := service.Db.Create(&user); created.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"email": []string{"Email already exists"}})
		return
	}

	c.JSON(http.StatusOK, user.Transform())
}

func (service AuthService) Login(c *gin.Context, jwtService *JwtService) {
	loginDto := c.MustGet("data").(*dto.LoginDto)

	user := users.UserModel{}

	fmt.Println(loginDto, user)

	service.Db.Find(&user, &users.UserModel{Email: loginDto.Email})

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

	c.JSON(http.StatusOK, result)
}

func (service *AuthService) GetMe(c *gin.Context) {
	u := c.MustGet("user").(*users.UserModel)
	c.JSON(200, u.Transform())
}
