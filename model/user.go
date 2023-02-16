package model

import (
	"github.com/neverTanking/TiktokByGo/db"
)

var user db.User

/**
 * @description: 创建用户
 * @param {string} username
 * @param {string} password
 * @return {*} user_id
 */
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

func SearchUser(id int64) User {
	res := db.DB.Find(&user)
	if res.Error != nil {
		return User{}
	}
	return User{
		Id:            id,
		Name:          user.Name,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}
