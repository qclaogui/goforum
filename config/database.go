package config

import (
	"github.com/jinzhu/gorm"
	"github.com/qclaogui/goforum/migrations"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDatabase(dbType string) *gorm.DB {
	db, err := gorm.Open(dbType, "root:zzz123@tcp(127.0.0.1:3306)/gogogo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	migrations.ForumTablesAutoMigrate(db)
	return db
}
