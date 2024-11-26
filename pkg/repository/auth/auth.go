package auth

import (
	"context"
	"github.com/jmoiron/sqlx"
	"pkg/repository/db"
)

type Auth struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) AddUser(ctx context.Context, username string) {
	query := `INSERT INTO "user" (name, role_id) 
			  VALUES ($1, find_id('role', 'name', 'user')) ON CONFLICT DO NOTHING`
	_ = db.Exec(ctx, a.db, query, username)
}
