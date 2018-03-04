package mysql

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/model"
)

//1.获取数据
//2.记录日志
//3,其他

// UserRepository represents a mysql implementation of services.UserServiceProvider.
type UserRepository struct {
	GormDB *gorm.DB
}

// FindByID returns a user for a given id.
func (us *UserRepository) FindByID(id string) *model.User {

	u := model.User{}

	us.GormDB.Debug().First(&u, id)

	return &u
}

// FindByName returns a user for a given name.
func (us *UserRepository) FindByName(name string) *model.User {

	u := model.User{}

	us.GormDB.Debug().Where("name = ?", name).First(&u)

	return &u
}

// FindByUsername returns a user login account name.
func (us *UserRepository) FindByUsername(username string) *model.User {

	u := model.User{}

	us.GormDB.Debug().Where(u.Username()+"= ?", username).First(&u)

	return &u
}

// FindAll returns all users.
func (us *UserRepository) FindAll() []*model.User {

	var u []*model.User

	us.GormDB.Debug().Find(&u)

	return u
}

// Store a new user.
func (us *UserRepository) Store(c *gin.Context) *model.User {

	u := model.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if us.GormDB.NewRecord(u) {
		u.RememberToken = model.RandomString(40)
		u.Password = model.BCryptPassword(u.Password + u.RememberToken)
		us.GormDB.Create(&u)
	}

	return &u
}

// Destroy destroy a user.
func (us *UserRepository) Destroy(u *model.User) {

	us.GormDB.Debug().Delete(u)
}
