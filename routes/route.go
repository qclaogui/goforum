package routes

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/controller"
	"github.com/qclaogui/goforum/middleware"
)

var welcomeCtl = new(controller.WelcomeController)
var threadCtl = new(controller.ThreadController)
var homeCtl = new(controller.HomeController)
var authCtl = new(controller.AuthController)
var forgotPwdCtl = new(controller.ForgotPwdController)

//InitRoutes initialized forum routes
func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"csrf_field": controller.CsrfField,
		"csrf_token": controller.CsrfTokenValue,
		"Check":      controller.AuthCheck,
		"mix":        controller.Mix,
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
		middleware.JwtAuthMiddleware(
			welcomeCtl.Index,
			authCtl.ShowLoginPage,
			authCtl.Create,
		),
		middleware.VerifyCsrfToken(),
	)

	AuthGroup(r)

	WebGroup(r)

	APIGroup(r)

	WebSocketGroup(r)

	return r
}
