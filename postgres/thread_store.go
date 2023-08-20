package postgres

import (
	"fmt"

	"github.com/ctheil/goreddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ThreadStore struct {
	*sqlx.DB
}

// imp CRUD methods

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	// query single thread from db based on ID
	var t goreddit.Thread
	/*
		FIRST_PARAM: reference to thread that GET will populate with the db return
		SECOND_PARAM: SQL query statement
		THIRD_PARAM: populate query to avoid SQL INJECTION
	*/
	// wrap call in error handler
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		// return empty thread and formatted error statement
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err) // %w special placeholder for errors
	}
	return t, nil // not sure what return is t &! *t

}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var tt []goreddit.Thread
	if err := s.Select(&tt, `SELECT * FROM threads`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("Error getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`, t.ID, t.Title, t.Description); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}
func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `
	UPDATE threads SET 
	title = $1, 
	description = $2 
	WHERE id  = $3 
	RETURNING *`,
		t.Title,
		t.Description,
		t.ID); err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}
func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
