package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	discordToken := os.Getenv("DISCORD_TOKEN")
	fmt.Printf("Discord token: %s\n", discordToken)
}