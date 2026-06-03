package main

import (
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get_posts", routes.GetPosts)
	router.GET("/get_post/:id", routes.GetPost)

	router.Run(":5005")
}