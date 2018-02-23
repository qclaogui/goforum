/*
|--------------------------------------------------------------------------
| Welcome Controller
|--------------------------------------------------------------------------
|
| This controller is WelcomePage
|
*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WelcomeController struct{}

//Home page
func (w *WelcomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome/index.html", gin.H{
		"host":       "http://" + c.Request.Host,
		"title":      "Welcome",
		"content":    "Let's go a forum",
		"ginContext": c,
	})
}
