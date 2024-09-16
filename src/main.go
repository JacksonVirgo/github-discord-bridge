package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
	server := api.InitApi()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("API server error: %v\n", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		discord.StartDiscordBot()
	}()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-signalChannel

	fmt.Println("Shutting down...")
	discord.CloseDiscordBot()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("API server forced to shutdown: %v\n", err)
	}
	wg.Wait()

	fmt.Println("Bot shut down successfully.")
}
