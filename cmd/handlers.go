package main

import (
	. "game-collection/configs"
	"game-collection/games"
	"game-collection/logger"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMsg(msg tg.Chattable) {
	_, sendErr := bot.Send(msg)
	logger.CheckError(sendErr)
}

// MsgHandler recognize and send message
func MsgHandler(Message *tg.Message) {
	msg := tg.NewMessage(Message.Chat.ID, "")
	msg.ParseMode = "HTML"

	switch Message.Text {
	case "/start":
		msg.Text = StartMsg
		msg.ReplyMarkup = StartKeyboard
	case Return:
		msg.Text = SelectGame
		msg.ReplyMarkup = StartKeyboard
	case Bunker:
		msg.Text = BunkerOptions
		msg.ReplyMarkup = BunkerKeyboard
	case BunkerCharacter:
		msg.Text = games.BunkerCharacter()
		msg.ReplyMarkup = CharacterKeyboard
	case BunkerShelter:
		msg.Text = games.BunkerInfo()
		msg.ReplyMarkup = BunkerKeyboard
	case BunkerCatastrophe:
		msg.Text = games.BunkerCatastrophe()
		msg.ReplyMarkup = BunkerKeyboard
	default:
		msg.Text = UnknownCommand
		msg.ReplyMarkup = StartKeyboard
	}
	SendMsg(msg)
}

// CallbackHandler recognize requests and send message
func CallbackHandler(ChatId int64, MessageId int, CallbackData string) {
	var text string
	switch CallbackData {
	case Regenerated:
		text = games.BunkerCharacter()
	}

	// Send a changed message
	msg := tg.NewEditMessageTextAndMarkup(ChatId, MessageId, text, CharacterKeyboard)
	msg.ParseMode = "HTML"
	SendMsg(msg)
}
