package controller

import (
	"mmorpg-bot/src/helpers"
	"mmorpg-bot/src/services"

	"github.com/bwmarrin/discordgo"
	"github.com/gomodule/redigo/redis"
)

var pool = helpers.NewPool()

func GetStory(s *discordgo.Session, m *discordgo.MessageCreate) {
	res, err := services.GetStory(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed")
	}
	msg, _ := s.ChannelMessageSend(m.ChannelID, res.Data.DisplayText)
	s.MessageReactionAdd(msg.ChannelID, msg.ID, "⏩")
	client := pool.Get()
	if _, err = client.Do("SET", msg.ID, m.Author.ID, "EX", "15"); err != nil {
		panic(err)
	}
	if res.Data.StoryType == "chest" {
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "🎁")
	}
	if err != nil {
		panic(err)
	}
}

func AdvanceStory(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	client := pool.Get()
	msg, err := redis.String(client.Do("GET", r.MessageID))
	if err != nil {
		return
	}
	if r.UserID != msg {
		return
	}
	res, _ := services.AdvanceStory(r.UserID)
	m, _ := s.ChannelMessageSend(r.ChannelID, res.Data.DisplayText)
	client.Do("SET", m.ID, res.Data.Chest, "EX", 15)
	if res.Data.StoryType == "chest" {
		s.MessageReactionAdd(r.ChannelID, m.ID, "🎁")
	}
}
