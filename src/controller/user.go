package controller

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"mmorpg-bot/src/services"
)

func UserDetails(s *discordgo.Session, m *discordgo.MessageCreate) {
	res, err := services.StartUser(m.Author.ID, m.Author.Username)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed")
	}
	sendText := fmt.Sprintf("Name: %s\nLevel: %d\nXP: %d", res.Data.Name, res.Data.Level, res.Data.XP)
	s.ChannelMessageSend(m.ChannelID, sendText)
}
