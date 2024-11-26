package puzzles

import (
	"common"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"pkg/repository/db"
)

var insufficientPermissionError = errors.New("insufficient permissions to access the data")

type Puzzles struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Puzzles {
	return &Puzzles{db: db}
}

type Puzzle struct {
	Id       int
	Position string
	Solution string
	AuthorId int
}

func (p *Puzzles) GetPuzzles(ctx context.Context, username string) ([]*common.PuzzleData, error) {
	query := `SELECT * FROM puzzles_view WHERE username = $1`
	return db.GetSlice[common.PuzzleData](ctx, p.db, query, username)
}

func (p *Puzzles) GetPuzzleById(ctx context.Context, puzzleId int) (*common.PuzzleData, error) {
	query := `SELECT pos.fen AS "position", 
					 puz.solution AS "solution", 
					 u.name AS "author"
			  FROM puzzle puz JOIN position pos ON pos.id = puz.position_id
				   JOIN "user" u ON u.id = puz.author_id
			  WHERE puz.id = $1`
	return db.GetRow[common.PuzzleData](ctx, p.db, query, puzzleId)
}

func (p *Puzzles) GetSolvedPuzzles(ctx context.Context, username string) ([]*common.PuzzleData, error) {
	query := `SELECT * FROM puzzles_view WHERE username = $1 AND is_correct`
	return db.GetSlice[common.PuzzleData](ctx, p.db, query, username)
}

func (p *Puzzles) GetUnsolvedPuzzles(ctx context.Context, username string) ([]*common.PuzzleData, error) {
	query := `SELECT * FROM puzzles_view WHERE user_id = $1 AND NOT is_correct`
	return db.GetSlice[common.PuzzleData](ctx, p.db, query, username)

}

func (p *Puzzles) GetRandomPuzzle(ctx context.Context, username string) (*common.PuzzleData, error) {
	panic("not implemented")
}

func (p *Puzzles) DeletePuzzle(ctx context.Context, username string, puzzleId int) error {
	if err := p.isAdmin(ctx, username); err != nil {
		return err
	}
	query := `DELETE FROM puzzle WHERE id = $1`
	return db.Exec(ctx, p.db, query, puzzleId)
}

func (p *Puzzles) CreatePuzzle(ctx context.Context, username string, puzzle *Puzzle) error {
	if err := p.isAdmin(ctx, username); err != nil {
		return err
	}

	query := `INSERT INTO position (fen) VALUES ($1) ON CONFLICT (fen) 
			  DO UPDATE SET fen = EXCLUDED.fen RETURNING id`
	posId, err := db.GetRow[int](ctx, p.db, query, puzzle.Position)
	if err != nil {
		return err
	}
	query = `INSERT INTO puzzle (position_id, solution, author_id)
			 VALUES ($1, $2, $3)`
	return db.Exec(ctx, p.db, query, posId, puzzle.Solution, puzzle.AuthorId)
}

func (p *Puzzles) UpdatePuzzle(ctx context.Context, username string) error {
	if err := p.isAdmin(ctx, username); err != nil {
		return err
	}
	return nil
}

func (p *Puzzles) AddAttempt(ctx context.Context, username string) {}

func (p *Puzzles) isAdmin(ctx context.Context, username string) error {
	query := `SELECT r.name FROM "user" u JOIN role r 
    		  ON u.role_id = r.id 
    		  WHERE u.name = $1`
	role, err := db.GetRow[string](ctx, p.db, query, username)
	if err != nil {
		return err
	}
	if *role == "user" {
		return insufficientPermissionError
	}
	return nil
}
