package goreddit

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error // modify the actual thread
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadId uuid.UUID) ([]Post, error)
	CreatePost(t *Post) error // modify the actual thread
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postId uuid.UUID) ([]Comment, error)
	CreateComment(t *Comment) error // modify the actual thread
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
