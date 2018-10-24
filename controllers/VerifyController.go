package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/models"
	"autumn/common/response"
	"autumn/tools/ga"
	"autumn/result"
	"autumn/tools/exception"
	"fmt"
)

type VerifyController struct {
	BaseController
}

//获取GA验证秘钥
func (i *VerifyController) GetGASecret(c *gin.Context) {
	defer exception.Catch(c)

	uid := i.GetUid(c)

	uga := &models.UserGA{}
	secret := ga.Secret()

	info := uga.Info(uid)
	user := (&models.User{}).Info(uid)

	if info.Status == 1 {
		response.Fail(c, 14000)
		return
	}

	usename := i.service.user.GetFullUsername(user)

	if info.Uid == 0 {
		uga.Uid = uid
		uga.Secret = secret
		uga.Status = 0

		if uga.Create() {
			url := "otpauth://totp/rrbank-"+usename+"?secret=" + secret
			response.Success(c, result.GaSecret{url, secret})
			return
		}

		response.Fail(c, 1)
		return
	}else {

		info.Secret = secret
		if info.Update() {
			url := "otpauth://totp/rrbank-"+usename+"?secret=" + secret
			response.Success(c, result.GaSecret{url, secret})
			return
		}
	}

	response.Fail(c, 1)
	return
}

//确认GA验证
func (i *VerifyController) GaConfirm(c *gin.Context) {
	defer exception.Catch(c)

	verify_code := c.PostForm("verify_code")
	ga_code := c.PostForm("ga_code")
	uid := i.GetUid(c)

	uga := (&models.UserGA{}).Info(uid)
	info := i.UserInfo(c)

	if info.GoogleAuth == 1 {
		response.Fail(c, 14000)
		return
	}

	if uga.Status == 1 {
		response.Fail(c, 14000)
		return
	}

	if v := i.service.sms.Verify(i.UserInfo(c).Mobile, verify_code, true); v > 0 {
		response.Fail(c, v)
		return
	}

	if ga.Code(uga.Secret) == ga_code{
		uga.Status = 1

		fmt.Println("uga:", uga)
		if uga.Update() {

			info.GoogleAuth = 1
			info.Update()

			response.Success(c, nil)
			return
		}
	}else{
		response.Fail(c, 14004)
		return
	}

	response.Fail(c, 1)
	return
}

//验证邮箱
func (i *VerifyController) Email(c *gin.Context) {
	defer exception.Catch(c)

	email   := c.PostForm("email")
	code    := c.PostForm("verify_code")

	info := i.UserInfo(c)
	if info.Email != "" {
		response.Fail(c, 13001)
		return
	}

	v := i.service.email.Verify(email, code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	if info.CheckVerifyEmail(email) {
		response.Fail(c, 13002)
		return
	}

	info.Email = email
	if info.Update() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}

//验证手机
func (i *VerifyController) Mobile(c *gin.Context) {
	defer exception.Catch(c)

	mobile  := c.PostForm("mobile")
	country := c.PostForm("country")
	code    := c.PostForm("verify_code")

	info := i.UserInfo(c)
	if info.Mobile != "" {
		response.Fail(c, 12004)
		return
	}

	v := i.service.sms.Verify(mobile, code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	if info.CheckVerifyMobile(mobile) {
		response.Fail(c, 12005)
		return
	}

	info.Country = country
	info.Mobile = mobile
	if info.Update() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}
