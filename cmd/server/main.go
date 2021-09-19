package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dacharat/circuit-breaker-go/internal/httpbreaker"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

func main() {
	fmt.Println("Start server")

	client := httpbreaker.NewDefaultClient()

	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		body, err := client.Get(context.Background(), "http://localhost:3000", nil)
		if err != nil {
			if errors.Is(err, gobreaker.ErrOpenState) {
				c.JSON(http.StatusServiceUnavailable, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		var data struct {
			Message string `json:"message"`
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
	})

	router.Run()
}
