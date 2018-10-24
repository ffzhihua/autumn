package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/common/response"
	"autumn/tools/exception"
	"autumn/models"
	"autumn/tools/crypt"
	"autumn/tools/validate"
)

type RegisterController struct {
	BaseController
}

func (i *RegisterController) Mobile(c *gin.Context) {
	defer exception.Catch(c)

	mobile          := c.PostForm("mobile")
	country         := c.PostForm("country")
	verify_code     := c.PostForm("verify_code")
	password        := c.PostForm("password")
	password_cfm    := c.PostForm("password_cfm")

	if len(mobile) != 11 {
		response.Fail(c, 12000)
		return
	}

	if country == "" {
		country = "86"
	}

	if password != password_cfm {
		response.Fail(c, 12002)
		return
	}

	v := i.service.sms.Verify(mobile, verify_code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	user := models.User{}

	//用户已存在
	if user.CheckByMobile(mobile) {
		response.Fail(c, 12001)
		return
	}

	user.Country    = country
	user.Mobile     = mobile
	user.Password   = crypt.GeneratorPassword(password)
	user.RegType    = 0
	user.RegIP      = c.ClientIP()

	if user.Create() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}

func (i *RegisterController) Email(c *gin.Context) {
	defer exception.Catch(c)

	email           := c.PostForm("email")
	verify_code     := c.PostForm("verify_code")
	password        := c.PostForm("password")
	password_cfm    := c.PostForm("password_cfm")


	if validate.IsEmail(email) == false {
		response.Fail(c, 13000)
		return
	}

	if password != password_cfm {
		response.Fail(c, 12002)
		return
	}

	v := i.service.email.Verify(email, verify_code, true)
	if v > 0 {
		response.Fail(c, v)
		return
	}

	user := models.User{}

	//用户已存在
	if user.CheckByEmail(email) {
		response.Fail(c, 12001)
		return
	}

	user.Email      = email
	user.Password   = crypt.GeneratorPassword(password)
	user.RegType    = 1
	user.RegIP      = c.ClientIP()

	if user.Create() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)

}