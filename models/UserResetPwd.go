package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type UserResetPwd struct {
	BaseModel
	Id      int
	Uid     int
	Token   string
	Status  int
	Expire  int64
	Typ     int
}

func (i *UserResetPwd) Table() *gorm.DB {
	return i.DB().Table("user_reset_pwd")
}

func (i *UserResetPwd) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.UserResetPwd.Create:", db.Error)
	}

	return true
}

func (i *UserResetPwd) Info(token string) (info UserResetPwd) {
	db := i.Table().Where(UserResetPwd{Token:token}).First(&info)

	if db.Error != nil {
		log.Println("Model.UserResetPwd.Info:", db.Error)
	}

	return
}

func (i *UserResetPwd) Update() bool {

	db := i.Table().Where(UserResetPwd{Id:i.Id}).Save(i)
	if db.Error != nil {
		log.Println("Model.UserResetPwd.Update:", db.Error)
		return false
	}

	return true
}