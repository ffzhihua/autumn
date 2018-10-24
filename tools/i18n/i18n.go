package i18n

import (
	"autumn/tools/cfg"
	"github.com/gin-gonic/gin"
	"strings"
)

func Get(typ string,  path string, lan string) string  {
	return cfg.Lang(typ, lan, path).String()
}

func Lang(c *gin.Context) string {
	lang := c.GetHeader("language")

	switch lang {
	case "en":
	case "kor":
	default:
		lang = "cn"
	}

	return strings.ToLower(lang)
}