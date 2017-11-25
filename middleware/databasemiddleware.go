/*
|--------------------------------------------------------------------------
| DatabaseMiddleware
|--------------------------------------------------------------------------
|
| init database
|
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//初始化数据库
func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
}
