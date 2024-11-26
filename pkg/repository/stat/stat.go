package stat

import (
	"common"
	"context"
	"github.com/jmoiron/sqlx"
	"pkg/repository/db"
)

type Stat struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Stat {
	return &Stat{db: db}
}

func (s *Stat) GetUserStat(ctx context.Context, username string) (*common.UserStat, error) {
	query := `SELECT rating, top_place, amount_games, amount_wins, amount_loses, amount_draws 
			  FROM stat_view WHERE username = $1`
	return db.GetRow[common.UserStat](ctx, s.db, query, username)
}

func (s *Stat) GetPuzzleStat(ctx context.Context, puzzleId int) (*common.PuzzleStat, error) {
	query := `SELECT * FROM puzzle_stat_view WHERE puz_id = $1`
	return db.GetRow[common.PuzzleStat](ctx, s.db, query, puzzleId)
}

func (s *Stat) GetTopStat(ctx context.Context, amount int) ([]*common.TopStat, error) {
	query := `SELECT username, rating, top_place 
			  FROM top_view LIMIT $1`
	return db.GetSlice[common.TopStat](ctx, s.db, query, amount)
}

func (s *Stat) GetTimeControlsStat(ctx context.Context, username string) ([]*common.TimeControlsStat, error) {
	query := `SELECT username, time_control, amount_games, amount_wins, amount_loses, amount_draws 
			  FROM time_controls_stat_view WHERE username = $1`
	return db.GetSlice[common.TimeControlsStat](ctx, s.db, query, username)
}
