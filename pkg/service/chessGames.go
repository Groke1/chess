package service

import (
	"pkg/chess-logic/player"
	"pkg/chess-logic/position"
	"pkg/chess-logic/result"
	"pkg/config"
)

type Game struct {
	cfg       config.SingleGameConfig
	pos       position.Position
	player1   player.Player
	player2   player.Player
	nowPlayer player.Player
}

func NewGame(cfg config.SingleGameConfig, pos position.Position, player1, player2 player.Player) *Game {
	return &Game{
		cfg:       cfg,
		pos:       pos,
		player1:   player1,
		player2:   player2,
		nowPlayer: player1,
	}
}

func (g *Game) Move(move string) (result.Result, error) {
	res, err := g.pos.MakeMoveByString(move)
	if err != nil {
		return res, err
	}
	g.changePlayer()
	switch typePlayer := g.nowPlayer.(type) {
	case *player.StockfishPlayer:
		res, _ := g.pos.MakeMoveByString(typePlayer.Move(g.pos))
		g.changePlayer()
		return res, nil
	default:
		return res, nil
	}
}

func (g *Game) changePlayer() {
	if g.nowPlayer == g.player1 {
		g.nowPlayer = g.player2
		return
	}
	g.nowPlayer = g.player1
}
