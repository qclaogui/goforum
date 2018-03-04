package controller

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/config"
	"github.com/qclaogui/goforum/middleware"
	"github.com/qclaogui/goforum/model"
	"github.com/spf13/viper"
)

//forum config
var forumC *config.APP

func init() {
	forumC = config.AppConfig
}

// CurrentUser return a login user payload data
func CurrentUser(c *gin.Context) *model.Payload {

	p, exists := c.Get("payload")
	if !exists {
		return nil
	}

	data, _ := p.(*model.Payload)

	return data
}

// CurrentUserID return a login user ID
func CurrentUserID(c *gin.Context) uint {
	return CurrentUser(c).UserID
}

// CurrentUserName return a login user name
func CurrentUserName(c *gin.Context) string {
	return CurrentUser(c).Name
}

// AuthUser return a login user payload data
func AuthUser(c *gin.Context) (*model.Payload, bool) {

	p, exists := c.Get("payload")
	if !exists {
		return nil, false
	}

	data, ok := p.(*model.Payload)
	if !ok || (data == nil) {
		return nil, false
	}

	return data, true
}

//CsrfField return a input contains csrf_token
func CsrfField(c *gin.Context) template.HTML {
	return template.HTML(fmt.Sprintf("<input type=%q name=%q value=%q>", "hidden", "_token", middleware.CsrfToken(c)))
}

//CsrfTokenValue return csrf_token
func CsrfTokenValue(c *gin.Context) string {

	return fmt.Sprintf("%s", middleware.CsrfToken(c))
}

//AuthCheck return bool
func AuthCheck(c *gin.Context) bool {
	return authCheck(c)
}

// Mix Get the path to a versioned Mix file.
func Mix(s string) string {
	v := viper.New()
	v.SetConfigName("mix-manifest")
	v.AddConfigPath(filepath.Join(os.Getenv("GOPATH"), "src/github.com/qclaogui/goforum/public"))
	v.SetConfigType("json")
	v.ReadInConfig()
	return forumC.Config.GetString("APP_URL") + "/assets" + v.GetString("/"+s)
}

//authCheck check user login
func authCheck(c *gin.Context) bool {

	if _, ok := AuthUser(c); ok {
		return true
	}

	return false
}
