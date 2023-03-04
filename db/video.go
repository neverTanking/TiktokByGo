package db

import (
	"github.com/neverTanking/TiktokByGo/config"
)

// Save 保存视频记录
func SaveDao(videoName string, imageName string, authorId uint, title string) error {
	var video Video

	video.PlayUrl = config.PlayUrlPrefix + videoName + ".mp4"
	video.CoverUrl = config.CoverUrlPrefix + imageName + ".jpg"
	video.UserID = authorId
	video.Title = title
	result := DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetVideosByAuthorId(authorId uint) ([]Video, error) {
	//建立结果集接收
	var data []Video
	//初始化db
	//Init()
	result := DB.Where(&Video{UserID: authorId}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// Save update value in database, if the value doesn't have primary key, will insert it
