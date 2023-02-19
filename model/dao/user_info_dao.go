package dao

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
	"sync"
)

type UserInfoDAO struct {
}

var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

func (u *UserInfoDAO) QueryUserInfoById(userId int64, user *db.User) error {
	if user == nil {
		return ErrIvdPtr
	}
	//DB.Where("id=?",userId).First(userinfo)
	db.DB.Where("id=?", userId).Select([]string{"ID", "name"}).First(user)
	//id为零值，说明sql执行失败
	if user.ID == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}
