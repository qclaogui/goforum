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

//Home page
func HomeControllerActionIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"host":    "http://" + c.Request.Host,
		"css":     "http://" + c.Request.Host + "/assets/css/app.css",
		"js":      "http://" + c.Request.Host + "/assets/js/app.js",
		"isLogin": authCheck(c),
		"title":   "Welcome go forum",
		"content": "You are logged in!",
	})
}
