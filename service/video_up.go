package service

import (
	"log"
	"mime/multipart"

	"github.com/neverTanking/TiktokByGo/middleware/minio"

)

type VideoServiceImpl struct {
	// UserService
	// LikeService
	// CommentService
}

// Publish
// 将传入的视频流保存在minio服务器中，并存储在mysql表中
func Publish_up(data *multipart.FileHeader, userId int64, title string,videoName string, imageName string) error {
	//将视频流上传到视频服务器，保存视频链接
	//在服务器上执行ffmpeg 从视频中获取第一帧截图，并上传minio服务器，保存图片链接

	log.Printf("生成视频名称%v", videoName)
	err := minio.VideoToMinio(data, videoName, imageName, title)
	if err != nil {
		log.Printf("VideoToMinio 失败%v", err)
		return err
	}
	log.Printf("VideoToMinio 成功")

	//组装并持久化
	err = minio.Save(videoName, imageName, userId, title)
	if err != nil {
		log.Printf("方法db.Save(videoName, imageName, userId) 失败%v", err)
		return err
	}
	log.Printf("方法db.Save(videoName, imageName, userId) 成功")
	return nil
}
