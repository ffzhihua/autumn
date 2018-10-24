package middleware

import (
"github.com/gin-gonic/gin"

	"autumn/tools/i18n"
	"autumn/tools/token"
	"autumn/result"
	"autumn/tools/cfg"
)

func ForceJwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证token，返回错误码和uid
		code, uid := token.Verify(c.GetHeader("Auth-Token"))

		//dev开启是测试专用通道
		if cfg.Get("env", "dev").Bool() && uid == "" {
			uid = c.GetHeader("uid")
		}

		if uid == "" {
			c.AbortWithStatusJSON(200, gin.H{
				"code": code,
				"msg":  i18n.Get("code", code, i18n.Lang(c)),
				"data": result.Fail{},
			})

			return
		}


		data := make(map[string]interface{})
		data["uid"] = uid

		c.Keys = data

		c.Next()

	}
}


func NotForceJwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证token，返回错误码和uid
		_, uid := token.Verify(c.GetHeader("Auth-Token"))

		//dev开启是测试专用通道
		if cfg.Get("env", "dev").Bool() && uid == "" {
			uid = c.GetHeader("uid")
		}

		data := make(map[string]interface{})
		data["uid"] = uid

		c.Keys = data

		c.Next()

	}
}