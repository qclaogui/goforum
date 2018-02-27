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

type AuthController struct{}

/**
 * the current user is authenticated.
 *
 */
func (a *AuthController) Login(c *gin.Context) {

	if err := ValidatePostFromParams(c, "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	user := &User{}
	if err := forumC.DB.Where("email=?", c.PostForm("email")).Limit(1).Find(&user).Error; err != nil {
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

func (a *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/home")
}

//show register form
func (a *AuthController) Create(c *gin.Context) {

	isLogin := authCheck(c)
	if isLogin {
		c.Redirect(http.StatusFound, "/home")
		return
	}

	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"data":       c.Request.Proto,
		"ginContext": c,
	})
}

//register a new user
func (a *AuthController) Store(c *gin.Context) {

	if err := ValidatePostFromParams(c, "name", "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	u := User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	if forumC.DB.Where("email=?", u.Email).Limit(1).Find(&u).RecordNotFound() {
		u.RememberToken = RandomString(40)
		u.Password = BCryptPassword(u.Password + u.RememberToken)
		forumC.DB.Create(&u)

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

func (a *AuthController) ShowLoginPage(c *gin.Context) {

	isLogin := authCheck(c)
	if isLogin {
		c.Redirect(http.StatusFound, "/home")
		return
	}

	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"content":    "Login",
		"checked":    "checked",
		"ginContext": c,
	})
}
