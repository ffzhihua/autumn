package services

import (
	"autumn/models"
	"strings"
	"autumn/tools/fn"
	"autumn/tools/crypt"
	"strconv"
	"time"
	"autumn/tools/cfg"
)

type UserService struct {

}

func (i *UserService) GetSmileUsername(info models.User) string  {
	username := ""

	if info.RegType == 0 {
		username = strings.Replace(info.Mobile, fn.Substr(info.Mobile, 3, 4), "****", 1)
	}else{
		username = fn.NicknameByEmail(info.Email)
	}

	return username
}


func (i *UserService) GetFullUsername(info models.User) string  {
	username := ""

	if info.RegType == 0 {
		username = info.Mobile
	}else{
		username = info.Email
	}

	return username
}

func (i *UserService) GetForgotResetToken(uid int, typ int) string {
	rst := &models.UserResetPwd{}

	rst.Uid = uid
	rst.Token = crypt.SHA1(strconv.FormatInt(time.Now().UnixNano(), 10))
	rst.Expire = time.Now().Unix() + cfg.Get("env", "forgot_expire").Int()
	rst.Status = typ

	if rst.Create() {
		return rst.Token
	}

	return ""
}