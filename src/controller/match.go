package controller

import (
	"encoding/json"
	"fmt"
	"mmorpg-bot/src/services"

	"github.com/bwmarrin/discordgo"
	"github.com/gomodule/redigo/redis"
)

func RandomMatch(s *discordgo.Session, m *discordgo.MessageCreate) {
	res, err := services.RandomMatch(m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed")
	}
	sendText := "You're matched against\n\n"
	for _, item := range res.Data.TeamB {
		sendText = fmt.Sprintf("%sname:%s\nHealth:%d/%d", sendText, item.Name, item.Hp, item.MaxHp)
	}
	msg, _ := s.ChannelMessageSend(m.ChannelID, sendText)
	client := pool.Get()

	u, _ := json.Marshal(res)

	client.Do("SET", msg.ID, string(u), "EX", "15")
	s.MessageReactionAdd(msg.ChannelID, msg.ID, "👊")
	s.MessageReactionAdd(msg.ChannelID, msg.ID, "ℹ️")
}


func HandleAttack(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	client := pool.Get()
	msg, err := redis.String(client.Do("GET", r.MessageID))
	if err != nil {
		return
	}
	var match services.Match
	json.Unmarshal([]byte(msg), &match)
	fmt.Print(match)
}

func GetInfo(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	client := pool.Get()
	msg, err := redis.String(client.Do("GET", r.MessageID))
	if err != nil {
		return
	}
	var match services.Match
	json.Unmarshal([]byte(msg), &match)
	fmt.Print(match)
}
