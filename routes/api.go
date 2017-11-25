package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/controller"
)

func ApiGroup(r *gin.Engine) {
	api := r.Group("api/v1")
	{
		api.GET("t", ThreadControllerActionIndex)
		api.GET("t/v/:tid", ThreadControllerActionShow)
	}
	return
}
