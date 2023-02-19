package service

import (
	"github.com/neverTanking/TiktokByGo/db"
	//"middleware/ffmpeg"
	"log"
	"mime/multipart"

	"github.com/neverTanking/TiktokByGo/middleware/ffmpeg"
	uuid "github.com/satori/go.uuid"
)

type VideoServiceImpl struct {
	// UserService
	// LikeService
	// CommentService
}

// Publish
// 将传入的视频流保存在文件服务器中，并存储在mysql表中
func Publish_up(data *multipart.FileHeader, userId int64, title string) error {
	//将视频流上传到视频服务器，保存视频链接
	//生成一个uuid作为视频的名字
	videoName := uuid.NewV4().String()
	log.Printf("生成视频名称%v", videoName)
	err := db.VideoToMinio(data, videoName)
	if err != nil {
		log.Printf("方法dao.VideoFTP(file, videoName) 失败%v", err)
		return err
	}
	log.Printf("方法dao.VideoFTP(file, videoName) 成功")
	//在服务器上执行ffmpeg 从视频流中获取第一帧截图，并上传图片服务器，保存图片链接
	imageName := uuid.NewV4().String()
	//向队列中添加消息
	ffmpeg.Ffchan <- ffmpeg.Ffmsg{
		videoName,
		imageName,
	}
	//组装并持久化
	err = db.Save(videoName, imageName, userId, title)
	if err != nil {
		log.Printf("方法dao.Save(videoName, imageName, userId) 失败%v", err)
		return err
	}
	log.Printf("方法dao.Save(videoName, imageName, userId) 成功")
	return nil
}
