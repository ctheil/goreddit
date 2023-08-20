package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening the database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}
	return &Store{
		ThreadStore:  &ThreadStore{DB: db},
		PostStore:    &PostStore{DB: db},
		CommentStore: &CommentStore{DB: db},
	}, nil
}

// store meant to have concrete implementations for stores-> remote goreddit. and replace with pointer
type Store struct {
	*ThreadStore
	*CommentStore
	*PostStore
}
