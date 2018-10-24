package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/tools/crypt"
	"autumn/common/response"
	"strings"
	"autumn/models"
	"autumn/tools/token"
	"autumn/result"
	"autumn/tools/exception"
	"time"
)

type UserController struct {
	BaseController
}

func (i *UserController) UpdatePassword(c *gin.Context) {
	defer exception.Catch(c)

	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")
	new_password_cfm := c.PostForm("new_password_cfm")

	info := i.UserInfo(c)
	if ! crypt.VerifyPassword(info.Password, old_password) {
		response.Fail(c, 10002)
		return
	}

	if old_password == new_password {
		response.Fail(c, 12007)
		return
	}

	if len(new_password) < 6 {
		response.Fail(c, 12006)
		return
	}

	if new_password != new_password_cfm {
		response.Fail(c, 12002)
		return
	}

	info.Password = crypt.GeneratorPassword(new_password)
	if info.Update() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}

func (i *UserController) Login(c *gin.Context) {
	defer exception.Catch(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		response.Fail(c, 10004)
		return
	}

	if password == "" {
		response.Fail(c, 10006)
		return
	}

	var info models.User

	if strings.Index(username, "@") > 0 {
		info = (&models.User{}).InfoByUser(models.User{Email: username})
	}else{
		info = (&models.User{}).InfoByUser(models.User{Mobile: username})
	}

	if info.Id == 0 {
		response.Fail(c, 10003)
		return
	}

	if !crypt.VerifyPassword(info.Password, password) {
		response.Fail(c, 10005)
		return
	}

	//记录登录日志
	login := &models.UserLogin{}
	login.Uid = info.Id
	login.LoginIp = c.ClientIP()

	login.Create()

	//更新最后登录时间
	info.LastLogin = time.Now().Unix()
	info.Update()


	response.Success(c , result.Login{token.Generator(info.Id)})
}

func (i *UserController) Info(c *gin.Context) {
	defer exception.Catch(c)

	info := (&models.User{}).Info(i.GetUid(c))

	var (
		email_valid int
		mobile_valid int
	)

	username := i.service.user.GetSmileUsername(info)

	if info.Mobile != "" {
		mobile_valid = 1
	}

	if info.Email != "" {
		mobile_valid = 1
	}

	ret := result.UserInfo{Uid:info.Id,
			Username:username,
			LastLogin:info.LastLogin,
		MobileValid:mobile_valid,
		EmailValid:email_valid,
		NameValid:info.RealName,
		GaValid:info.GoogleAuth,
		RegTime:info.CreatedAt,
	}

	response.Success(c, ret)
}
