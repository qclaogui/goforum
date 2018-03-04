package goforum

import (
	"github.com/qclaogui/goforum/config"
	"github.com/qclaogui/goforum/routes"
)

//forum config
var forumC *config.APP

func init() {
	forumC = config.AppConfig
}

//Run start forum server
func Run() {
	routes.InitRoutes().Run(":" + forumC.Config.GetString("APP_SERVER_PORT"))

}
