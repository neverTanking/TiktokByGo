package db

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	MyModel
	UserID      int64
	User        User
	VideoID     int64
	CommentText string
}
type Like struct {
	MyModel
	UserID  int64
	VideoID int64
}
type User struct {
	MyModel
	Name     string
	Password string

	Videos []Video
	Likes  []Like
}
type Video struct {
	MyModel
	UserID   int64
	PlayUrl  string
	CoverUrl string
	Title    string

	Likes    []Like
	Comments []Comment
}
type MyModel struct {
	ID        int64          `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
