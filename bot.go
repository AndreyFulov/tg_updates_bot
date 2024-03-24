package tg_updates_bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Response struct {
	PhotoURL string
	Text     string
}

// For work we need a bot token and channel, where we send struct data, and if u want, u can set a TG channel id
func Bot(tgbotkey *string, tgchannel *int64, ch chan Response) {
	/*
		tgbotkey = flag.String("token", "", "Telegram bot token")
		tgchannel = flag.Int64("ch", 0, "Telegram channel id")
	*/

	//Check that tgbot is set up
	if *tgbotkey == "" {
		log.Fatal("Bot token must be set")
	}
	//Creating new bot instance
	bot, err := tgbotapi.NewBotAPI(*tgbotkey)
	if err != nil {
		log.Fatal("Something went wrong!")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		//Check that channel id is currect or setted to 0(default value)
		if update.ChannelPost.Chat.ID == *tgchannel || *tgchannel == 0 {
			//Any updates from channel
			if update.ChannelPost != nil {
				//Creating instance of struct that we need for response to channel
				response := &Response{}
				response.Text = update.ChannelPost.Text
				if len(update.ChannelPost.Photo) != 0 {
					file, err := bot.GetFile(tgbotapi.FileConfig{FileID: update.ChannelPost.Photo[len(update.ChannelPost.Photo)-1].FileID})
					if err != nil {
						log.Fatal("Error! ")
					}
					response.PhotoURL = file.Link(*tgbotkey)
					response.Text = update.ChannelPost.Caption
				}
				ch <- *response
			}
		}
	}
}
