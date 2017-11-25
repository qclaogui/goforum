package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/controller"
	. "github.com/qclaogui/goforum/middleware"
)

func WebGroup(r *gin.Engine) {

	r.GET("", WelcomeControllerActionIndex)
	r.GET("/home", HomeControllerActionIndex)

	thread := r.Group("/t", JwtAuthMiddleware(
		ThreadControllerActionShow,
		ThreadControllerActionIndex,
	))
	{
		thread.GET("", ThreadControllerActionIndex)
		thread.GET("show/:tid", ThreadControllerActionShow)

		//Add Thread
		thread.GET("create", ThreadControllerActionShowCreatePage)
		thread.POST("store", ThreadControllerActionStore)

		//Edit Thread
		thread.GET("edit/:tid", ThreadControllerActionShowEditPage)
		thread.POST("edit/:tid", ThreadControllerActionEdit)

		//Delete Thread
		thread.POST("delete/:tid", ThreadControllerActionDestroy)

	}

	return
}
