package model

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
)

var fakeUserId uint = 0
var fakeUser = User{}
var errExistedUser = errors.New("user exist")
var errNotFound = errors.New("user not found")
var errWrongPassword = errors.New("wrong password")

//TODO: 需要一个检查数据库中是否有重复数据的维护函数

func CreatUser(username string, password string) (userID uint, err error) {
	//确保用户名不会重复
	var curUser db.User
	var _, err1 = SearchUserByName(username)
	if err1 == nil {
		return fakeUserId, errExistedUser
	}
	if !errors.Is(err1, errNotFound) {
		return fakeUserId, fmt.Errorf("unknown search user err, create stop:%v", err1)
	}
	curUser = db.User{
		Name:     username,
		Password: password,
	}
	res := db.DB.Create(&curUser)
	if res.Error != nil {
		return fakeUserId, fmt.Errorf("create user %v failed beacuse database error: %v", username, res.Error)
	}
	return curUser.ID, nil
}

func SearchUserByID(id uint) (user User, err error) {
	var curUser db.User
	res := db.DB.First(&curUser, id)
	if res.Error != nil {
		return fakeUser, fmt.Errorf("database err:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return fakeUser, errNotFound
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
		TotalFavorited:  "",
		WorkCount:       0,
		IsFollow:        false,
	}, nil
}

func SearchUserByName(username string) (user User, err error) {
	var curUser db.User
	res := db.DB.Where("name = ?", username).Find(&curUser)
	if res.Error != nil {
		return fakeUser, fmt.Errorf("database err:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return fakeUser, errNotFound
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
		TotalFavorited:  "",
		WorkCount:       0,
		IsFollow:        false,
	}, nil
}

func SearchUserForVerify(id uint, password string) error {
	var curUser db.User
	res := db.DB.Find(&curUser, id)
	//
	if res.Error != nil {
		return fmt.Errorf("database error:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return errNotFound
	}
	if curUser.Password != password {
		return errWrongPassword
	}

	return nil

}
