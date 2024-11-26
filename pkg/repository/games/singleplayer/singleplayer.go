package singleplayer

import (
	"common"
	"context"
	"github.com/jmoiron/sqlx"
	"pkg/config"
	"pkg/repository/db"
	"pkg/repository/games"
)

type Singleplayer struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Singleplayer {
	return &Singleplayer{db: db}
}

func (s *Singleplayer) GetGames(ctx context.Context, username string) (games.GameData, error) {
	query := `SELECT * FROM single_games_view WHERE username = $1`
	return db.GetSlice[common.SingleplayerData](ctx, s.db, query, username)
}

func (s *Singleplayer) GetGameById(ctx context.Context, gameId int) (games.GameData, error) {
	query := `SELECT * FROM single_games_view WHERE game_id = $1`
	return db.GetRow[common.SingleplayerData](ctx, s.db, query, gameId)
}

func (s *Singleplayer) GetGamesByResult(ctx context.Context, username string, result string) (games.GameData, error) {
	query := `SELECT * FROM single_games_view WHERE username = $1 AND result = $2`
	return db.GetSlice[common.SingleplayerData](ctx, s.db, query, username, result)
}

func (s *Singleplayer) GetByTimeControl(ctx context.Context, username string, timeMinutes, increment int) (games.GameData, error) {
	query := `SELECT * FROM single_games_view WHERE username = $1 AND time_control = CONCAT($1, "+", $2)`
	return db.GetSlice[common.SingleplayerData](ctx, s.db, query, username, timeMinutes, increment)
}

func (s *Singleplayer) AddGame(ctx context.Context, cfg config.SingleGameConfig) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query := `INSERT INTO time_control (time_minutes, increment) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err = tx.ExecContext(ctx, query, cfg.TimeControlMinutes, cfg.TimeControlIncrement)
	if err != nil {
		return err
	}
	query = `INSERT INTO position (fen) VALUES ($1) ON CONFLICT DO NOTHING`
	_, err = tx.ExecContext(ctx, query, cfg.PositionFen)
	if err != nil {
		return err
	}

	query = `INSERT INTO singleplayer (player_id, engine_id, engine_level, 
                          				is_player_white, date, result_id, position_id, 
                          				time_control_id, moves) 
			  VALUES (
			          find_id('user', 'name' $1), 
			          find_id('engine', 'name', $2), 
			          $3, $4, NOW(), 
			          find_id('result', 'name', $5), 
			          find_id('position', 'fen', $6), 
			          find_time_control_id($7, $8), $9
			  )`

	_, err = tx.ExecContext(ctx, query, cfg.Username, cfg.EngineName, cfg.EngineLevel,
		cfg.IsPlayerWhite, cfg.Result, cfg.PositionFen, cfg.TimeControlMinutes,
		cfg.TimeControlIncrement, cfg.Moves)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}
