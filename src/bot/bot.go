package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gomodule/redigo/redis"
	"mmorpg-bot/src/config"
	"mmorpg-bot/src/helpers"
	"mmorpg-bot/src/services"
)

var BotId string
var pool = helpers.NewPool()

// var goBot *discordgo.Session

func Start() {
	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Making our bot a user using User function .
	u, err := goBot.User("@me")

	//Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Storing our id from u to BotId .
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(reactionAddListener)

	err = goBot.Open()

	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "bot help" {
		msg, _ := s.ChannelMessageSend(m.ChannelID, "You have pressed help\n\nCommands:\ninfo - Get your current details\nshop - Open shop\nswitch - switch to Co-op/Single-player\nstory - Advance to the next story point\ndungeon - Starts duel with npc ranked similar to you\ndaily - Collect your daily reward\ni - Open your inventory")
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "⏩")
	}

	if m.Content == "bot info" {
		res, err := services.StartUser(m.Author.ID, m.Author.Username)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed")
		}
		sendText := fmt.Sprintf("Name: %s\nLevel: %d\nXP: %d", res.Data.Name, res.Data.Level, res.Data.XP)
		s.ChannelMessageSend(m.ChannelID, sendText)
	}

	if m.Content == "bot story" {
		res, err := services.GetStory(m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed")
		}
		// sendText := fmt.Sprintf("The beginning\n\n%s", res.Data.DisplayText)
		msg, _ := s.ChannelMessageSend(m.ChannelID, res.Data.DisplayText)
		s.MessageReactionAdd(msg.ChannelID, msg.ID, "⏩")
		client := pool.Get()
		_, err = client.Do("SET", msg.ID, m.Author.ID, "EX", "15")
		if res.Data.StoryType == "chest" {
			s.MessageReactionAdd(msg.ChannelID, m.ID, "🎁")
		}
		if err != nil {
			panic(err)
		}
	}
}

func reactionAddListener(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	client := pool.Get()
	if r.UserID == BotId {
		return
	}

	if r.Emoji.Name == "🎁" {
		msg, err := redis.String(client.Do("GET", r.MessageID))
		if err != nil {
			return
		}
		res, _ := services.OpenChest(r.UserID, msg)
		responseString := fmt.Sprintf("XP: +%d\nCoins: +%d", res.Data.XP, res.Data.Coins)
		for _, item := range res.Data.Contents { //discordgo package from the repo of b //discordgo package  //discordgo package from the repo of bwmarrin .from the repo of bwmarrin .wmarrin .
			responseString = fmt.Sprintf("%s\n%s", responseString, item.Name)
		}
		m, _ := s.ChannelMessageSend(r.ChannelID, responseString)
		s.MessageReactionAdd(m.ChannelID, m.ID, "⏩")
	}

	if r.Emoji.Name == "⏩" {
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
}
