package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Thread model
type Thread struct {
	Model
	UserID       uint `gorm:"not null" json:"user_id"`
	User         User
	Replies      []Reply
	ChannelID    uint   `gorm:"not null" json:"channel_id"`
	RepliesCount uint   `json:"replies_count"`
	Title        string `gorm:"not null" json:"title"`
	Body         string `gorm:"not null;type:text" json:"body"`
}

//ThreadPath return a path
func (t *Thread) ThreadPath() string {
	return "/t/show"
}

//WithUser load User data
func (t *Thread) WithUser(db *gorm.DB) {
	var u User
	db.Debug().Where("id=?", t.UserID).Limit(1).Find(&u)
	t.User.ID = u.ID
	t.User.Name = u.Name
	t.User.Email = u.Email
}

//WithReplies load Replies data
func (t *Thread) WithReplies(db *gorm.DB) {
	var replies []Reply
	db.Debug().Where("thread_id=?", t.ID).Limit(1000).Find(&replies)
	t.Replies = replies
}

//Create new Thread
func (t *Thread) Create(c *gin.Context, db *gorm.DB) {

	t.UserID = 1
	t.Title = c.PostForm("title")
	t.Body = c.PostForm("body")

	db.Create(t)
}

//Edit a Thread
func (t *Thread) Edit(c *gin.Context, db *gorm.DB) error {

	t, err := t.FindByID(c, db)
	if err != nil {
		return err
	}

	t.Title = c.PostForm("title")
	t.Body = c.PostForm("body")

	err = db.Save(t).Error

	if err != nil {
		return err
	}

	return nil
}

//FindByID get Thread by thread id
func (t *Thread) FindByID(c *gin.Context, db *gorm.DB, with ...string) (*Thread, error) {

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

//DestroyByID delete Thread by thread id
func (t *Thread) DestroyByID(c *gin.Context, db *gorm.DB) error {

	t, err := t.FindByID(c, db)
	if err != nil {
		return err
	}

	err = db.Debug().Delete(t).Error

	if err != nil {
		return err
	}

	return nil
}
