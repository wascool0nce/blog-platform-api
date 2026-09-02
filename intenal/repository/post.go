package repository

import (
	"database/sql"
	"fmt"

	"github.com/wascool0nce/blog-platform-api/config"
	"github.com/wascool0nce/blog-platform-api/intenal/model"
)

type DBConnect interface {
	ConnectDB(cfg config.Config) (*sql.DB, error)
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(title, content, category string) (*model.Post, error) {
	const query = `
		INSERT INTO posts (
			title,
			content,
			category
		)
		VALUES ($1, $2, $3)
		RETURNING id, title, content, category, created_at
	`

	var created model.Post

	err := r.db.QueryRow(query, title, content, category).Scan(&created.ID, &created.Title, &created.Content, &created.Category, &created.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}
	return &created, nil
}
