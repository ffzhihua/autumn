package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type UserSysMsg struct {
	BaseModel
	Id      int
	Uid     int
	MsgId   int
}

func (i *UserSysMsg) Table() *gorm.DB {
	return i.DB().Table("user_sys_msg")
}

func (i *UserSysMsg) Create() bool {
	db := i.Table().Save(i)

	if db.Error != nil {
		log.Println("Model.UserSysMsg.Create:", db.Error)
	}

	return true
}

func (i *UserSysMsg) SysReadDict(uid int, msgid []int) (map[int]int){
	dict := make(map[int]int)
	var list [] UserSysMsg

	db := i.Table().Where(UserSysMsg{Uid:uid}).Where("msg_id in (?)", msgid).Find(&list)

	if db.Error != nil {
		log.Println("Model.UserSysMsg.List:", db.Error)
	}

	for _,v := range list {
		dict[v.MsgId] = 1
	}

	return dict
}

func (i *UserSysMsg) IsRead(uid, msgid int) bool {
	var info UserSysMsg
	db := i.Table().Where(UserSysMsg{Uid:uid,MsgId:msgid}).First(&info)

	if db.Error != nil {
		log.Println("Model.UserSysMsg.IsRead:", db.Error)
		return false
	}

	if info.Id <= 0 {
		return false
	}

	return true
}