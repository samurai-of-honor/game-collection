package main

import tg "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	startMsg       = "👋 Привіт 🕹 Обирай гру:"
	unknownCommand = "Невідома команда 😕"

	theRescue = "\U0001F9BA Порятунок"
	bunker    = "☢️ Бункер"
	mafia     = "🎩 Мафія"
)

var startKeyboard = tg.NewReplyKeyboard(
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(theRescue),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(bunker),
	),
	tg.NewKeyboardButtonRow(
		tg.NewKeyboardButton(mafia),
	),
)
