package routes

import (
	"github.com/gin-gonic/gin"
	"autumn/middleware"
	"autumn/controllers"
)

func Bind(port string)  {
	router := gin.Default()
	router.Use(middleware.CrosSite()).Use(middleware.NotForceJwtCheck())

	send(router)
	user(router)
	forgot(router)
	register(router)
	verify(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"ret": 404})
	})

	router.Run(":" + port)

}

func user(r *gin.Engine)  {
	user := controllers.UserController{}
	umsg := controllers.UserMsgController{}

	route := r.Group("user").Use(middleware.ForceJwtCheck())
	{
		route.PUT("password", user.UpdatePassword)

		route.GET("info", user.Info)
		route.GET("msg/:page", umsg.List)
		route.PUT("msg/:msg_id", umsg.Read)
	}

	route2 := r.Group("user")
	{
		route2.POST("login", user.Login)
	}
}

func forgot(r *gin.Engine)  {
	f := controllers.ForgotController{}

	route := r.Group("forgot")
	{
		route.POST("check/mobile", f.VerifyCodeOfMobile)
		route.POST("check/email", f.VerifyCodeOfEmail)
		route.PUT("password", f.ResetPassword)
	}
}

func send(r *gin.Engine)  {
	send := controllers.SendController{}

	route := r.Group("send")
	{
		route.POST("mobile", send.MobileCode)
		route.POST("email", send.EmailCode)
	}
}

func verify(r *gin.Engine)  {
	verify := controllers.VerifyController{}
	route := r.Group("verify").Use(middleware.ForceJwtCheck())
	{
		route.GET("ga", verify.GetGASecret)

		route.POST("ga", verify.GaConfirm)
		route.POST("email", verify.Email)
		route.POST("mobile", verify.Mobile)
	}
}

func register(r *gin.Engine) {

	reg := controllers.RegisterController{}

	//middleware.JwtCheck()
	route := r.Group("register")
	{
		route.POST("mobile", reg.Mobile)
		route.POST("email", reg.Email)
	}
}