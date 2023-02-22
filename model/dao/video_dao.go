package dao

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"log"
	"sync"
)

type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

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

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}

// 查likes表里看看这个UserId喜欢的视频Id
func (v *VideoDAO) QueryFavorVideoListByUserId(userId int64, Like *[]*db.Like) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userId).Find(&Like).Error; err != nil {
			return err
		}
		if len(*Like) == 0 {
			return errors.New("点赞列表为空")
		}
		return nil
	})
}

// 增加一个赞
func (u *VideoDAO) AddOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&db.Like{UserID: UserId, VideoID: VideoId}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 减少一个赞
func (u *VideoDAO) SubOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM `likes` WHERE `user_id` = ? AND `video_id` = ?", UserId, VideoId).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ? AND video_id = ?", UserId, VideoId).Delete(&db.Like{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据videoId查video表所有信息
func (u *VideoDAO) QueryVideoInformationByVideoId(videoId uint, video *db.Video) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", videoId).First(&video).Error; err != nil {
			return err
		}
		return nil
	})
}
