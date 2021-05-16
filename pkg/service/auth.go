package service

import (
	"pishikot/pkg/models"
	"pishikot/pkg/util"

	"github.com/gin-gonic/gin"
)

func (s *Service) Login(c *gin.Context) {
	user := &models.User{}
	c.Bind(&user)
	if err := user.ValidateForLogin(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	util.MD5Password(&user.Password)
	if err := s.User.FindUser(user); err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := s.GenerateTokenPair(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	s.User.Clean(user)
	c.JSON(200, gin.H{
		"user" : user,
		"token":         token[0],
		"refresh_token": token[1],
	})
}

func (s *Service) Register(c *gin.Context) {
	user := &models.User{}
	c.Bind(&user)
	if err := user.ValidateForRegister(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	util.MD5Password(&user.Password)
	err := s.User.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := s.GenerateTokenPair(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	s.User.Clean(user)
	c.JSON(200, gin.H{
		"user" : user,
		"token":         token[0],
		"refresh_token": token[1],
	})
}
