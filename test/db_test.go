package test

import (
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"testing"
)

func TestDBInit(t *testing.T) {
	db.Init()
}

func TestUniqueUsername(t *testing.T) {
	//This test will fail
	db.Init()
	user1 := db.User{
		Model:           gorm.Model{},
		Name:            "duplicateUsername",
		Password:        "111",
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		Videos:          nil,
		Likes:           nil,
	}
	user2 := db.User{
		Model:           gorm.Model{},
		Name:            "duplicateUsername",
		Password:        "111",
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		Videos:          nil,
		Likes:           nil,
	}
	db.DB.Create(&user1)
	db.DB.Create(&user2)
}
