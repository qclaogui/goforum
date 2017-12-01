package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Thread struct {
	Model
	UserId       uint `gorm:"not null" json:"user_id"`
	User         User
	Replies      []Reply
	ChannelId    uint   `gorm:"not null" json:"channel_id"`
	RepliesCount uint   `json:"replies_count"`
	Title        string `gorm:"not null" json:"title"`
	Body         string `gorm:"not null;type:text" json:"body"`
}

func (t *Thread) ThreadPath() string {
	return "/t/show"
}

//是否需要加载User
func (t *Thread) WithUser(db *gorm.DB) {
	var u User
	db.Debug().Where("id=?", t.UserId).Limit(1).Find(&u)
	t.User.ID = u.ID
	t.User.Name = u.Name
	t.User.Email = u.Email
}

func (t *Thread) WithReplies(db *gorm.DB) {
	var replies []Reply
	db.Debug().Where("thread_id=?", t.ID).Limit(1000).Find(&replies)
	t.Replies = replies
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
func (t *Thread) FindById(c *gin.Context, db *gorm.DB, with ...string) (*Thread, error) {

	if err := db.Where("id=?", c.Param("tid")).Limit(1).Find(t).Error; err != nil {
		return nil, err
	}

	for _, v := range with {
		switch v {
		case "User":
			t.WithUser(db)
			fmt.Println("case User")
		case "Reply":
			t.WithReplies(db)
			fmt.Println("case Reply")
		}
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
