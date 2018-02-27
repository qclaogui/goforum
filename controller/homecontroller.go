/*
|--------------------------------------------------------------------------
| Home Controller
|--------------------------------------------------------------------------
|
| This controller is auth user homePage
|
*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

//Home page
func (h *HomeController) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}
