package test

import (
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"testing"
)

func TestUniqueUsername(t *testing.T) {
	//This test will fail
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

// Fail when DB have duplicate username
// Already set username as UNIQUE key, This test will always pass
func TestIfDBUsernameUnique(t *testing.T) {
	var userList []db.User
	res1 := db.DB.Distinct("name", "age").Find(&userList)
	res2 := db.DB.Find(&userList)
	if res1.RowsAffected != res2.RowsAffected {
		fmt.Errorf("%v lines duplicate in database", res2.RowsAffected-res1.RowsAffected)
	}
}
