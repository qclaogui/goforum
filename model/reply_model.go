package model

import "github.com/jinzhu/gorm"

//Reply model
type Reply struct {
	Model
	ThreadID uint
	UserID   uint `gorm:"not null" json:"user_id"`
	User     User
	Body     string
}

//With load data
func (r *Reply) With(DB *gorm.DB, with ...interface{}) *Reply {

	for _, v := range with {
		switch v {
		case v.(*User):
			r.user(DB)
		}
	}
	return r
}

func (r *Reply) user(DB *gorm.DB) { DB.Debug().Model(&r).Related(&r.User) }
