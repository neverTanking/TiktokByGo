package db

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID      uint
	User        User
	VideoID     uint
	CommentText string
}
