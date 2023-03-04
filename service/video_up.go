package service

import (
	"log"
	"mime/multipart"
	"sync"

	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/middleware/minio"
	"github.com/neverTanking/TiktokByGo/model"
)

type VideoServiceImpl struct {
	// UserService
	// LikeService
	// CommentService
}

// Publish
// 将传入的视频流保存在minio服务器中，并存储在mysql表中
func Publish_up(data *multipart.FileHeader, userId uint, title string, videoName string, imageName string) error {
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
	err = db.SaveDao(videoName, imageName, userId, title)
	if err != nil {
		log.Printf("方法db.Save(videoName, imageName, userId) 失败%v", err)
		return err
	}
	log.Printf("方法db.Save(videoName, imageName, userId) 成功")
	return nil
}

// List
// 通过userId来查询对应用户发布的视频，并返回对应的视频数组
func List(userId uint, curId uint) ([]model.Video, error) {
	//依据用户id查询所有的视频，获取视频列表
	data, err := db.GetVideosByAuthorId(userId)
	if err != nil {
		log.Printf("方法dao.GetVideosByAuthorId(%v)失败:%v", userId, err)
		return nil, err
	}
	log.Printf("方法dao.GetVideosByAuthorId(%v)成功:%v", userId, data)
	//提前定义好切片长度
	result := make([]model.Video, 0, len(data))
	//调用拷贝方法，将数据进行转换
	err = copyVideos(&result, &data, curId)
	if err != nil {
		log.Printf("方法videoService.copyVideos(&result, &data, %v)失败:%v", userId, err)
		return nil, err
	}
	//如果数据没有问题，则直接返回
	return result, nil
}

// 该方法可以将数据进行拷贝和转换，并从其他方法获取对应的数据
func copyVideos(result *[]model.Video, data *[]db.Video, curId uint) error {
	for _, temp := range *data {
		var video model.Video
		//将video进行组装，添加想要的信息,插入从数据库中查到的数据
		creatVideo(&video, &temp, curId)
		*result = append(*result, video)
	}
	return nil
}

// 将video进行组装，添加想要的信息,插入从数据库中查到的数据
func creatVideo(video *model.Video, data *db.Video, curId uint) {
	//建立协程组，当这一组的携程全部完成后，才会结束本方法
	var wg sync.WaitGroup
	wg.Add(4)
	var err error
	video.Video = *data
	//插入Author，这里需要将视频的发布者和当前登录的用户传入，才能正确获得isFollow，
	//如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	go func() {
		video.Author, err = GetUserByIdWithCurId(data.UserID, curId)
		if err != nil {
			log.Printf("方法videoService.GetUserByIdWithCurId(data.AuthorId, userId) 失败：%v", err)
		} else {
			log.Printf("方法videoService.GetUserByIdWithCurId(data.AuthorId, userId) 成功")
		}
		wg.Done()
	}()

	//插入点赞数量，同上所示，不将nil直接向上返回，数据没有就算了，给一个默认就行了
	go func() {
		video.FavoriteCount, err = FavouriteCount(data.ID)
		if err != nil {
			log.Printf("方法videoService.FavouriteCount(data.ID) 失败：%v", err)
		} else {
			log.Printf("方法videoService.FavouriteCount(data.ID) 成功")
		}
		wg.Done()
	}()

	//获取该视屏的评论数字
	go func() {
		video.CommentCount, err = CountFromVideoId(data.ID)
		if err != nil {
			log.Printf("方法videoService.CountFromVideoId(data.ID) 失败：%v", err)
		} else {
			log.Printf("方法videoService.CountFromVideoId(data.ID) 成功")
		}
		wg.Done()
	}()

	//获取当前用户是否点赞了该视频
	go func() {
		video.IsFavorite, err = IsFavourite(video.ID, curId)
		if err != nil {
			log.Printf("方法videoService.IsFavourit(video.Id, userId) 失败：%v", err)
		} else {
			log.Printf("方法videoService.IsFavourit(video.Id, userId) 成功")
		}
		wg.Done()
	}()

	wg.Wait()
}
