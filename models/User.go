package models

import (
	"log"
	"github.com/jinzhu/gorm"
)


type User struct {
	BaseModel
	Id          int
	Mobile      string
	Country     string
	Email       string
	Password    string
	RegType     int
	RegIP       string
	RealName    int
	GoogleAuth  int
	LastLogin   int64
}

func (i *User) CheckByMobile(mobile string) bool {
	var info User
	i.DB().Where(User{Mobile:mobile}).First(&info)

	if info.Id > 0 {
		return true
	}

	return false
}

func (i *User) CheckByEmail(email string) bool {
	var info User
	i.DB().Where(User{Email:email}).First(&info)

	if info.Id > 0 {
		return true
	}

	return false
}

func (i *User) Table() *gorm.DB {
	return i.DB().Table("users")
}


func (i *User) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.User.Create:", db.Error)
	}

	return true
}

func (i *User) Info(uid int) (info User) {


	db := i.Table().Where(User{Id:uid}).First(&info)

	if db.Error != nil {
		log.Println("Model.User.Info:", db.Error)
	}

	return
}

func (i *User) InfoByUser(query User) (info User) {
	db := i.Table().Where(query).First(&info)

	if db.Error != nil {
		log.Println("Model.User.Info:", db.Error)
	}

	return
}

func (i *User) Update() bool {

	db := i.Table().Where(User{Id:i.Id}).Save(i)
	if db.Error != nil {
		log.Println("Model.User.Update:", db.Error)
		return false
	}

	return true
}

//检测是否有其他人验证了这个邮箱
func (i *User) CheckVerifyEmail(email string) bool {
	var info User
	i.Table().Where("email = ? and id <> ?", email, i.Id).First(&info)
	if info.Id > 0 {
		return true
	}

	return false
}

//检测是否有其他人验证了这个手机
func (i *User) CheckVerifyMobile(mobile string) bool {
	var info User
	i.Table().Where("mobile = ? and id <> ?", mobile, i.Id).First(&info)
	if info.Id > 0 {
		return true
	}

	return false
}