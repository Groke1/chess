package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var GameSessions = make(map[int64]int)

type Commands interface {
	Start(update *tgbotapi.Update)
	StartSingleGame(update *tgbotapi.Update, level int)
	MakeMove(update *tgbotapi.Update, gameId int, move string)
	GetStat(update *tgbotapi.Update)
	GetTop(update *tgbotapi.Update)
	Help(update *tgbotapi.Update)
}
