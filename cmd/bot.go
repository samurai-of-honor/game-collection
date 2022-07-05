package main

import (
	"flag"
	"game-collection/logger"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

var token = flag.String("t", "", "Token for access bot")
var bot *tg.BotAPI

func Auth() {
	var err error
	flag.Parse()
	bot, err = tg.NewBotAPI(*token)
	logger.CheckError(err)

	bot.Debug = true
	logger.AuthSuccess(bot)
}

func main() {
	Auth()
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // ignore any non-Message Updates
			MsgHandler(update.Message)
		} else if update.CallbackQuery != nil {
			// Telling Tg to show the user a message with the data received
			callback := tg.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			_, clbErr := bot.Request(callback)
			logger.CheckError(clbErr)

			CallbackHandler(update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID, update.CallbackQuery.Data)
		}
	}
}
