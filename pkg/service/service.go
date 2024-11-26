package service

import (
	"common"
	"context"
	"github.com/sirupsen/logrus"
	"pkg/chess-logic/player"
	"pkg/chess-logic/position"
	"pkg/chess-logic/result"
	"pkg/config"
	pf "pkg/handler/parse-functions"
	"pkg/repository"
	"sync"
)

type Service interface {
	CreateSingleGame(game common.GameWithEngine) common.StartGameResponse
	CreateMultiGame()
	MakeMove(ctx context.Context, move common.Move) common.MoveResponse
}
type serviceImpl struct {
	repo  *repository.Repository
	games map[int]*Game
	id    int
	mx    sync.Mutex
}

func NewService(repo *repository.Repository) *serviceImpl {
	return &serviceImpl{
		repo:  repo,
		games: make(map[int]*Game),
		mx:    sync.Mutex{},
	}
}

func (s *serviceImpl) parseGameConfig(startPos position.Position, game common.GameWithEngine) config.SingleGameConfig {
	engineName := "stockfish"
	if game.EngineName != "" {
		engineName = game.EngineName
	}
	minutes, increment := 0, 0
	if game.TimeControl != "" {
		minutes, increment, _ = pf.ParseTimeControl(game.TimeControl)
	}
	return config.SingleGameConfig{
		Username:             game.Username,
		EngineName:           engineName,
		EngineLevel:          game.Level,
		PositionFen:          startPos.GetFEN(),
		IsPlayerWhite:        game.Color == "white" || game.Color == "",
		TimeControlMinutes:   minutes,
		TimeControlIncrement: increment,
	}
}

func (s *serviceImpl) CreateSingleGame(game common.GameWithEngine) common.StartGameResponse {
	var startPos position.Position
	var resp common.StartGameResponse
	if game.Fen == "" {
		startPos = position.NewStartPosition()
	} else {
		var err error
		startPos, err = position.NewFromFEN(game.Fen)
		if err != nil {
			logrus.Error(err)
			return resp
		}
	}

	cfg := s.parseGameConfig(startPos, game)
	pl1 := player.HumanPlayer{}
	pl2 := player.NewStockfishPlayer(game.Level)
	newGame := NewGame(cfg, startPos, &pl1, pl2)
	s.mx.Lock()
	s.id++
	s.games[s.id] = newGame
	resp.GameId = s.id

	s.mx.Unlock()

	return resp
}

func (s *serviceImpl) CreateMultiGame() {}

func (s *serviceImpl) MakeMove(ctx context.Context, move common.Move) common.MoveResponse {
	g := s.games[move.GameId]
	var resp common.MoveResponse
	res, err := g.Move(move.Move)
	if err != nil {
		logrus.Error(err)
		return resp
	}
	resp.Fen = g.pos.GetFEN()
	resp.IsValidMove = true
	resp.Result = res.String()
	if res == result.Draw {
		g.cfg.Result = "draw"
	} else if res == result.Win && g.player1 == g.nowPlayer {
		g.cfg.Result = "white win"
	} else {
		g.cfg.Result = "black win"
	}
	if res != result.Unknown {
		_ = s.repo.Singleplayer.AddGame(ctx, g.cfg)
	}
	return resp
}
