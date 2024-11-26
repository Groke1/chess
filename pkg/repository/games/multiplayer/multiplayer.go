package multiplayer

import (
	"common"
	"context"
	"github.com/jmoiron/sqlx"
	"pkg/repository/db"
	"pkg/repository/games"
)

type Multiplayer struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Multiplayer {
	return &Multiplayer{
		db: db,
	}
}

func (m *Multiplayer) GetGames(ctx context.Context, username string) (games.GameData, error) {
	query := `SELECT * FROM multi_games_view WHERE username = $1`
	return db.GetSlice[common.MultiplayerData](ctx, m.db, query, username)
}

func (m *Multiplayer) GetGameById(ctx context.Context, gameId int) (games.GameData, error) {
	query := `SELECT * FROM multi_games_view WHERE game_id = $1`
	return db.GetRow[common.MultiplayerData](ctx, m.db, query, gameId)
}

func (m *Multiplayer) GetGamesByResult(ctx context.Context, username string, result string) (games.GameData, error) {
	query := `SELECT * FROM multi_games_view WHERE username = $1 AND result = $2`
	return db.GetSlice[common.MultiplayerData](ctx, m.db, query, username, result)
}

func (m *Multiplayer) GetByTimeControl(ctx context.Context, username string, timeMinutes, increment int) (games.GameData, error) {
	query := `SELECT * FROM multi_games_view WHERE username = $1 AND time_control = CONCAT($2, "+", $3)`
	return db.GetSlice[common.MultiplayerData](ctx, m.db, query, username, timeMinutes, increment)
}

type MultiplayerData struct{}

func (m *Multiplayer) AddGame(ctx context.Context, data MultiplayerData) error {
	panic("not implemented")
}
