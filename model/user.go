package model

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
)

var curUser db.User
var fakeUserId uint = 0
var fakeUser = User{}
var errExistedUser = errors.New("user exist")

//TODO: 需要一个检查数据库中是否有重复数据的维护函数

func CreatUser(username string, password string) (userID uint, err error) {
	//确保用户名不会重复
	_, exist := SearchUserByName(username)
	if exist != false { //已经存在数据
		return fakeUserId, errExistedUser
	}
	curUser = db.User{
		Name:     username,
		Password: password,
	}
	res := db.DB.Create(&curUser)
	if res.Error != nil {
		return fakeUserId, fmt.Errorf("create user %v failed and db error: %v", username, res.Error)
	}
	return curUser.ID, nil
}

func SearchUserByID(id uint) (user User, ok bool) {
	res := db.DB.Find(&curUser, id)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			return fakeUser, false
		}
	}
	return User{
		ID:              curUser.ID,
		Name:            curUser.Name,
		Password:        curUser.Password,
		Avatar:          curUser.Avatar,
		BackgroundImage: curUser.BackgroundImage,
		Signature:       curUser.Signature,
		FavoriteCount:   0,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  "",
		WorkCount:       0,
		IsFollow:        false,
	}, true

}
func SearchUserByName(username string) (user User, ok bool) {
	res := db.DB.Where("name = ?", username).Find(&curUser)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			return fakeUser, false
		}
	}
	return User{
		ID:              curUser.ID,
		Name:            curUser.Name,
		Password:        curUser.Password,
		Avatar:          curUser.Avatar,
		BackgroundImage: curUser.BackgroundImage,
		Signature:       curUser.Signature,
		FavoriteCount:   0,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  "",
		WorkCount:       0,
		IsFollow:        false,
	}, true
}
