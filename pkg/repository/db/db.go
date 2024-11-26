package db

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var invalidQuery = errors.New("invalid query")

func GetSlice[T any](ctx context.Context, db *sqlx.DB, query string, args ...any) ([]*T, error) {
	slice := make([]*T, 0)
	if err := db.SelectContext(ctx, &slice, query, args...); err != nil {
		logrus.Error(invalidQuery)
		return nil, err
	}
	return slice, nil
}

func GetRow[T any](ctx context.Context, db *sqlx.DB, query string, args ...any) (*T, error) {
	var row T
	if err := db.GetContext(ctx, &row, query, args...); err != nil {
		logrus.Error(invalidQuery)
		return nil, err
	}
	return &row, nil
}

func Exec(ctx context.Context, db *sqlx.DB, query string, args ...any) error {
	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		logrus.Error(invalidQuery)
		return err
	}
	return nil
}
