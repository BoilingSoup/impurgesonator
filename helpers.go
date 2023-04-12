package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func isCaseInsensitiveNameMatch(member, checkAgainst *discordgo.Member) bool {
	usernameMatch := strings.ToLower(member.User.Username) == strings.ToLower(checkAgainst.User.Username)
	nickMatchesUsername := strings.ToLower(member.Nick) == strings.ToLower(checkAgainst.User.Username)

	return usernameMatch || nickMatchesUsername
}
