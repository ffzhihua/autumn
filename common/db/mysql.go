package db

import (
	"arutam/tools/cfg"
	"fmt"
	log "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tidwall/gjson"
	"strings"
	//"strings"
)

var DB map[string]*gorm.DB

func InitMySQL() {

	filename := cfg.ConfPath + "config/mysql.conf"
	content := cfg.Cfg_read(filename)

	DB = make(map[string]*gorm.DB)
	dat := gjson.Parse(content).Map()
	for k := range dat {
		DB[k] = mysql_connect(strings.ToLower(k))
	}

}

func CloseMySQL() {
	for k := range DB {
		DB[k].Close()
	}
}

func mysql_connect(project string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		cfg.Get("mysql", project+".user").String(),
		cfg.Get("mysql", project+".passwd").String(),
		cfg.Get("mysql", project+".host").String(),
		cfg.Get("mysql", project+".port").String(),
		cfg.Get("mysql", project+".database").String(),
		cfg.Get("mysql", project+".charset").String(),
	)
	var err error

	db, err := gorm.Open("mysql", dsn)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(10)
	if err != nil {
		log.Error("[GIN-MYSQL(" + project + ")] connect to mysql error:" + err.Error())
	}

	log.Info("[GIN-MYSQL(" + project + ")] connected success")

	dev := cfg.Get("env", "mysql_debug").Bool()

	if dev {
		db.LogMode(true)

		db.SetLogger(logger{"[SQL(" + project + ")]"})
	}

	return db
}

/**
 *自定义gorm log 输出
 *实现gorm logger print方法
 */

type logger struct {
	flag string
}

func (logger logger) Print(value ...interface{}) {

	log.Info(value,logger.flag)
}
