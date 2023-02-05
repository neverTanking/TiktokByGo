package db

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserID   uint
	PlayUrl  string
	CoverUrl string
	Title    string

	Likes    []Like
	Comments []Comment
}
