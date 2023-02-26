package dao

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"sync"
)

type LikeDAO struct {
}

var (
	likeDAO  *LikeDAO
	likeOnce sync.Once
)

func NewLikeDAO() *LikeDAO {
	likeOnce.Do(func() {
		likeDAO = new(LikeDAO)
	})
	return likeDAO
}

// 判断是否存在userId点赞了videoId
// 找不到就是没点赞
func (u *LikeDAO) IsLikeByUserIdAndVideoId(userId uint, videoId uint) bool {
	var like db.Like
	db.DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).First(&like)
	if like.ID == 0 {
		return false //没找到
	}
	return true
}

// 查likes表里看看这个UserId喜欢的视频Id
func (v *LikeDAO) QueryFavorVideoListByUserId(userId int64, Like *[]*db.Like) error {
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

// 查likes表里看看这个UserId喜欢的视频个数
func (v *LikeDAO) QueryLenFavorVideoListByUserId(userId int64) (int, error) {
	var likeList *[]*db.Like
	err := db.DB.Where("user_id = ?", userId).Find(&likeList).Error
	if err != nil {
		return 0, err
	}
	if len(*likeList) == 0 {
		return 0, errors.New("点赞列表为空")
	}
	return len(*likeList), nil
}

// 查看视频点赞总数
func (u *LikeDAO) QueryLenFavorVideoListByVideoId(videoId int64) (int, error) {
	var likeList *[]*db.Like
	err := db.DB.Where("video_id=?", videoId).Find(&likeList).Error
	if err != nil {
		return 0, err
	}
	if len(*likeList) == 0 {
		return 0, errors.New("没有人给这个视频点赞")
	}
	return len(*likeList), nil
}

// 增加一个赞
func (u *LikeDAO) AddOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&db.Like{UserID: UserId, VideoID: VideoId}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 减少一个赞
func (u *LikeDAO) SubOneLikeByUserIdAndVideoId(UserId uint, VideoId uint) error {
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
