package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var running = true

func main() {
	fmt.Println("Start 3rd api")

	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		if !running {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.POST("/toggle", func(c *gin.Context) {
		running = !running

		c.JSON(http.StatusOK, gin.H{"message": "ok", "status": running})
	})

	router.Run(":3000")
}
