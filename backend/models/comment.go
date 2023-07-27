package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Other model properties and methods related to comment data
