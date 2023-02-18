package dao

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
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

func (v *VideoDAO) QueryFavorVideoListByUserId(userId int64, videoList *[]*db.Video) error {
	if videoList == nil {
		return errors.New("QueryFavorVideoListByUserId videoList 空指针")
	}
	//多表查询，左连接得到结果，再映射到数据
	if err := db.DB.Raw("SELECT v.* FROM user_favor_videos u , videos v WHERE u.user_info_id = ? AND u.video_id = v.id", userId).Scan(videoList).Error; err != nil {
		return err
	}
	//如果id为0，则说明没有查到数据
	if len(*videoList) == 0 || (*videoList)[0].ID == 0 {
		return errors.New("点赞列表为空")
	}
	return nil
}
