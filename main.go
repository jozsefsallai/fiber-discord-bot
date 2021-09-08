package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jozsefsallai/fiber-discord-bot/config"
	"github.com/jozsefsallai/fiber-discord-bot/handlers"
)

func init() {
	config.Load()
}

func main() {
	client, err := discordgo.New(fmt.Sprintf("Bot %s", config.Config.Bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	client.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers

	client.AddHandler(handlers.HandleMemberJoin)
	client.AddHandler(handlers.HandleMessageReaction)

	err = client.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Fiber Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	client.Close()
}
