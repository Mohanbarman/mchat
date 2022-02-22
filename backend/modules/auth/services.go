package auth

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mchat.com/api/config"
	"mchat.com/api/lib"
	"mchat.com/api/models"
)

type Service struct {
	Config *config.Config
	Db     *gorm.DB
	Jwt    *lib.Jwt
	Redis  *lib.RedisClient
	Mail   *lib.MailClient
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

func (service Service) Login(loginDto *LoginDto) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{}

	records := service.Db.Find(&user, &models.UserModel{Email: loginDto.Email})

	if records.RowsAffected <= 0 {
		e = lib.Error(WrongPasswordErr)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		e = lib.Error(WrongPasswordErr)
		return
	}

	accessToken, aerr := service.Jwt.SignToken(user.UUID, lib.AccessToken)
	refreshToken, rerr := service.Jwt.SignToken(user.UUID, lib.RefreshToken)

	if aerr != nil || rerr != nil {
		e = lib.Error(TokenGenerateErr)
	}

	result = user.Transform()
	result["access_token"] = accessToken
	result["refresh_token"] = refreshToken

	return
}

func (s *Service) SendResetPasswordMail(dto *ResetPasswordDTO) {
	user := models.UserModel{}
	userRecord := s.Db.First(&user, &models.UserModel{Email: dto.Email})

	if userRecord.RowsAffected > 0 {
		hash := lib.GenRandomString([]byte(user.UUID))
		s.Redis.Set(hash, user.UUID, time.Minute*15)
		s.Mail.SendMail([]string{dto.Email}, "Test mail", fmt.Sprintf("localhost:8000/reset-password?id=%s", hash))
	}
}

func (s *Service) ResetPassword(dto *ResetPasswordChangeDTO) (result lib.H, e *lib.ServiceError) {
	uuid, err := s.Redis.Get(dto.Secret)

	if err != nil {
		e = lib.Error(ResetPasswordLinkExpErr)
		return
	}

	s.Redis.Remove(dto.Secret)

	user := models.UserModel{}

	records := s.Db.First(&user, &models.UserModel{UUID: uuid.(string)})

	if records.RowsAffected <= 0 {
		e = lib.Error(UserNotFoundErr)
		return
	}

	err = user.SetPassword(dto.Password)
	s.Db.Save(&user)

	if err != nil {
		e = lib.Error(HashingPassErr)
	}

	return
}

func (s *Service) RefreshToken(dto *RefreshTokenDTO) (result lib.H, e *lib.ServiceError) {
	token := dto.Token

	sub, err := s.Jwt.ParseToken(token, lib.RefreshToken)

	if err != nil {
		e = &lib.ServiceError{Code: TokenExpireErr}
		return
	}

	user := &models.UserModel{}
	if records := s.Db.First(&user, "uuid=?", sub); records.RowsAffected < 1 {
		e = &lib.ServiceError{Code: TokenExpireErr}
		return
	}

	newAccessToken, err := s.Jwt.SignToken(user.UUID, lib.AccessToken)
	if err != nil {
		e = &lib.ServiceError{Code: TokenGenerateErr}
		return
	}

	result = lib.H{"token": newAccessToken}
	return
}
