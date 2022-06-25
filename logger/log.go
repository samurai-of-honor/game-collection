package logger

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"time"
)

var (
	file, _ = os.OpenFile("game-collection/logger/game-collection.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile = log.New(file, time.Now().Format(time.RFC3339)+" ", 0)
)

func CheckError(err error) {
	if err != nil {
		LogFile.Fatalln(err)
	}
}

func AuthSuccess(bot *tgbotapi.BotAPI) {
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
