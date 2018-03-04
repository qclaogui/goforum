package mysql

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/model"
)

// ReplyRepository mysql ReplyRepository
type ReplyRepository struct {
	GormDB *gorm.DB
}

// FindByID returns a user for a given id.
func (rs *ReplyRepository) FindByID(id string) *model.Reply {

	r := model.Reply{}

	rs.GormDB.Debug().First(&r, id)

	return &r
}

// Store a new Reply.
func (rs *ReplyRepository) Store(c *gin.Context) *model.Reply {

	tid, _ := strconv.ParseUint(c.Param("tid"), 10, 0)
	userID := c.GetInt("user_id")
	r := model.Reply{
		ThreadID: uint(tid),
		UserID:   uint(userID),
		Body:     c.PostForm("body"),
	}

	if rs.GormDB.NewRecord(r) {
		rs.GormDB.Create(&r)
	}

	return &r
}
