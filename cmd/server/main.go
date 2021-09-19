package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dacharat/circuit-breaker-go/cmd/server/external"
	"github.com/dacharat/circuit-breaker-go/internal/httpbreaker"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

func main() {
	fmt.Println("Start server")

	client := httpbreaker.NewDefaultClient()
	thirdApi := external.NewThirdApi(client)

	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		data, err := thirdApi.GetData(c.Request.Context())
		if err != nil {
			if errors.Is(err, gobreaker.ErrOpenState) {
				c.JSON(http.StatusServiceUnavailable, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
	})

	router.Run()
}
