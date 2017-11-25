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
	. "github.com/qclaogui/goforum/model"
)

//forum database
func forumDB(c *gin.Context) *gorm.DB {
	return c.MustGet("database").(*gorm.DB)
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

/**
 * Determine if the current user is authenticated.
 *
 * @return bool
 */
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
func authGuest(c *gin.Context) bool {

	if _, ok := AuthUser(c); ok {
		return false
	}

	return true
}
