package telegram

import (
	"clients/commands"
	"clients/telegram/tgcommands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"strconv"
	"strings"
)

type Telegram struct {
	bot *tgbotapi.BotAPI
}

func New() (*Telegram, error) {
	botToken := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &Telegram{bot: bot}, nil
}

func (t *Telegram) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := t.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	cmds := tgcommands.New(t.bot)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		parts := strings.Fields(update.Message.Text)
		chatID := update.Message.Chat.ID

		if gameId, ok := commands.GameSessions[chatID]; ok {
			if parts[0] == "/move" {
				cmds.MakeMove(&update, gameId, parts[1])
			} else if parts[0] == "/stop" {
				delete(commands.GameSessions, chatID)
			}
		} else {
			switch parts[0] {
			case "/start":
				cmds.Start(&update)
			case "/help":
				cmds.Help(&update)
			case "/withEngine":
				level, err := strconv.Atoi(parts[1])
				if err != nil {
					continue
				}
				cmds.StartSingleGame(&update, level)
			case "/stat":
				cmds.GetStat(&update)
			case "/top":
				cmds.GetTop(&update)
			}

		}
	}
	return nil
}
