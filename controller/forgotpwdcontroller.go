package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ForgotPwdController deal with user password
type ForgotPwdController struct{}

//ShowRequestPage return request page
func (f *ForgotPwdController) ShowRequestPage(c *gin.Context) {
	c.HTML(http.StatusOK, "passwords/email.html", gin.H{
		"content":    "RequestForm",
		"ginContext": c,
	})
}

//SendResetLinkEmail send a reset password email
func (f *ForgotPwdController) SendResetLinkEmail(c *gin.Context) {

	//	TODO sendResetLinkEmail
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}

//ShowResetPage show reset page
func (f *ForgotPwdController) ShowResetPage(c *gin.Context) {
	//ShowResetForm
	token := c.Param("token")
	if token == "" {
		token = "default token"
	}
	c.HTML(http.StatusOK, "passwords/reset.html", gin.H{
		"token":      token,
		"content":    "ResetForm",
		"ginContext": c,
	})
}

//ResetPassword to reset current user password
func (f *ForgotPwdController) ResetPassword(c *gin.Context) {

	//	TODO Reset Password
	c.String(http.StatusOK, c.PostForm("token"))
}
