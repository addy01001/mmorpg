package main

import (
	"fmt"
	"mmorpg-bot/src/bot"
	"mmorpg-bot/src/config"
)

func main() {
  err := config.ReadConfig()

  if err != nil {
  	fmt.Println(err.Error())
  	return
  }

  bot.Start()

  <-make(chan struct{})
  return
}