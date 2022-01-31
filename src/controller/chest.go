package controller

import (
	"fmt"
	"mmorpg-bot/src/helpers"
	"mmorpg-bot/src/services"

	"github.com/bwmarrin/discordgo"
	"github.com/gomodule/redigo/redis"
)

func OpenChest(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	client := helpers.NewPool().Get()
	msg, err := redis.String(client.Do("GET", r.MessageID))
	if err != nil {
		return
	}
	res, _ := services.OpenChest(r.UserID, msg, s.State.User.Token)
	responseString := fmt.Sprintf("XP: +%d\nCoins: +%d", res.Data.XP, res.Data.Coins)
	for _, item := range res.Data.Contents { //discordgo package from the repo of b //discordgo package  //discordgo package from the repo of bwmarrin .from the repo of bwmarrin .wmarrin .
		responseString = fmt.Sprintf("%s\n%s", responseString, item.Name)
	}
	m, _ := s.ChannelMessageSend(r.ChannelID, responseString)
	s.MessageReactionAdd(m.ChannelID, m.ID, "⏩")
}
