package model

//User model
type User struct {
	Model
	Name          string `gorm:"not null" json:"name" json:"name"`
	Email         string `gorm:"not null;unique_index:users_email_unique" json:"email"`
	Password      string `gorm:"not null" json:"-"`
	RememberToken string `gorm:"size:100" json:"-"`
}
