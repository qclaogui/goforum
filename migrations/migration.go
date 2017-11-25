/*
|--------------------------------------------------------------------------
| ForumTablesAutoMigrate
|--------------------------------------------------------------------------
|
| Forum Tables AutoMigrate
|
*/
package migrations

import (
	"github.com/jinzhu/gorm"
	. "github.com/qclaogui/goforum/model"
)

func ForumTablesAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Thread{},
		&Channel{},
		&Reply{},
	)
}
