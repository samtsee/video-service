package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostID      uint64     `json:"post_id"` // maybe make this uuid as well
	UserID      uuid.UUID  `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Location    Location   `json:"location"`
	VideoID     string     `json:"video_id"` // whatever is generated in AWS S3
	CreatedAt   *time.Time `json:"created_at"`
	EditedAt    *time.Time `json:"edited_at"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
