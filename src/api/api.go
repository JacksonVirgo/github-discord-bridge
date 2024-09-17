package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/JacksonVirgo/github-discord-bridge/src/github"
	"github.com/gin-gonic/gin"
)

type echo struct {
	Value string `json:"value"`
}

type payload struct {
	Action string `json:"action"`
}

func verifyGithubSignature(body []byte, signature string) bool {
	secret := github.GetWebhookSecret()
	if secret == "" {
		fmt.Println("Webhook secret not found")
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	eq := hmac.Equal([]byte(signature), []byte(expectedMAC))
	fmt.Printf("Signature verification: %t\n", eq)
	return eq
}

func InitApi() *http.Server {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, echo{Value: "pong"})
	})

	router.POST("/payload", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read body"})
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewReader(body))

		signature := c.GetHeader("X-Hub-Signature-256")
		if signature == "" {
			c.JSON(400, gin.H{"error": "Missing signature header"})
			return
		}

		if !verifyGithubSignature(body, signature) {
			c.JSON(401, gin.H{"error": "Invalid signature"})
			return
		}

		var jsonBody map[string]interface{}
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		jsonString, err := json.Marshal(jsonBody)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to convert to JSON string"})
			return
		}

		c.JSON(200, gin.H{"status": "success", "received": string(jsonString)})
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
