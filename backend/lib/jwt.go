package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"mchat.com/api/config"
)

type Jwt struct {
	Config *config.JwtConfig
}

type TokenType string

const (
	RefreshToken TokenType = "refresh_token"
	AccessToken  TokenType = "access_token"
)

func (s *Jwt) SignToken(sub string, t TokenType) (signedToken string, err error) {
	var exp int64
	var scope string

	if t == RefreshToken {
		exp = time.Now().UTC().Add(time.Hour * 24 * time.Duration(s.Config.RefreshTokenExpDays)).Unix()
		scope = "refresh_token"
	} else {
		exp = time.Now().UTC().Add(time.Hour * 2).Unix()
		scope = "access_token"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":   sub,
		"exp":   exp,
		"iat":   time.Now().UTC().Unix(),
		"scope": scope,
	})

	signedToken, err = token.SignedString([]byte(s.Config.Secret))
	return
}

func (s *Jwt) ParseToken(tokenString string, t TokenType) (sub string, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(s.Config.Secret), nil
	})

	if token == nil {
		err = errors.New("invalid token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["scope"].(string) != string(t) {
			err = errors.New("token is invalid")
			return
		}

		sub = claims["sub"].(string)
		return
	}

	err = errors.New("token is invalid")
	return
}
