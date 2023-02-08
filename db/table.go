package db

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID      uint
	User        User
	VideoID     uint
	CommentText string
}
type Like struct {
	gorm.Model
	UserID  uint
	VideoID uint
}
type User struct {
	gorm.Model
	Name     string
	Password string

	Videos []Video
	Likes  []Like
}
type Video struct {
	gorm.Model
	UserID   uint
	PlayUrl  string
	CoverUrl string
	Title    string

	Likes    []Like
	Comments []Comment
}
