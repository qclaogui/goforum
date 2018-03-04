package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomeController user home page
type HomeController struct{}

//Index return user Home page
func (h *HomeController) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}
