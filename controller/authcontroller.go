package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
)

//AuthController deal with user auth
type AuthController struct{}

//Login deal with user login
func (a *AuthController) Login(c *gin.Context) {

	if err := model.ValidatePostFromParams(c, "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	user := &model.User{}
	if err := forumC.DB.Where("email=?", c.PostForm("email")).Limit(1).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": 40002,
			"message":    err.Error(),
		})
		return
	}

	if !model.VerifyingPassword(user.Password, c.PostForm("password")+user.RememberToken) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": 40002,
			"message":    "密码错误",
		})
		return
	}

	//jwtToken
	token, _ := model.GenerateJwtAuthToken(&model.PayloadClaims{
		Data: model.Payload{
			UserID: user.ID,
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

//Logout deal with user Logout
func (a *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/home")
}

//Create deal with user Create
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

//Store deal with user Store
func (a *AuthController) Store(c *gin.Context) {

	if err := model.ValidatePostFromParams(c, "name", "email", "password"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	u := model.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	if forumC.DB.Where("email=?", u.Email).Limit(1).Find(&u).RecordNotFound() {
		u.RememberToken = model.RandomString(40)
		u.Password = model.BCryptPassword(u.Password + u.RememberToken)
		forumC.DB.Create(&u)

		//jwtToken
		token, _ := model.GenerateJwtAuthToken(&model.PayloadClaims{
			Data: model.Payload{
				UserID: u.ID,
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

//ShowLoginPage return login page
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
