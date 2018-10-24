package db

import (
	"autumn/tools/cfg"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"strings"

	//"strings"
)

var DB map[string]*gorm.DB

func InitMySQL() {

	filename:= "config/mysql.conf"
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
	if err != nil {
		log.Fatal("[GIN-MYSQL(" + project + ")] connect to mysql error:" + err.Error())
	}

	log.Println("[GIN-MYSQL(" + project + ")] connected success")

	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "[GIN-MYSQL("+project+")]", 0))

	return db
}
