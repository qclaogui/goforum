package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
)

//ExceptHandles 表示需要放行的handlerFunc
var ExceptHandles = make(map[string]bool, 50)

// JwtAuthMiddleware jwt token 验证
func JwtAuthMiddleware(except ...gin.HandlerFunc) gin.HandlerFunc {
	for _, v := range except {
		ExceptHandles[runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()] = true
	}

	return func(c *gin.Context) {
		var data *model.Payload
		var err error
		//根据请求头来判断返回的数据格式
		if "application/json" != c.GetHeader("Content-Type") {

			session := sessions.Default(c)
			token, _ := session.Get("jwt-token").(string)
			session.Save()

			data, err = model.ValidateAuthToken(token)

			if !ExceptHandles[c.HandlerName()] {
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/login")
					fmt.Fprintf(gin.DefaultErrorWriter, "invalid token:%v", err.Error())
					return
				}
			}
		} else {
			//API注意请求头为"application/json"
			if data, err = model.ValidateAuthToken(strings.TrimSpace(c.GetHeader("token"))); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error_code": 10000,
					"message":    "invalid token",
				})
				fmt.Fprintf(gin.DefaultErrorWriter, "invalid token:%v", err.Error())
				return
			}
		}

		c.Set("payload", data)
		c.Next()
	}

}
