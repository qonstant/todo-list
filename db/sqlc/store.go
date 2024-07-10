package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("panic recovery: rollback error: %v\n", rbErr)
			}
			panic(p) // re-panic after rollback
		} else if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("error: rollback error: %v\n", rbErr)
			}
		} else {
			err = tx.Commit()
			if err != nil {
				fmt.Printf("error: commit error: %v\n", err)
			}
		}
	}()

	q := New(tx)
	err = fn(q)
	if err != nil {
		return fmt.Errorf("error executing transaction function: %w", err)
	}
	return nil
}
