package dao

import (
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

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
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

// 根据VideoId查询视频是否存在
func (u *VideoDAO) IsVideoExistById(id int64) bool {
	var videoInfo db.Video
	if err := db.DB.Where("id=?", id).First(&videoInfo).Error; err != nil {
		log.Println(err)
	}
	if videoInfo.ID == 0 {
		return false
	}
	return true
}
