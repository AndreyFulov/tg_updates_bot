package main

import (
	"flag"
	"fmt"

	"github.com/AndreyFulov/tg_updates_bot"
)

var tgbotkey *string
var tgchannel *int64

func main() {
	tgbotkey = flag.String("token", "", "Telegram bot token")
	tgchannel = flag.Int64("ch", 0, "Telegram channel id")

	flag.Parse()
	responseChan := make(chan tg_updates_bot.Response)

	go tg_updates_bot.Bot(tgbotkey, tgchannel, responseChan)
	for {
		response := <-responseChan
		fmt.Printf("Text: %s \t PhotoURL: %s\n", response.Text, response.PhotoURL)
	}
}
