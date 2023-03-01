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
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(http.StatusBadRequest) // 400
		log.Println(result.Error)
		return // 务必返回
	}

	// return it
	c.JSON(http.StatusOK, gin.H{ // 200
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// respond with them
	c.JSON(http.StatusOK, gin.H{ // 200
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get `id` off URL param
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id) // 使用`主键`查询，返回`第一个`符合条件的记录

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// 使用主键`id`找到数据库中的记录
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{ // 200
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
	c.JSON(http.StatusOK, gin.H{ // 200
		"deleteID": id,
	})
}
