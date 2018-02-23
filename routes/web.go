package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/middleware"
)

func WebGroup(r *gin.Engine) {

	r.GET("", welcomeCtl.Index)

	r.GET("/home", homeCtl.Index)

	thread := r.Group("/t", JwtAuthMiddleware(
		threadCtl.Show,
		threadCtl.Index,
	))
	{
		thread.GET("", threadCtl.Index)
		thread.GET("show/:tid", threadCtl.Show)

		//Add Thread
		thread.GET("create", threadCtl.Create)
		thread.POST("store", threadCtl.Store)

		//Edit Thread
		thread.GET("edit/:tid", threadCtl.Edit)
		thread.POST("edit/:tid", threadCtl.Update)

		//Delete Thread
		thread.POST("delete/:tid", threadCtl.Destroy)

	}

	return
}
