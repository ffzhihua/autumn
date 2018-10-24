package main

import (
	"autumn/common/db"
	"autumn/common/redis"
	"autumn/tools/cfg"
	"autumn/routes"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"runtime"
)



func main() {


	_init()

	NumCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(int(math.Max(float64(NumCPU-1), 1)))

	routes.Bind("8282")

	defer _deferer()
}

func _init() {
	_check_config()
	_check_language()
	_check_online()

	db.InitMySQL()
	redis.InitRedis()
}

func _deferer() {
	db.CloseMySQL()
	redis.CloseRedis()
}

/*检查配置文件格式*/
func _check_config() {
	filename, _ :=cfg.WalkDir("config","")
	for _, s := range filename {

		cfg.Valid(s)
	}
	log.Println("init check config  success")
}

func _check_language() {
	filename, _ :=cfg.WalkDir("i18n","")
	for _, s := range filename {
		cfg.Valid(s)
	}

	log.Println("init check language config  success")
}

func _check_online(){
	dev := cfg.Get("env", "dev").Bool()
	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}
}