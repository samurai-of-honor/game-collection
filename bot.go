package main

import (
	"flag"
	"game-collection/games"
	"game-collection/logger"
	. "game-collection/settings"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

var token = flag.String("t", "", "Token for access bot")

func Auth() *tg.BotAPI {
	flag.Parse()
	bot, err := tg.NewBotAPI(*token)
	logger.CheckError(err)

	bot.Debug = true
	logger.AuthSuccess(bot)

	return bot
}

func SendMsg(bot *tg.BotAPI, msg tg.Chattable) {
	_, sendErr := bot.Send(msg)
	logger.CheckError(sendErr)
}

func UpdateHandler(bot *tg.BotAPI, updates tg.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tg.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "HTML"

		switch update.Message.Text {
		case "/start":
			msg.Text = StartMsg
			msg.ReplyMarkup = StartKeyboard
		case Bunker:
			msg.Text = BunkerOptions
			msg.ReplyMarkup = BunkerKeyboard
		case BunkerCharacter:
			msg.Text = games.BunkerCharacter()
			msg.ReplyMarkup = BunkerKeyboard
		case BunkerShelter:
			msg.Text = games.BunkerInfo()
			msg.ReplyMarkup = BunkerKeyboard
		default:
			msg.Text = UnknownCommand
			msg.ReplyMarkup = StartKeyboard
		}
		SendMsg(bot, msg)
	}
}

func main() {
	bot := Auth()
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	UpdateHandler(bot, updates)
}
