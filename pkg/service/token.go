package service

import (
	"pishikot/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func (s *Service) GenerateTokenPair(user *models.User) (*[2]string, error) {
	_jwt := jwt.New(jwt.SigningMethodHS256)
	claims := _jwt.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["login"] = user.Login
	claims["user_type"] = user.UserType
	claims["exp"] = time.Now().Add(24 * 30 * time.Hour).Unix()
	token, err := _jwt.SignedString([]byte(s.Config.Jwt.Key))
	if err != nil {
		return nil, err
	}
	_jwt = jwt.New(jwt.SigningMethodHS256)
	claims = _jwt.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	claims["id"] = user.ID
	claims["login"] = user.Login
	claims["user_type"] = user.UserType
	claims["orig_iat"] = time.Now().Add(time.Minute).Unix()
	refreshToken, err := _jwt.SignedString([]byte(s.Config.Jwt.Key))
	if err != nil {
		return nil, err
	}
	return &[2]string{token, refreshToken}, nil
}

func (s *Service) Accessible(c *gin.Context) {
	c.Status(200)
}
