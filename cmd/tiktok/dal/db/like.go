package db

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID  uint
	VideoID uint
}
