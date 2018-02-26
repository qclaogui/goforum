package controller

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/middleware"
)

func CsrfField(c *gin.Context) template.HTML {
	return template.HTML(fmt.Sprintf("<input type=%q name=%q value=%q>", "hidden", "_token", middleware.CsrfToken(c)))
}

func CsrfTokenValue(c *gin.Context) string {

	//log.Println(forumC(c).Config.GetString("DB_DATABASE"))

	return fmt.Sprintf("%s", middleware.CsrfToken(c))
}

/**
 * Determine if the current user is authenticated.
 *
 * @return bool
 */
func AuthCheck(c *gin.Context) bool {
	return authCheck(c)
}

/**
 * Get the path to a versioned Mix file.
 *
 * @return string
 */
func Mix(s string) string {
	return "assets/" + s
}

func authCheck(c *gin.Context) bool {

	if _, ok := AuthUser(c); ok {
		return true
	}

	return false
}

/**
 * Determine if the current user is a guest.
 *
 * @return bool
 */
func AuthGuest(c *gin.Context) bool {

	if _, ok := AuthUser(c); ok {
		return false
	}

	return true
}
