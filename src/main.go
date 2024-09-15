package main

import (
	"log"
	"sync"

	"github.com/JacksonVirgo/github-discord-bridge/src/api"
	"github.com/JacksonVirgo/github-discord-bridge/src/discord"
	"github.com/JacksonVirgo/github-discord-bridge/src/github"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load()

	err := github.LoadGithubContext()
	if err != nil {
		log.Fatal(err)
	}

	err = discord.LoadDiscordContext()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		api.InitApi()
		wg.Done()
	}()

	discord.StartDiscordBot()

	wg.Wait()
}
