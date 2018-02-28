package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
)

//ThreadController deal with thread
type ThreadController struct{}

//Index get All Threads
func (t *ThreadController) Index(c *gin.Context) {

	var threads []model.Thread

	forumC.DB.Debug().Find(&threads)

	for i, v := range threads {
		v.WithUser(forumC.DB)
		threads[i] = v
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

	thread, err := (&model.Thread{}).FindByID(c, forumC.DB, "User", "Reply")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

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

	(&model.Thread{}).Create(c, forumC.DB)

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

	thread, err := (&model.Thread{}).FindByID(c, forumC.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

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

	if err := (&model.Thread{}).Edit(c, forumC.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

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

	if err := (&model.Thread{}).DestroyByID(c, forumC.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/t")
}
