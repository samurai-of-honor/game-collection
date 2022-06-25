package settings

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot keyboards

var StartKeyboard = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(TheRescue),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Bunker),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(Mafia),
	),
)
