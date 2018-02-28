package routes

import (
	"github.com/gin-gonic/gin"
)

//AuthGroup return auth group
func AuthGroup(r *gin.Engine) {

	//Authentication Routes...
	r.GET("/login", authCtl.ShowLoginPage)

	r.POST("/login", authCtl.Login)

	r.POST("/logout", authCtl.Logout)

	//Registration Routes...
	r.GET("/register", authCtl.Create)
	r.POST("/register", authCtl.Store)

	//Password Reset Routes...
	p := r.Group("/password")
	{
		p.GET("reset", forgotPwdCtl.ShowRequestPage)

		p.POST("email", forgotPwdCtl.SendResetLinkEmail)

		p.GET("reset/:token", forgotPwdCtl.ShowResetPage)

		p.POST("reset", forgotPwdCtl.ResetPassword)
	}
}
