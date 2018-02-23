package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiGroup(r *gin.Engine) {
	api := r.Group("api/v1")
	{
		api.GET("t", threadCtl.Index)
		api.GET("t/v/:tid", threadCtl.Show)
	}
	return
}
