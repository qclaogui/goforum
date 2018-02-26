/*
|--------------------------------------------------------------------------
| InitConfigMiddleware
|--------------------------------------------------------------------------
|
| init config
|
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/config"
)

func InitConfigMiddleware(config *config.APP) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
