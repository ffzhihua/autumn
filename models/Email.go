package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Email struct {
	BaseModel
	Id        int
	Email     string
	Code      string
	Status    int
	Expire    int64
}

func (i *Email) Table() *gorm.DB {
	return i.DB().Table("email")
}

func (i *Email) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		fmt.Println("Model.Email.Create", db.Error)
		return false
	}

	return true
}

func (i *Email) Info(email string, code string) (info Email) {
	db := i.Table().Where(&Email{Email: email, Code:code}).Order("id DESC").First(&info)

	if db.Error != nil {
		log.Println("Model.Email.Info:", db.Error)
	}

	return
}

func (i *Email) Update() bool {
	db := i.Table().Where(Email{Id:i.Id}).Save(i)
	if db.Error != nil {
		log.Println("Model.Email.Update:", db.Error)
		return false
	}

	return true
}
