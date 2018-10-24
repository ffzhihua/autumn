package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/tools/crypt"
	"autumn/tools/i18n"
	"autumn/common/response"
	"log"
	"autumn/models"
	"time"
	"autumn/tools/cfg"
	"autumn/tools/validate"
	"autumn/tools/exception"
)

type SendController struct {
	BaseController
}

func (i *SendController) MobileCode(c *gin.Context) {
	defer exception.Catch(c)

	mobile  := c.PostForm("mobile")
	country := c.PostForm("country")

	if country == "" {
		country = "86"
	}

	uid := i.GetUid(c)
	if mobile == "" && uid > 0 {
		mobile = i.UserInfo(c).Mobile
		if mobile == "" {
			response.Fail(c, 12003)
			return
		}
	}

	if len(mobile) != 11 {
		response.Fail(c, 12000)
		return
	}

	code := crypt.RandCode()

	sms := models.Sms{}

	sms.Code = code
	sms.Expire = time.Now().Unix() + cfg.Get("sms", "expire").Int()
	sms.Mobile = mobile
	sms.Country = country
	sms.Status = 0

	if !sms.Create() {
		response.Fail(c ,11004)
		return
	}

	err := i.service.sms.Send(mobile, code, country, i18n.Lang(c))
	if err != nil {
		response.Fail(c, 11004)
		log.Println("Send.MobileCode:", err)
		return
	}

	response.Success(c, nil)
}

func (i *SendController) EmailCode(c *gin.Context) {
	defer exception.Catch(c)

	email  := c.PostForm("email")

	if validate.IsEmail(email) == false {
		response.Fail(c, 13000)
		return
	}

	code := crypt.RandCode()

	eml := models.Email{}

	eml.Code = code
	eml.Expire = time.Now().Unix() + cfg.Get("sms", "expire").Int()
	eml.Email = email
	eml.Status = 0

	if !eml.Create() {
		response.Fail(c ,11004)
		return
	}

	send := i.service.email.SendCode(email, code, i18n.Lang(c))
	if send == false {
		response.Fail(c, 11004)
		return
	}

	response.Success(c, nil)
}