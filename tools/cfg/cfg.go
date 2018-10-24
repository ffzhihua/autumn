package cfg

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var cfgCache map[string]string = make(map[string]string)

//获取多语言配置
func Lang(typ string, lan string, path string) gjson.Result {
	filename := "i18n/" + lan + "/" + typ + ".conf"

	content := Cfg_read(filename)
	if !gjson.Valid(content) {
		return gjson.Result{}
	}

	return gjson.Get(content, path)
}

//解析配置文件
func Get(conf string, path string) gjson.Result {
	filename := "config/" + conf + ".conf"

	content := Cfg_read(filename)
	if !gjson.Valid(content) {
		return gjson.Result{}
	}

	return gjson.Get(content, path)
}

//验证配置文件
func Valid(filename string) {

	content := Cfg_read(filename)
	if !gjson.Valid(content) {
		//配置文件格式不正确，请检查配置文件
		panic("配置文件格式不正确，请检查配置文件:" + filename)
	}

	return
}

func Cfg_read(filename string) string {
	content := cache_get(filename)
	if content != "" {
		return content
	} else {
		ret, _ := cache_in(filename)

		if ret == false {
			panic("config file not found:" + filename)
		}

		data, e := ioutil.ReadFile(filename)
		if e != nil {
			panic("read config file {" + filename + "} error:" + e.Error())
		}

		content = string(data)
		cache_set(filename, content)
	}

	return content
}

func cache_get(key string) string {
	value, exists := cfgCache[key]
	if exists {
		return value
	}

	return ""
}

func cache_set(key string, val string) {
	_, exists := cfgCache[key]
	if exists {
		cfgCache[key] = val
		return
	}

	cfgCache[key] = val
}

func cache_in(filename string) (bool, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
