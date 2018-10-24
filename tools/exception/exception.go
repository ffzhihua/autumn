package exception

import (
"github.com/gin-gonic/gin"

"fmt"
"reflect"
"runtime"
"runtime/debug"
	"autumn/tools/i18n"
	"strconv"
)

//异常捕获
func Catch(c *gin.Context) {

	var data = map[string]string{}

	if err := recover(); err != nil {
		debug.PrintStack()

		fmt.Println("Exception:", err, reflect.TypeOf(err))
		code := 99999
		msg := i18n.Get("code", strconv.Itoa(code), i18n.Lang(c))

		switch err.(type) {
		case *runtime.TypeAssertionError:
			c.AbortWithStatusJSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
		case *runtime.Error:
			c.AbortWithStatusJSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
		case string:
			c.AbortWithStatusJSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
		case int:
			c.AbortWithStatusJSON(200, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
		default:
			c.AbortWithStatusJSON(200, gin.H{
				"code": err,
				"msg":  i18n.Get("code", err.(string), i18n.Lang(c)),
				"data": data,
			})
		}
	}

	return
}
