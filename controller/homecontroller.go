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
		"host":       "http://" + c.Request.Host,
		"css":        "http://" + c.Request.Host + "/assets/css/forum.css",
		"js":         "http://" + c.Request.Host + "/assets/js/forum.js",
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}
