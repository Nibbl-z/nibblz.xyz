package routes

import (
	"bufio"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostMeta struct {
	Name string;
	Description string;
	Timestamp int;
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
		
		posts = append(posts, PostMeta {
			Name: name,
			Description: description,
			Timestamp: timestamp,
			ID: file.Name(),
		})
	}

	result := Response {
		Posts: posts,
	}

	c.JSON(http.StatusOK, result)
}