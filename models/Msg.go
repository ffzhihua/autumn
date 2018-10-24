package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Msg struct {
	BaseModel
	Id      int     `json:"msg_id"`
	Uid     int     `json:"uid" gorm:"column:receiver_id"`
	Title   string  `json:"title"`
	Content string  `json:"content" gorm:"column:message"`
	Status  int     `json:"status"`
}

func (i *Msg) Table() *gorm.DB {
	return i.DB().Table("msg")
}

func (i *Msg) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.Msg.Create:", db.Error)
	}

	return true
}

func (i *Msg) List(uid, limit, offset int) (list []Msg) {
	db := i.Table().Where("receiver_id in (?)", []int{5, 0}).
		Order("id desc").Limit(limit).Offset(offset).Find(&list)

	if db.Error != nil {
		log.Println("Model.Msg.List:", db.Error)
	}

	return
}

func (i *Msg) Info(id int) (info Msg) {
	db := i.Table().Where(Msg{Id:id}).First(&info)

	if db.Error != nil {
		log.Println("Model.Msg.Info:", db.Error)
	}

	return
}

func (i *Msg) Update() bool {
	db := i.Table().Where(Msg{Id:i.Id}).Save(i)

	if db.Error != nil {
		log.Println("Model.Msg.Update:", db.Error)
		return false
	}

	return true
}