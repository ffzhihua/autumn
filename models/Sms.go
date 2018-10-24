package models

import (
	"log"
	"github.com/jinzhu/gorm"
)

type Sms struct {
	BaseModel
	Id          int
	Mobile      string
	Country     string
	Code        string
	Status      int
	Expire      int64
}

func (i *Sms) Table() *gorm.DB {
	return i.DB().Table("sms")
}

func (i *Sms) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.Sms.Create:", db.Error)
		return false
	}

	return true
}

func (i *Sms) Info(mobile string, code string) (info Sms) {
	i.Table().Where(Sms{Mobile:mobile,Code:code}).Order("id DESC").Find(&info)
	return
}

func (i *Sms) Update() bool {
	db := i.Table().Where(Sms{Id:i.Id}).Save(i)
	if db.Error != nil {
		log.Println("Model.Sms.Update:", db.Error)
		return false
	}

	return true
}