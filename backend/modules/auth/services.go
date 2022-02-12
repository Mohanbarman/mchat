package auth

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/lib/jwt"
	"mchat.com/api/models"
)

type Service struct {
	Config *config.Config
	Db     *gorm.DB
}

func (service *Service) Register(registerDto *RegisterDto) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{
		Email:  registerDto.Email,
		Name:   registerDto.Name,
		Status: registerDto.Name,
	}

	err := user.SetPassword(registerDto.Password)

	if err != nil {
		e = lib.Error(HashingPassErr)
		return
	}

	if created := service.Db.Create(&user); created.Error != nil {
		e = lib.Error(EmailExistsErr)
		return
	}

	result = user.Transform()
	return
}

func (service Service) Login(loginDto *LoginDto, jwtService *jwt.JwtService) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{}

	records := service.Db.Find(&user, &models.UserModel{Email: loginDto.Email})

	if records.RowsAffected <= 0 {
		e = lib.Error(UserNotFoundErr)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		e = lib.Error(WrongPasswordErr)
		return
	}

	accessToken, aerr := jwtService.SignToken(user.UUID, jwt.AccessToken)
	refreshToken, rerr := jwtService.SignToken(user.UUID, jwt.RefreshToken)

	if aerr != nil || rerr != nil {
		e = lib.Error(TokenGenerateErr)
	}

	result = user.Transform()
	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return
}

func (service *Service) SendResetPasswordMail(dto *ResetPasswordDTO, s *lib.MailClient, rdb *lib.RedisClient) {
	user := models.UserModel{}
	userRecord := service.Db.First(&user, &models.UserModel{Email: dto.Email})

	if userRecord.RowsAffected > 0 {
		hash := lib.GenRandomString([]byte(user.UUID))
		rdb.Set(hash, user.UUID, time.Minute*15)
		s.SendMail([]string{dto.Email}, "Test mail", fmt.Sprintf("localhost:8000/reset-password?id=%s", hash))
	}
}

func (service *Service) ResetPassword(dto *ResetPasswordChangeDTO, rdb *lib.RedisClient) (result lib.H, e *lib.ServiceError) {
	uuid, err := rdb.Get(dto.Secret)

	if err != nil {
		e = lib.Error(ResetPasswordLinkExpErr)
		return
	}

	rdb.Remove(dto.Secret)

	user := models.UserModel{}

	records := service.Db.First(&user, &models.UserModel{UUID: uuid.(string)})

	if records.RowsAffected <= 0 {
		e = lib.Error(UserNotFoundErr)
		return
	}

	err = user.SetPassword(dto.Password)
	service.Db.Save(&user)

	if err != nil {
		e = lib.Error(HashingPassErr)
	}

	return
}
