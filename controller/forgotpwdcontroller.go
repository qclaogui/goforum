/*
|--------------------------------------------------------------------------
| ForgotPwd Controller
|--------------------------------------------------------------------------
|
| This controller is responsible for handling password reset
|
*/
package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//重置密码请求页面
func ForgotPwdControllerActionShowRequestPage(c *gin.Context) {
	//showLinkRequestForm
	c.HTML(http.StatusOK, "passwords/email.html", gin.H{
		"host":    "http://" + c.Request.Host,
		"css":     "http://" + c.Request.Host + "/assets/css/app.css",
		"js":      "http://" + c.Request.Host + "/assets/js/app.js",
		"content": "RequestForm",
	})
}

//发送重置密码邮件
func ForgotPwdControllerActionSendResetLinkEmail(c *gin.Context) {
	//	sendResetLinkEmail
	time.Sleep(2 * time.Second)
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//重置页面
func ForgotPwdControllerActionShowResetPage(c *gin.Context) {
	//ShowResetForm
	token := c.Param("token")
	if token == "" {
		token = "default token"
	}
	c.HTML(http.StatusOK, "passwords/reset.html", gin.H{
		"host":    "http://" + c.Request.Host,
		"css":     "http://" + c.Request.Host + "/assets/css/app.css",
		"js":      "http://" + c.Request.Host + "/assets/js/app.js",
		"token":   token,
		"content": "ResetForm",
	})
}

//重置
func ForgotPwdControllerActionReset(c *gin.Context) {
	c.String(http.StatusOK, c.PostForm("token"))
}
