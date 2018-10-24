package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/common/response"
	"autumn/models"
	"time"
	"autumn/tools/crypt"
	"autumn/result"
	"autumn/tools/validate"
)

type ForgotController struct {
	BaseController
}

func (i *ForgotController) ResetPassword(c *gin.Context)  {
	token           := c.PostForm("token")
	password        := c.PostForm("password")
	password_cfm    := c.PostForm("password_cfm")

	if len(token) != 40 {
		response.Fail(c, 10007)
		return
	}

	if len(password) < 6 {
		response.Fail(c, 12006)
		return
	}

	if password != password_cfm {
		response.Fail(c, 12002)
		return
	}

	tInfo := (&models.UserResetPwd{}).Info(token)

	if tInfo.Id == 0 || tInfo.Status == 1 {
		response.Fail(c, 10007)
		return
	}

	if tInfo.Expire < time.Now().Unix() {
		response.Fail(c, 10008)
		return
	}

	uInfo := (&models.User{}).Info(tInfo.Uid)

	uInfo.Password = crypt.GeneratorPassword(password)
	if uInfo.Update() {

		tInfo.Status = 1
		tInfo.Update()

		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}

func (i *ForgotController) VerifyCodeOfMobile(c *gin.Context){
	mobile      := c.PostForm("mobile")
	verify_code := c.PostForm("verify_code")

	if len(mobile) != 11 {
		response.Fail(c, 12000)
		return
	}

	v := i.service.sms.Verify(mobile, verify_code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	info := (&models.User{}).InfoByUser(models.User{Mobile: mobile})
	if info.Id == 0 {
		response.Fail(c, 10003)
		return
	}

	rToken := i.service.user.GetForgotResetToken(info.Id, 0)

	if rToken != "" {
		response.Success(c, result.ForgotVerify{rToken})
		return
	}

	response.Fail(c, 1)
}

func (i *ForgotController) VerifyCodeOfEmail(c *gin.Context) {
	email      := c.PostForm("email")
	verify_code := c.PostForm("verify_code")

	if !validate.IsEmail(email){
		response.Fail(c, 13000)
		return
	}

	v := i.service.email.Verify(email, verify_code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	info := (&models.User{}).InfoByUser(models.User{Email: email})
	if info.Id == 0 {
		response.Fail(c, 10003)
		return
	}

	rToken := i.service.user.GetForgotResetToken(info.Id, 1)

	if rToken != "" {
		response.Success(c, result.ForgotVerify{rToken})
		return
	}

	response.Fail(c, 1)
}