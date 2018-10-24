package response

import (
	"github.com/gin-gonic/gin"
	"autumn/tools/i18n"
	"strconv"
	"autumn/result"
)

//输出错误
func Fail(c *gin.Context, code int) {

	c.Header("Content-Type", "application/json; charset=UTF-8")

	c.JSON(200, gin.H{
		"code": code,
		"msg":  i18n.Get("code", strconv.Itoa(code), i18n.Lang(c)),
		"data": result.Fail{},
	})
}

//成功
func Success(c *gin.Context, data interface{}) {

	c.Header("Content-Type", "application/json; charset=UTF-8")

	if data == nil {
		data = result.Success{}
	}

	code := 0
	c.JSON(200, gin.H{
		"code": code,
		"msg":  i18n.Get("code", strconv.Itoa(code), i18n.Lang(c)),
		"data": data,
	})
}
