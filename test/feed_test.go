package test

import (
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func TestCreateDemoDB(t *testing.T) {
	db.Init()
	user1 := db.User{
		Model:    gorm.Model{},
		Name:     "user1",
		Password: "password",
		Videos:   nil,
		Likes:    nil,
	}
	user2 := db.User{
		Model:    gorm.Model{},
		Name:     "user2",
		Password: "password",
		Videos:   nil,
		Likes:    nil,
	}
	db.DB.Create(&user1)
	db.DB.Create(&user2)

	var user = db.User{}
	db.DB.First(&user)

	video1 := db.Video{
		Model:    gorm.Model{},
		UserID:   user.ID,
		PlayUrl:  "https://1",
		CoverUrl: "https://1_1",
		Title:    "Title1",
	}
	video2 := db.Video{
		Model:    gorm.Model{},
		UserID:   user.ID,
		PlayUrl:  "https://2",
		CoverUrl: "https://2_1",
		Title:    "Title2",
	}
	video3 := db.Video{
		Model:    gorm.Model{},
		UserID:   user.ID,
		PlayUrl:  "https://3",
		CoverUrl: "https://3_1",
		Title:    "Title3",
	}
	db.DB.Create(&video1)
	db.DB.Create(&video2)
	db.DB.Create(&video3)
}
func TestGetVideoList(t *testing.T) {
	db.Init()
	curUser := model.User{
		ID:   1,
		Name: "user1",
	}

	model.GetVideoList(strconv.FormatInt(time.Now().Add(time.Minute*-180).Unix(), 10), curUser, 5)
}
func TestCleanDB(t *testing.T) {
	db.Init()
	db.DB.Exec("DELETE FROM videos")
	db.DB.Exec("DELETE FROM users")
}
