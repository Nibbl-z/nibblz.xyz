package routes

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")

	if !strings.Contains(id, ".md") { // something malicious is brewing. probably. idk
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "that post doesn't exist!",
		})
		return
	}

	file, err := os.Open("posts/" + id)

	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message" : "that post doesn't exist!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message" : "failed to fetch post contents! - " + err.Error(),
			})
		}
		
		return
	}

	scanner := bufio.NewScanner(file)
	// skips tha first 3 lines
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	// ok,, lets read!

	var buffer bytes.Buffer

	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteString("\n")
	}
	
	if err := scanner.Err(); err != nil {
        fmt.Printf("error reading file: %s\n", err)
		c.String(http.StatusInternalServerError, "failed to fetch post contents! - " + err.Error())
		return
    }

	fmt.Println(buffer.String())

	c.String(http.StatusOK, buffer.String())
}