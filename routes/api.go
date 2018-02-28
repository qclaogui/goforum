package routes

import (
	"github.com/gin-gonic/gin"
)

//APIGroup return api group
func APIGroup(r *gin.Engine) {
	api := r.Group("api/v1")
	{
		api.GET("t", threadCtl.Index)
		api.GET("t/v/:tid", threadCtl.Show)
	}
	return
}
