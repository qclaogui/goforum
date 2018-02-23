package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/controller"
)

func AuthGroup(r *gin.Engine) {

	//Authentication Routes...
	r.GET("/login", authCtl.ShowLoginPage)

	r.POST("/login", authCtl.Login)

	r.POST("/logout", authCtl.Logout)

	//Registration Routes...
	r.GET("/register", authCtl.Create)
	r.POST("/register", authCtl.Store)

	//Password Reset Routes...TODO//
	p := r.Group("/password")
	{
		p.GET("reset", ForgotPwdControllerActionShowRequestPage)
		p.POST("email", ForgotPwdControllerActionSendResetLinkEmail)
		p.GET("reset/:token", ForgotPwdControllerActionShowResetPage)
		p.POST("reset", ForgotPwdControllerActionReset)
	}
}
