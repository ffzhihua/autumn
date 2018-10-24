package models

import (
	"autumn/common/db"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type BaseModel struct {
	CreatedAt int64 `gorm:"type:int(11)" json:"created_at"`
	UpdatedAt int64 `gorm:"type:int(11)" json:"updated_at"`
}

//返回数据库对象
func (i *BaseModel) DB(name ...string) *gorm.DB {
	dbname := strings.ToLower(name[0])
	if dbname == "" {
		return db.DB["mysql"]
	}
	return db.DB[dbname]
}


func (i *BaseModel) BeforeCreate() {
	i.CreatedAt = time.Now().Unix()
	i.UpdatedAt = time.Now().Unix()
}

func (i *BaseModel) BeforeUpdate() {
	i.UpdatedAt = time.Now().Unix()
}

func (i *BaseModel) BeforeSave() {
	i.UpdatedAt = time.Now().Unix()
}
