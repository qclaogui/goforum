/*
|--------------------------------------------------------------------------
| Thread Controller
|--------------------------------------------------------------------------
|
| This controller is Thread(add, edit,delete)
|
*/
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/qclaogui/goforum/model"
)

type ThreadController struct{}

//get All Threads
func (t *ThreadController) Index(c *gin.Context) {

	var threads []Thread

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

func (t *ThreadController) Show(c *gin.Context) {
	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	thread, err := (&Thread{}).FindById(c, forumC.DB, "User", "Reply")

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

//ShowCreatePage
func (t *ThreadController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "thread/create.html", gin.H{
		"title":      "Welcome go forum",
		"content":    "You are logged in!",
		"ginContext": c,
	})
}

func (t *ThreadController) Store(c *gin.Context) {

	if err := ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	(&Thread{}).Create(c, forumC.DB)

	c.Redirect(http.StatusFound, "/t")
}

//ShowEditPage
func (t *ThreadController) Edit(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	thread, err := (&Thread{}).FindById(c, forumC.DB)
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

func (t *ThreadController) Update(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := ValidatePostFromParams(c, "title", "body"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := (&Thread{}).Edit(c, forumC.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/t")
}

func (t *ThreadController) Destroy(c *gin.Context) {

	if err := ValidateParams(c, "tid"); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := (&Thread{}).DestroyById(c, forumC.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/t")
}
