package db

import (
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"sync"
)

type LikeState struct {
	UserId     uint
	VideoId    uint
	actionType uint
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

// 保证只初始化一次
func NewVideoDao() *VideoDao {
	videoOnce.Do(func() {
		videoDao = new(VideoDao)
	})
	return videoDao
}

// 增加一个赞
func (u *VideoDao) AddOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&db.Like{UserID: UserId, VideoID: VideoId}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 减少一个赞
func (u *VideoDao) SubOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND video_id = ?", UserId, VideoId).Delete(&db.Like{}).Error; err != nil {
			return err
		}
		return nil
	})
}
