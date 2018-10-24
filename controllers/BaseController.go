package controllers

import "github.com/gin-gonic/gin"
import (
	"strconv"
	"autumn/services"
	"autumn/tools/cfg"
	"autumn/models"
)
type Service struct {
	sms   services.SmsService
	email services.EmailService
	ga    services.GaService
	user  services.UserService
}
type BaseController struct {
	service Service
}

func (i *BaseController) GetUid(c *gin.Context) int {
	var uid string

	id,ok := c.Keys["uid"]
	if ok {
		uid = id.(string)
	}

	if uid == "" && cfg.Get("env", "dev").Bool() {
		uid = c.GetHeader("uid")
	}

	iid,_ := strconv.Atoi(uid)
	return iid
}

func (i *BaseController) UserInfo(c *gin.Context) models.User {
	return (&models.User{}).Info(i.GetUid(c))
}
