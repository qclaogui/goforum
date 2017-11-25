package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Thread struct {
	Model
	UserId       uint   `gorm:"not null" json:"user_id"`
	ChannelId    uint   `gorm:"not null" json:"channel_id"`
	RepliesCount uint   `json:"replies_count"`
	Title        string `gorm:"not null" json:"title"`
	Body         string `gorm:"not null;type:text" json:"body"`
}

func (t *Thread) ThreadPath() string {
	return "/t/show"
}

//
func (t *Thread) Replies(db *gorm.DB) (replies []Reply) {
	db.Debug().Where("thread_id=?", t.ID).Find(&replies)
	return
}

//add Thread
func (t *Thread) Create(c *gin.Context, db *gorm.DB) {

	t.UserId = 1
	t.Title = c.PostForm("title")
	t.Body = c.PostForm("body")

	db.Create(t)
}

//Edit Thread
func (t *Thread) Edit(c *gin.Context, db *gorm.DB) error {

	t, err := t.FindById(c, db)
	if err != nil {
		return err
	}

	t.Title = c.PostForm("title")
	t.Body = c.PostForm("body")

	if err = db.Save(t).Error; err != nil {
		return err
	}

	return nil
}

//get Thread by thread id
func (t *Thread) FindById(c *gin.Context, db *gorm.DB) (*Thread, error) {

	if err := db.Where("id=?", c.Param("tid")).Limit(1).Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

//delete Thread by thread id
func (t *Thread) DestroyById(c *gin.Context, db *gorm.DB) error {

	t, err := t.FindById(c, db)
	if err != nil {
		return err
	}

	if err = db.Debug().Delete(t).Error; err != nil {
		return err
	}

	return nil
}
