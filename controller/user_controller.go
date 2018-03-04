package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/repository/mysql"
)

var us *mysql.UserRepository

func init() {
	// Create services.
	us = &mysql.UserRepository{GormDB: forumC.DB}
}

//UserController return home page
type UserController struct{}

//Show return user Home page
func (u *UserController) Show(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg":            "ok",
		"FindByID":       us.FindByID("1"),
		"FindByName":     us.FindByName("Go"),
		"FindByUsername": us.FindByUsername("go@qq.com"),
	})
}
