package main

import (
	"TarkovBot/bot"
	"TarkovBot/commands"
	"TarkovBot/config"
	"log"
)

func init() {
	// load json data first
	if err := config.LoadFile(); err != nil {
		log.Fatal(err.Error())
	}
	// Load Initial bot data
	if err := config.LoadSettings(); err != nil {
		log.Fatal(err.Error())
	}
	go commands.Load()
}

func main() {
	// start bot
	bot.NewBot()
}
