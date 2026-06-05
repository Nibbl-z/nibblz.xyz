package main

import (
	"backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	router.Use(cors.New(config))

	router.RedirectTrailingSlash = true
	router.GET("/get_posts", routes.GetPosts)
	router.GET("/get_post/:id", routes.GetPost)
	router.Static("/thumbnails", "./post_thumbnails")
	router.Run(":5005")
}