package dao

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"log"
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
func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo db.User
	if err := db.DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.ID == 0 {
		return false
	}
	return true
}

func (u *UserInfoDAO) QueryUserIdByVideoIdInVideos(videoId int64, like *db.Like) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("video_id = ?", videoId).First(&like).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserInfoDAO) QueryUserInfoById(userId int64, user *db.User) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", userId).First(&user).Error; err != nil {
			return err
		}
		return nil
	})
}

// 获取用户发布了多少个作品
func (u *UserInfoDAO) QueryLenUserInfoById(userId int64) (int, error) {
	var user *[]*db.User
	db.DB.Where("id = ?", userId).Find(&user)
	if len(*user) == 0 {
		return 0, errors.New("用户没有发布作品")
	}
	return len(*user), nil
}
