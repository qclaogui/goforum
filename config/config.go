package config

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	. "github.com/qclaogui/goforum/model"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var AppConfig *APP

type APP struct {
	Config *viper.Viper
	DB     *gorm.DB
}

func init() {
	AppConfig = LoadConfig()
}

func LoadConfig() *APP {
	a := new(APP)

	a.readConfig()

	a.initDB()

	return a
}

func (a *APP) readConfig() {
	v := viper.New()
	v.SetConfigName("app")
	v.AddConfigPath(filepath.Join(os.Getenv("GOPATH"), "src/github.com/qclaogui/goforum"))
	v.SetConfigType("yaml")
	v.ReadInConfig()

	a.Config = v
}

func (a *APP) initDB() {

	dbType := a.Config.GetString("DB_CONNECTION")
	DSN := a.Config.GetString("DB_USERNAME") + ":" + a.Config.GetString("DB_PASSWORD") + "@tcp(" +
		a.Config.GetString("DB_HOST") + ":" + a.Config.GetString("DB_PORT") + ")/" +
		a.Config.GetString("DB_DATABASE") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(dbType, DSN)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&User{},
		&Thread{},
		&Channel{},
		&Reply{},
	)

	a.DB = db
}
