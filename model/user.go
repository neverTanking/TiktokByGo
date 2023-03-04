package model

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
)

var curUser db.User
var fakeUserId uint = 0
var fakeUser = User{}

func CreatUser(username string, password string) (userID uint, err error) {
	curUser = db.User{
		Name:     username,
		Password: password,
	}
	res := db.DB.Create(&curUser)
	if res.Error != nil {
		return fakeUserId, res.Error
	}
	return curUser.ID, errors.New("no error")
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
		Avatar:          curUser.Avatar,
		BackgroundImage: curUser.BackgroundImage,
		Signature:       curUser.Signature,
		FavoriteCount:   0,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  0,
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
		Avatar:          curUser.Avatar,
		BackgroundImage: curUser.BackgroundImage,
		Signature:       curUser.Signature,
		FavoriteCount:   0,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  0,
		WorkCount:       0,
		IsFollow:        false,
	}, true
}
