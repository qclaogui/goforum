package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/model"
)

// ThreadRepository  mysql threadRepository
type ThreadRepository struct {
	GormDB *gorm.DB
}

// FindByID returns a Thread for a given id from mysql.
func (ts *ThreadRepository) FindByID(id string) *model.Thread {

	t := model.Thread{}

	ts.GormDB.Debug().First(&t, id)

	return &t
}

// FindByName returns a Thread for a given name.
func (ts *ThreadRepository) FindByName(name string) *model.Thread {

	t := model.Thread{}

	ts.GormDB.Debug().Where("name = ?", name).Find(&t)

	return &t
}

// FindAll returns all threads.
func (ts *ThreadRepository) FindAll() []*model.Thread {

	var t []*model.Thread

	ts.GormDB.Debug().Find(&t)

	return t
}

//Store user store thread
func (ts *ThreadRepository) Store(c *gin.Context) *model.Thread {

	userID := c.GetInt("user_id")

	t := model.Thread{
		ChannelID: 1,
		UserID:    uint(userID),
		Title:     c.PostForm("title"),
		Body:      c.PostForm("body"),
	}

	if ts.GormDB.NewRecord(t) {
		ts.GormDB.Create(&t)
	}

	return &t
}

//Update user Update thread
func (ts *ThreadRepository) Update(c *gin.Context) *model.Thread {

	t := ts.FindByID(c.Param("tid"))

	ts.GormDB.Model(&t).Updates(model.Thread{Title: c.PostForm("title"), Body: c.PostForm("body")})

	return t
}

// Destroy destroy a thread.
func (ts *ThreadRepository) Destroy(c *gin.Context) {

	t := ts.FindByID(c.Param("tid"))

	ts.GormDB.Debug().Delete(&t)

}

// Threads returns all threads.
func (ts *ThreadRepository) Threads() []*model.Thread {

	return ts.FindAll()

}
