package model

type Channel struct {
	Model
	Name string `gorm:"not null" json:"name"`
	Slug string `gorm:"not null"`
}
