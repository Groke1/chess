package tgcommands

import (
	"bytes"
	"clients/commands"
	"common"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"strings"
)

type TelegramCmd struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *TelegramCmd {
	return &TelegramCmd{
		bot: bot,
	}
}

func HttpHelper[K, R any](url string, value K) (R, error) {
	jsonData, err := json.Marshal(value)
	var res R
	if err != nil {
		return res, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

func (t *TelegramCmd) MakeMove(update *tgbotapi.Update, gameId int, move string) {
	newMove := common.Move{GameId: gameId, Move: move}
	moveResp, err := HttpHelper[common.Move, common.MoveResponse]("http://server:8080/api/move", newMove)

	if err != nil {
		log.Println(err)
		return
	}

	url := "https://backscattering.de/web-boardimage/board.svg?fen=" + strings.Fields(moveResp.Fen)[0]

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, url)
	_, err = t.bot.Send(msg)
	if err != nil {
		log.Println("Error send msg")
	}

}

func (t *TelegramCmd) StartSingleGame(update *tgbotapi.Update, level int) {
	newGame := common.GameWithEngine{Username: update.Message.From.UserName, Level: level}
	startGameResp, err := HttpHelper[common.GameWithEngine, common.StartGameResponse]("http://server:8080/api/single", newGame)
	if err != nil {
		log.Println(err)
		return
	}
	commands.GameSessions[update.Message.Chat.ID] = startGameResp.GameId
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Game with stockfish %d", level))
	t.bot.Send(msg)
}

func (t *TelegramCmd) Start(update *tgbotapi.Update) {
	url := "http://server:8080/auth"
	newUser := common.User{Username: update.Message.From.UserName}
	jsonData, _ := json.Marshal(newUser)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
}

func (t *TelegramCmd) Help(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Help")
	t.bot.Send(msg)
}

func (t *TelegramCmd) GetTop(update *tgbotapi.Update) {
	resp, err := http.Get("http://server:8080/api/stat/top/10")
	if err != nil {
		log.Println(err)
		return
	}
	var topStat []common.TopStat
	if err := json.NewDecoder(resp.Body).Decode(&topStat); err != nil {
		log.Println(err)
		return
	}
	for _, stat := range topStat {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s %d %d", stat.Username, stat.Rating, stat.TopPlace))
		t.bot.Send(msg)
	}
}

func (t *TelegramCmd) GetStat(update *tgbotapi.Update) {
	resp, err := http.Get("http://server:8080/api/stat/" + update.Message.From.UserName)
	if err != nil {
		log.Println(err)
		return
	}
	var stat common.UserStat
	if err := json.NewDecoder(resp.Body).Decode(&stat); err != nil {
		log.Println(err)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("%d %d %d %d %d", stat.Rating, stat.TopPlace, stat.AmountWins,
			stat.AmountLoses, stat.AmountDraws))
	t.bot.Send(msg)
}
