package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
)

type echo struct {
	Value string `json:"value"`
}

type payload struct {
	Action string `json:"action"`
}

func InitApi() *http.Server {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, echo{Value: "pong"})
	})

	router.POST("/payload", func(c *gin.Context) {
		var p payload
		if err := c.BindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		pp.Println(p)
		c.JSON(200, gin.H{"status": "success", "received": p})
	})

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server
}
