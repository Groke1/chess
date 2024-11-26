package repository

import (
	"github.com/jmoiron/sqlx"
	"pkg/repository/auth"
	"pkg/repository/games/multiplayer"
	"pkg/repository/games/singleplayer"
	"pkg/repository/puzzles"
	"pkg/repository/stat"
)

type Repository struct {
	*puzzles.Puzzles
	*multiplayer.Multiplayer
	*singleplayer.Singleplayer
	*stat.Stat
	*auth.Auth
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Puzzles:      puzzles.New(db),
		Multiplayer:  multiplayer.New(db),
		Singleplayer: singleplayer.New(db),
		Stat:         stat.New(db),
		Auth:         auth.New(db),
	}
}
