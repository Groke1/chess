package games

import (
	"context"
)

type GameData interface {
}

type TypeGames interface {
	GetGames(ctx context.Context, username string) (GameData, error)
	GetGameById(ctx context.Context, gameId int) (GameData, error)
	GetGamesByResult(ctx context.Context, username string, result string) (GameData, error)
	GetByTimeControl(ctx context.Context, username string, timeMinutes, increment int) (GameData, error)
}
