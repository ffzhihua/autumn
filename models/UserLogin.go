package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type UserLogin struct {
	BaseModel
	Id      int
	Uid     int
	LoginIp string
}

func (i *UserLogin) Table() *gorm.DB {
	return i.DB().Table("user_login")
}

func (i *UserLogin) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.UserLogin.Create:", db.Error)
	}

	return true
}