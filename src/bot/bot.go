package bot

import (
	"fmt"
	"mmorpg-bot/src/config"
	"mmorpg-bot/src/controller"
	"github.com/bwmarrin/discordgo"
)

var BotId string

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
		controller.UserDetails(s, m)
	}

	if m.Content == "bot match" {
		controller.RandomMatch(s, m)
	}

	if m.Content == "bot story" {
		controller.UserDetails(s, m)
	}
}

func reactionAddListener(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == BotId {
		return
	}

	if r.Emoji.Name == "👊" {
		controller.HandleAttack(s, r)
	}

	if r.Emoji.Name == "ℹ️" {
		controller.GetInfo(s, r)
	}

	if r.Emoji.Name == "🎁" {
		controller.OpenChest(s, r)
	}

	if r.Emoji.Name == "⏩" {
		controller.AdvanceStory(s, r)
	}
}
