package service

import (
	"pishikot/pkg/models"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUser(c *gin.Context) {
	var user *models.User
	id, _ := strconv.Atoi(c.Query("id"))
	if id != 0 {
		user = &models.User{
			ID: uint(id),
		}
	} else {
		claims := jwt.ExtractClaims(c)
		user = &models.User{
			ID: uint(claims["id"].(float64)),
		}
	}
	exists := s.User.GetUser(user)
	s.User.Clean(user)
	c.JSON(200, gin.H{
		"exists": exists,
		"user":   user,
	})
}
