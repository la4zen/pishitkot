package service

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (s *Service) UpgradeConnection(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	s.EventService.NewEventConnection(c.Writer, c.Request, uint(claims["id"].(float64)))
}
