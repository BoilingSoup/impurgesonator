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
var modID string

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&modID, "m", "", "Moderator Discord ID")
	flag.Parse()
}

func main() {
	if token == "" {
		fmt.Println("Required flags not provided. Please run: impurgesonator -t <bot token> -m <moderator Discord ID>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(chunk)
	dg.AddHandler(guildAdd)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Impurgesonator is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// scan for usernames that are impersonators
	// guild, err := event.Guild(event.Application.GuildID)
	// guild, err := event.Guilds
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	for _, guilds := range event.Guilds {
		err := s.RequestGuildMembers(guilds.ID, "", 0, "", true)
		if err != nil {
			fmt.Println(err)
		}

		// guild, err := s.State.Guild(guilds.ID)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// fmt.Println(guild.Members)
	}
}

func chunk(s *discordgo.Session, event *discordgo.GuildMembersChunk) {
	for _, member := range event.Members {
		fmt.Println(member.User.Username)
	}
}

func guildAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	fmt.Printf("Nick: %v\n", event.Member.Nick)
	fmt.Printf("UserID: %v\n", event.Member.User.ID)
	fmt.Printf("Username: %v\n", event.Member.User.Username)

	// err := s.RequestGuildMembers(event.GuildID, "", 0, "", true)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	guild, err := s.State.Guild(event.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", guild.Members)

	// st, err := s.GuildMembers(event.GuildID, "", 1000)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%v\n", st)
}
