//Let‘s go a forum
package main

import (
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/config"
	. "github.com/qclaogui/goforum/controller"
	. "github.com/qclaogui/goforum/middleware"
	"github.com/qclaogui/goforum/routes"
)

func main() {
	forum := InitRoutes()
	forum.Run(":8321")
}

func InitRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	//初始化数据库
	db := config.InitDatabase("mysql")
	//session
	store := sessions.NewCookieStore([]byte("forum"))
	r := gin.New()

	r.SetFuncMap(template.FuncMap{
		"csrf_field": CsrfField,
		"Check":      AuthCheck,
	})

	respath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/qclaogui/goforum")

	r.StaticFile("/favicon.ico", respath+"/resources/assets/favicon.ico")
	r.Static("/assets", respath+"/resources/assets")
	r.LoadHTMLGlob(respath + "/resources/views/**/*")

	r.Use(
		gin.Logger(),
		gin.Recovery(),
		DatabaseMiddleware(db),
		sessions.Sessions("FID", store),
		JwtAuthMiddleware(
			WelcomeControllerActionIndex,
			AuthControllerActionShowLoginPage,
			AuthControllerActionShowRegisterPage,
		),
		VerifyCsrfToken(),
	)

	//about login register
	routes.AuthGroup(r)

	//web
	routes.WebGroup(r)

	//api
	routes.ApiGroup(r)

	//WebSockets
	routes.WebSocketGroup(r)

	return r
}
