/*
|--------------------------------------------------------------------------
| Auth Controller
|--------------------------------------------------------------------------
|
| This controller handles authenticating users login
|
*/
package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/model"
)

/**
 * the current user is authenticated.
 *
 */
func AuthControllerActionLogin(c *gin.Context) {

	if err := ValidatePostFromParams(c, "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	user := &User{}
	if err := forumDB(c).Where("email=?", c.PostForm("email")).Limit(1).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": 40002,
			"message":    err.Error(),
		})
		return
	}

	if !VerifyingPassword(user.Password, c.PostForm("password")+user.RememberToken) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": 40002,
			"message":    "密码错误",
		})
		return
	}

	//jwtToken
	token, _ := GenerateJwtAuthToken(&PayloadClaims{
		Data: Payload{
			UserId: user.ID,
			Name:   user.Name,
		},
	})

	if "application/json" == c.ContentType() {
		c.JSON(http.StatusOK, gin.H{
			"error_code": 0,
			"message":    "login success",
			"token":      token,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("jwt-token", token)
	session.Save()
	c.Redirect(http.StatusFound, "/home")
}

//logout
func AuthControllerActionLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/home")
}

//register
func AuthControllerActionRegister(c *gin.Context) {

	if err := ValidatePostFromParams(c, "name", "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	db := forumDB(c)
	u := User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	if db.Where("email=?", u.Email).Limit(1).Find(&u).RecordNotFound() {
		u.RememberToken = RandomString(40)
		u.Password = BCryptPassword(u.Password + u.RememberToken)
		db.Create(&u)

		//jwtToken
		token, _ := GenerateJwtAuthToken(&PayloadClaims{
			Data: Payload{
				UserId: u.ID,
				Name:   u.Name,
			},
		})

		if "application/json" == c.ContentType() {
			c.JSON(http.StatusOK, gin.H{
				"error_code": 0,
				"message":    "register success",
				"token":      token,
			})
			return
		}

		session := sessions.Default(c)
		session.Set("jwt-token", token)
		session.Save()
		c.Redirect(http.StatusFound, "/home")

	} else {
		c.JSON(http.StatusOK, gin.H{
			"error_code": 10002,
			"message":    "user exists",
		})
		return
	}

}

//show register form
func AuthControllerActionShowRegisterPage(c *gin.Context) {

	isLogin := authCheck(c)
	if isLogin {
		c.Redirect(http.StatusFound, "/home")
		return
	}

	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"host":    "http://" + c.Request.Host,
		"css":     "http://" + c.Request.Host + "/assets/css/app.css",
		"js":      "http://" + c.Request.Host + "/assets/js/app.js",
		"isLogin": isLogin,
		"data":    c.Request.Proto,
	})
}

//show login form
func AuthControllerActionShowLoginPage(c *gin.Context) {

	isLogin := authCheck(c)
	if isLogin {
		c.Redirect(http.StatusFound, "/home")
		return
	}

	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"host":    "http://" + c.Request.Host,
		"css":     "http://" + c.Request.Host + "/assets/css/app.css",
		"js":      "http://" + c.Request.Host + "/assets/js/app.js",
		"content": "Login",
		"isLogin": isLogin,
		"checked": "checked",
	})
}
