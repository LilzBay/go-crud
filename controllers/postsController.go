package controllers

import (
	"log"
	"net/http"

	"github.com/LilzBay/go-crud/initializers"
	"github.com/LilzBay/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var reqInfo struct {
		Title string
		Body  string
	}
	c.Bind(&reqInfo)

	// create a post
	post := models.Post{
		Title: reqInfo.Title,
		Body:  reqInfo.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		log.Println(result.Error)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get `id` off URL
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id) // 使用`主键`查询

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var reqInfo struct {
		Title string
		Body  string
	}
	c.Bind(&reqInfo)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: reqInfo.Title,
		Body:  reqInfo.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{
		"deleteID": id,
	})
}
