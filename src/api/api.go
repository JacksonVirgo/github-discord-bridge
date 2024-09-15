package api

import "github.com/gin-gonic/gin"

type echo struct {
	Value string `json:"value"`
}

func InitApi() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, echo{Value: "pong"})
	})

	router.Run("0.0.0.0:8080")
}
