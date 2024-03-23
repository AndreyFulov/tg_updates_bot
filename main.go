package main

import (
	"flag"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgbotkey *string

func main() {
	tgbotkey = flag.String("token", "", "Telegram bot token")
	flag.Parse()

	if *tgbotkey == "" {
		log.Fatal("Bot token must be set")
	}

	bot, err := tgbotapi.NewBotAPI(*tgbotkey)
	if err != nil {
		log.Fatal("Something went wrong!")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.ChannelPost != nil {
			fmt.Println(update.ChannelPost.Photo[0].FileID)
			file, err := bot.GetFile(tgbotapi.FileConfig{FileID: update.ChannelPost.Photo[0].FileID})
			if err != nil {
				log.Fatal("Error! ")
			}
			fmt.Println(file.Link(*tgbotkey))
		}
	}
}
