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
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/config"
	. "github.com/qclaogui/goforum/model"
)

//forum database
func forumDB(c *gin.Context) *gorm.DB {
	return forumC(c).DB
}

//forum database
func forumC(c *gin.Context) *config.APP {
	return c.MustGet("config").(*config.APP)
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
