package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func getCurrentMembers(s *discordgo.Session, event *discordgo.Ready) {
	for _, guilds := range event.Guilds {
		err := s.RequestGuildMembers(guilds.ID, "", 0, "", false)
		if err != nil {
			fmt.Println("Error requesting guild members:", err)
		}
	}
}

func checkCurrentMembers(s *discordgo.Session, event *discordgo.GuildMembersChunk) {
	real, err := s.GuildMember(event.GuildID, doNotImpersonateID)
	if err != nil {
		fmt.Println("Error finding the user who must not be impersonated:", err)
		return
	}

	guild, err := s.Guild(event.GuildID)
	if err != nil {
		fmt.Println("Error finding guild:", err)
		return
	}

	fmt.Printf("Searching server %q for impersonators of %v\n", guild.Name, real.User.String())

	var count int
	for _, member := range event.Members {

		if !isCaseInsensitiveNameMatch(member, real) {
			continue
		}

		if member.User.ID == doNotImpersonateID {
			continue
		}

		err := s.GuildBanCreate(event.GuildID, member.User.ID, 7) // bans user and deletes the last 7 (max) days of comments from this user
		if err != nil {
			fmt.Println("Error banning impersonator:", err)
			continue
		}

		fmt.Println("Banned ", member.User.String())
		count++
	}

	fmt.Printf("Initial search complete. Banned %d users.\n\n", count)
	fmt.Printf("Listening for name changes and new joins...\n")
}

func checkMemberUpdateEvent(s *discordgo.Session, event *discordgo.GuildMemberUpdate) {
	real, err := s.GuildMember(event.GuildID, doNotImpersonateID)
	if err != nil {
		fmt.Println("Error finding the user who must not be impersonated:", err)
		return
	}

	if !isCaseInsensitiveNameMatch(event.Member, real) {
		fmt.Println(event.Member.Nick)
		return
	}

	if event.Member.User.ID == doNotImpersonateID {
		return
	}

	err = s.GuildBanCreate(event.GuildID, event.Member.User.ID, 7) // bans user and deletes the last 7 (max) days of comments from this user
	if err != nil {
		fmt.Println("Error banning impersonator:", err)
		return
	}

	fmt.Println("Banned", event.Member.User.String())
}

func checkNewMemberJoin(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	real, err := s.GuildMember(event.GuildID, doNotImpersonateID)
	if err != nil {
		fmt.Println("Error finding the user who must not be impersonated:", err)
		return
	}

	if !isCaseInsensitiveNameMatch(event.Member, real) {
		return
	}

	if event.Member.User.ID == doNotImpersonateID {
		return
	}

	err = s.GuildBanCreate(event.GuildID, event.Member.User.ID, 7) // bans user and deletes the last 7 (max) days of comments from this user
	if err != nil {
		fmt.Println("Error banning impersonator:", err)
		return
	}

	fmt.Println("Banned", event.Member.User.String())
}
