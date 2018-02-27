package routes

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/controller"
	. "github.com/qclaogui/goforum/middleware"
)

var welcomeCtl = new(WelcomeController)
var threadCtl = new(ThreadController)
var homeCtl = new(HomeController)
var authCtl = new(AuthController)

func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"csrf_field": CsrfField,
		"csrf_token": CsrfTokenValue,
		"Check":      AuthCheck,
		"mix":        Mix,
	})

	respath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/qclaogui/goforum")
	r.StaticFile("/favicon.ico", respath+"/public/favicon.ico")
	r.Static("/assets", respath+"/public")
	r.LoadHTMLGlob(respath + "/resources/views/**/*")

	//session
	store := sessions.NewCookieStore([]byte("forum"))
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		sessions.Sessions("FID", store),
		JwtAuthMiddleware(
			welcomeCtl.Index,
			authCtl.ShowLoginPage,
			authCtl.Create,
		),
		VerifyCsrfToken(),
	)

	//about login register
	AuthGroup(r)

	//web
	WebGroup(r)

	//api
	ApiGroup(r)

	//WebSockets
	WebSocketGroup(r)

	return r
}
