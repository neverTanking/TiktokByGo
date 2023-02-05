package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string

	Videos []Video
	Likes  []Like
}
