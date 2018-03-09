package model

import (
	"github.com/gin-gonic/gin"
)

//ValidatePostFromParams 校验参数from c.PostForm不能为空
func ValidatePostFromParams(c *gin.Context, params ...string) []gin.H {
	var errors []gin.H
	for _, p := range params {
		if "" == c.PostForm(p) {
			errors = append(errors, gin.H{
				"error_code": 40004,
				"message":    p + "不能为空",
			})
		}
	}
	return errors
}

//ValidateParams 校验参数from c.Param 不能为空
func ValidateParams(c *gin.Context, params ...string) []gin.H {
	var errors []gin.H
	for _, p := range params {
		if "" == c.Param(p) {
			errors = append(errors, gin.H{
				"error_code": 40004,
				"message":    p + "不能为空",
			})
		}
	}
	return errors
}
