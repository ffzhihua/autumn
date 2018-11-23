package main

import (
	"arutam/common/db"
	"arutam/common/redis"
	"arutam/routes"
	"arutam/tools/cfg"
	"flag"
	log "github.com/alecthomas/log4go"
	"math"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {

	_init()
	port := cfg.Get("env", "port").String()
	if port == ""{
		port = "8282"
	}
	routes.Bind(port)


	defer _deferer()

}

func init() {

	ConfPath := flag.String("c", "", " set arutam config file path")
	flag.Parse()
	cfg.ConfPath = string(*ConfPath)

	NumCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(int(math.Max(float64(NumCPU-1), 1)))

	log.LoadConfiguration(cfg.ConfPath + "config/log.xml")

}

func _init() {
	_check_config()
	_check_language()
	_check_online()

	db.InitMySQL()
	myredis.InitRedis()
}

func _deferer() {
	db.CloseMySQL()
	myredis.CloseRedis()
	log.Close()
}

/*检查配置文件格式*/
func _check_config() {
	filename, _ := cfg.WalkDir("config", "")
	for _, s := range filename {

		cfg.Valid(s)
	}
	log.Info("init check config  success")
}

func _check_language() {
	filename, _ := cfg.WalkDir("i18n", "")
	for _, s := range filename {
		cfg.Valid(s)
	}

	log.Info("init check language config  success")
}

func _check_online() {
	dev := cfg.Get("env", "dev").Bool()
	if !dev {
		gin.SetMode(gin.ReleaseMode)

	}
}
