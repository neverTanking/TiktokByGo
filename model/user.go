package model

import (
	"github.com/neverTanking/TiktokByGo/controller"
	"github.com/neverTanking/TiktokByGo/db"
)

var user db.User

func CreatUser(username string, password string) int64 {
	user = db.User{
		Name:     username,
		Password: password,
	}
	res := db.DB.Create(&user)
	if res.Error != nil {
		return -1
	}
	return user.ID
}

func SearchUser(id int64) controller.User {
	res := db.DB.Find(&user)
	if res.Error != nil {
		return controller.User{}
	}
	return controller.User{
		Id:            id,
		Name:          user.Name,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}
