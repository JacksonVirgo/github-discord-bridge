package main

import (
	"log"

	"github.com/JacksonVirgo/github-discord-bridge/src/discord"
	"github.com/JacksonVirgo/github-discord-bridge/src/github"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = github.LoadGithubContext()
	if err != nil {
		log.Fatal(err)
	}

	err = discord.LoadDiscordContext()
	if err != nil {
		log.Fatal(err)
	}

	discord.StartDiscordBot()
}
