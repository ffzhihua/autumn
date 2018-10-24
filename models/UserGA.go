package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
)

type UserGA struct {
	BaseModel
	Uid int
	Secret string
	Status int
}

func (i *UserGA) Table() *gorm.DB {
	return i.DB().Table("user_ga")
}

func (i *UserGA) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		fmt.Println("Model.UserGA.Create", db.Error)
		return false
	}

	return true
}

func (i *UserGA) Info(uid int) (info UserGA) {
	i.Table().Where(UserGA{Uid:uid}).First(&info)

	return
}

func (i *UserGA) Update() bool {

	db := i.Table().Where(UserGA{Uid:i.Uid}).Update(i)
	if db.Error != nil {
		log.Println("Model.UserGA.Update:", db.Error)
		return false
	}

	return true
}
