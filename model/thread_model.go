package model

import (
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
func (t *Thread) ThreadPath() string { return "/t/show" }

//With load data
func (t *Thread) With(DB *gorm.DB, with ...interface{}) *Thread {

	for _, v := range with {
		switch v {
		case v.(*User):
			t.user(DB)
		case v.(*Reply):
			t.replies(DB)
		}
	}

	return t
}

//user load thread userInfo data
func (t *Thread) user(DB *gorm.DB) { DB.Debug().Model(&t).Related(&t.User) }

//replies load  thread Replies data
func (t *Thread) replies(DB *gorm.DB) { DB.Debug().Model(&t).Related(&t.Replies) }
