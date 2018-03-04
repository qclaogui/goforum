package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
	"github.com/qclaogui/goforum/repository/mysql"
)

var rs *mysql.ReplyRepository

func init() {
	rs = &mysql.ReplyRepository{GormDB: forumC.DB}
}

//ReplyController deal with thread replies
type ReplyController struct{}

//Store store a thread reply
func (r *ReplyController) Store(c *gin.Context) {

	if err := model.ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := model.ValidatePostFromParams(c, "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.Set("user_id", int(CurrentUserID(c)))

	rs.Store(c)

	c.Redirect(http.StatusFound, "/t/show/"+c.Param("tid"))
}
