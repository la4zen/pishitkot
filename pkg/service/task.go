package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"pishikot/pkg/models"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type RequestWithCode struct {
	Code string `json:"code,omitempty"`
}

func (s *Service) CheckTask(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user := &models.User{
		ID:       uint(claims["id"].(float64)),
		Login:    claims["login"].(string),
		UserType: claims["user_type"].(string),
	}
	req := &RequestWithCode{}
	c.Bind(&req)
	if req.Code == "" {
		c.JSON(400, "code is not defined")
		return
	}
	var commandpath string = "/tmp/" + user.Login + ".py"
	ioutil.WriteFile(commandpath, []byte(req.Code), 0777)
	cmd, _ := exec.Command("pylint", commandpath, "-E", "-f", "json").Output()
	if string(cmd) != "[]\n" {
		_json := &[]map[string]interface{}{}
		json.Unmarshal(cmd, &_json)
		c.JSON(200, gin.H{
			"error":  _json,
			"output": nil,
		})
		return
	}
	cmd, err := exec.Command("python", commandpath).Output()
	os.Remove(commandpath)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  nil,
			"output": string(cmd),
		})
		return
	}
	c.JSON(200, gin.H{
		"error":  nil,
		"output": string(cmd),
	})
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
			"error": err.Error(),
		})
		return
	}
	s.Task.GetTask(uint(id))
}
