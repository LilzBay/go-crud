package main

import (
	"github.com/LilzBay/go-crud/controllers"
	"github.com/LilzBay/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// 创建一篇博文
	r.POST("/posts", controllers.PostsCreate)
	// 查询所有文章
	r.GET("/posts", controllers.PostsIndex)
	// 查询某篇文章
	r.GET("/posts/:id", controllers.PostsShow)
	// 更新某篇文章
	r.PUT("/posts/:id", controllers.PostsUpdate)
	// 删除某篇文章
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
