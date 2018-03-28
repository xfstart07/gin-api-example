package main

import (
	"gin-api-example/handler"
	"gin-api-example/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.ApiVersion())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go",
		})
	})

	articleHandler := handler.ArticleHandler{}
	router.GET("/articles", articleHandler.Index)

	router.Run(":8888")
}
