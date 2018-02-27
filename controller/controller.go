/*
|--------------------------------------------------------------------------
| Controller
|--------------------------------------------------------------------------
|
| This controller contains common func for controller
|
*/
package controller

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/config"
	"github.com/qclaogui/goforum/middleware"
	. "github.com/qclaogui/goforum/model"
	"github.com/spf13/viper"
)

var forumC *config.APP

func init() {
	forumC = config.AppConfig
}

//return forum database
func forumDB() *gorm.DB {
	return forumC.DB
}

/**
 * the current user is authenticated.
 *
 * @return user Payload
 */
func AuthUser(c *gin.Context) (*Payload, bool) {

	p, exists := c.Get("payload")
	if !exists {
		return nil, false
	}

	data, ok := p.(*Payload)
	if !ok || (data == nil) {
		return nil, false
	}

	return data, true
}

func CsrfField(c *gin.Context) template.HTML {
	return template.HTML(fmt.Sprintf("<input type=%q name=%q value=%q>", "hidden", "_token", middleware.CsrfToken(c)))
}

func CsrfTokenValue(c *gin.Context) string {

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
	v := viper.New()
	v.SetConfigName("mix-manifest")
	v.AddConfigPath(filepath.Join(os.Getenv("GOPATH"), "src/github.com/qclaogui/goforum/public"))
	v.SetConfigType("json")
	v.ReadInConfig()
	return forumC.Config.GetString("APP_URL") + "/assets" + v.GetString("/"+s)
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
