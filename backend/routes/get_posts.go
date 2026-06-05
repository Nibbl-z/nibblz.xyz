package routes

import (
	"bufio"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostMeta struct {
	Name string;
	Description string;
	Timestamp int;
	Image string;
	ID string;
}

type Response struct {
	Posts []PostMeta
}

func GetPosts(c *gin.Context) {
	files, err := os.ReadDir("posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "failed to fetch posts! - " + err.Error(),
		})
		return
	}

	var posts []PostMeta

	for _, file := range files {
		opened, err := os.Open("posts/" + file.Name())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message" : "failed to read post data! - " + err.Error(),
			})
			return
		}

		scanner := bufio.NewScanner(opened)
		scanner.Scan()
		name := scanner.Text()
		scanner.Scan()
		description := scanner.Text()
		scanner.Scan()
		timestamp, err := strconv.Atoi(scanner.Text())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message" : "failed to read post data! - " + err.Error(),
			})
			return
		}

		opened.Close()

		image := "http://localhost:5005/thumbnails/" + file.Name() + ".png"

		_, err = os.Stat("post_thumbnails/" + file.Name() + ".png")
		if errors.Is(err, os.ErrNotExist) {
			image = "/biribiriuo.webp" // todo: proper placeholder
		}
		
		posts = append(posts, PostMeta {
			Name: name,
			Description: description,
			Timestamp: timestamp,
			Image: image,
			ID: file.Name(),
		})
	}

	result := Response {
		Posts: posts,
	}

	c.JSON(http.StatusOK, result)
}