package main

import (
	"flag"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var token = flag.String("t", "", "Token for access bot")
var bot *tg.BotAPI

func Auth() {
	var err error
	flag.Parse()
	bot, err = tg.NewBotAPI(*token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func SendMsg(msg tg.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func main() {
	Auth()
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tg.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case "/start":
			msg.Text = startMsg
			msg.ReplyMarkup = startKeyboard
		default:
			msg.Text = unknownCommand
			msg.ReplyMarkup = startKeyboard
		}

		SendMsg(msg)
	}
}
