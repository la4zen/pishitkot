package service

import (
	"fmt"
	"pishikot/pkg/models"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func (s *Service) CheckTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user := &models.User{
		ID:       uint(claims["id"].(float64)),
		Login:    claims["login"].(string),
		UserType: claims["user_type"].(string),
	}
	fmt.Println(user) // тут логика отправки тасков на тест, допилю как нибудь когда нибудь.
}

func (s *Service) GetTask(c *gin.Context) {
	_id, exists := c.Params.Get("task_id")
	if !exists {
		c.JSON(404, gin.H{
			"error": "Такого задания не существует!",
		})
		return
	}
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.JSON(500, gin.H{
			"errpr": err.Error(),
		})
		return
	}
	s.Task.GetTask(uint(id))
}
