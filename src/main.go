package main

import (
	"log"
	"os"

	"github.com/JacksonVirgo/github-discord-bridge/src/controllers"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	err = controllers.LoadGithubContext()
	if err != nil {
		log.Fatal(err)
	}

	discordToken := os.Getenv("DISCORD_TOKEN")
	controllers.StartDiscordBot(discordToken)
}