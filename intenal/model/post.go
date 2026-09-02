package model

import (
	"errors"
	"time"
)

var (
	ErrNotFoundPost = errors.New("post not found")
)

type Post struct {
	ID        int64
	Title     string
	Content   string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
