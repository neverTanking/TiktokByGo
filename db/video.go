package db

import (
	"time"

	"github.com/neverTanking/TiktokByGo/config"
)

type TableVideo struct {
	Id          int64 `json:"id"`
	AuthorId    int64
	PlayUrl     string `json:"play_url"`
	CoverUrl    string `json:"cover_url"`
	PublishTime time.Time
	Title       string `json:"title"`
}

// TableName
//
//	将TableVideo映射到videos，
//	这样我结构体到名字就不需要是Video了，防止和我Service层到结构体名字冲突
func (TableVideo) TableName() string {
	return "videos"
}

// Save 保存视频记录
func SaveDao(videoName string, imageName string, authorId int64, title string) error {
	var video TableVideo
	video.PublishTime = time.Now()
	video.PlayUrl = config.PlayUrlPrefix + videoName + ".mp4"
	video.CoverUrl = config.CoverUrlPrefix + imageName + ".jpg"
	video.AuthorId = authorId
	video.Title = title
	result := DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetVideosByAuthorId(authorId int64) ([]TableVideo, error) {
	//建立结果集接收
	var data []TableVideo
	//初始化db
	//Init()
	result := DB.Where(&TableVideo{AuthorId: authorId}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// Save update value in database, if the value doesn't have primary key, will insert it
