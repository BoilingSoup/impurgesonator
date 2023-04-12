package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string
var doNotImpersonateID string

func init() {
	flag.StringVar(&token, "t", "", "Discord Bot Token")
	flag.StringVar(&doNotImpersonateID, "u", "", "Discord ID of the user who must not be impersonated.")
	flag.Parse()
}

func main() {
	if token == "" {
		fmt.Println("Required flags not provided. Please run: impurgesonator -t <bot token> -u <Discord ID>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(getCurrentMembers)
	dg.AddHandler(checkCurrentMembers)
	dg.AddHandler(checkMemberUpdateEvent)
	dg.AddHandler(checkNewMemberJoin)

	// Not sure if this is needed for non-public bots... It was part of the boilerplate so leaving it here.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Printf("Impurgesonator is now running.  Press CTRL-C to exit.\n\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
