package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/controller"
)

func AuthGroup(r *gin.Engine) {

	//Authentication Routes...
	r.GET("/login", AuthControllerActionShowLoginPage)
	r.POST("/login", AuthControllerActionLogin)

	r.POST("/logout", AuthControllerActionLogout)

	//Registration Routes...
	r.GET("/register", AuthControllerActionShowRegisterPage)
	r.POST("/register", AuthControllerActionRegister)

	//Password Reset Routes...
	p := r.Group("/password")
	{
		p.GET("reset", ForgotPwdControllerActionShowRequestPage)
		p.POST("email", ForgotPwdControllerActionSendResetLinkEmail)
		p.GET("reset/:token", ForgotPwdControllerActionShowResetPage)
		p.POST("reset", ForgotPwdControllerActionReset)
	}
}
