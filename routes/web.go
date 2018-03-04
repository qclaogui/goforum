package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/middleware"
)

//WebGroup return web group
func WebGroup(r *gin.Engine) {

	r.GET("", welcomeCtl.Index)

	r.GET("/home", homeCtl.Index)
	r.GET("/p", userCtl.Show)

	thread := r.Group("/t", middleware.JwtAuthMiddleware(
		threadCtl.Show,
		threadCtl.Index,
	))
	{
		thread.GET("", threadCtl.Index)
		thread.GET("show/:tid", threadCtl.Show)

		//replyCtl
		//Add Thread
		thread.GET("create", threadCtl.Create)
		thread.POST("store", threadCtl.Store)

		//Edit Thread
		thread.GET("edit/:tid", threadCtl.Edit)
		thread.POST("edit/:tid", threadCtl.Update)

		//reply Thread
		thread.POST("reply/:tid", replyCtl.Store)

		//Delete Thread
		thread.POST("delete/:tid", threadCtl.Destroy)

	}

	return
}
