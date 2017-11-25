/*
|--------------------------------------------------------------------------
| JwtAuthMiddleware
|--------------------------------------------------------------------------
|
| init database
|
*/
package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/model"
)

var ShouldCheckHandles = make(map[string]bool, 50)

/**
 * jwt token 验证
 *
 * except 表示需要放行的handlerFunc
 */
func JwtAuthMiddleware(except ...gin.HandlerFunc) gin.HandlerFunc {
	for _, v := range except {
		ShouldCheckHandles[runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()] = false
	}

	return func(c *gin.Context) {
		//同一个请求只验证一次就好了
		if c.GetBool("JwtAuthMiddleware") {
			c.Next()
			return
		}

		var data *Payload
		var err error
		//根据请求头来判断返回的数据格式
		if "application/json" != c.GetHeader("Content-Type") {

			token, _ := sessions.Default(c).Get("jwt-token").(string)
			data, err = ValidateAuthToken(token)

			//是否需要验证这个handles
			if ShouldCheckHandles[c.HandlerName()] {
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/login")
					fmt.Fprintf(gin.DefaultErrorWriter, "invalid token:%v", err.Error())
					return
				}
			}

		} else {

			//API注意请求头为"application/json"
			if data, err = ValidateAuthToken(strings.TrimSpace(c.GetHeader("token"))); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error_code": 10000,
					"message":    "invalid token",
				})
				fmt.Fprintf(gin.DefaultErrorWriter, "invalid token:%v", err.Error())
				return
			}
		}

		c.Set("JwtAuthMiddleware", true)
		c.Set("payload", data)
		c.Next()
	}

}
