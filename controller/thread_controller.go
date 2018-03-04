package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
	"github.com/qclaogui/goforum/repository/mysql"
)

//数据源
var ts *mysql.ThreadRepository

func init() {
	ts = &mysql.ThreadRepository{GormDB: forumC.DB}
}

//ThreadController deal with thread
type ThreadController struct{}

//Index get All Threads
func (t *ThreadController) Index(c *gin.Context) {

	threads := ts.FindAll()

	for _, v := range threads {
		v.With(ts.GormDB, &model.User{})
	}

	if "application/json" == c.ContentType() {
		c.JSON(http.StatusOK, gin.H{
			"data": threads,
		})
		return
	}

	c.HTML(http.StatusOK, "thread/index.html", gin.H{
		"threads":    threads,
		"ginContext": c,
	})
}

//Show a Thread
func (t *ThreadController) Show(c *gin.Context) {
	if err := model.ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	//从数据源服务获取数据
	thread := ts.FindByID(c.Param("tid")).With(ts.GormDB, &model.Reply{}, &model.User{})

	if "application/json" == c.ContentType() {
		c.JSON(http.StatusOK, gin.H{
			"data": thread,
		})
		return
	}

	c.HTML(http.StatusOK, "thread/show.html", gin.H{
		"thread":     thread,
		"ginContext": c,
	})
}

//Create return thread create page
func (t *ThreadController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "thread/create.html", gin.H{
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}

//Store a thread form request
func (t *ThreadController) Store(c *gin.Context) {

	if err := model.ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.Set("user_id", int(CurrentUserID(c)))

	ts.Store(c)

	c.Redirect(http.StatusFound, "/t")
}

//Edit a thread form request
func (t *ThreadController) Edit(c *gin.Context) {

	if err := model.ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	thread := ts.FindByID(c.Param("tid"))

	c.HTML(http.StatusOK, "thread/edit.html", gin.H{
		"thread":     thread,
		"ginContext": c,
	})
}

//Update a thread form request
func (t *ThreadController) Update(c *gin.Context) {

	if err := model.ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := model.ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ts.Update(c)

	c.Redirect(http.StatusFound, "/t")
}

//Destroy a thread form request
func (t *ThreadController) Destroy(c *gin.Context) {

	if err := model.ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ts.Destroy(c)

	c.Redirect(http.StatusFound, "/t")
}
