package db

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go"
	"github.com/neverTanking/TiktokByGo/config"
	"github.com/neverTanking/TiktokByGo/middleware/myminio"
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

func VideoToMinio(file *multipart.FileHeader, videoname string) error {
	// 初使化 minio client对象
	minioClient, err := myminio.Init()
	if err != nil {
		return err
	}

	// 创建一个叫 mybucket 的存储桶。
	bucketName := config.BucketName
	location := config.Location

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", bucketName)
		} else {
			log.Fatalln("查询存储桶状态异常", err)
			return err
		}
	}
	log.Printf("创建存储桶 %s 成功", bucketName)

	// 使用PutObject上传一个文件
	video, err := file.Open()
	if err != nil {
		log.Printf("方法file.Open() 失败%v", err)
		return err
	}

	log.Printf("方法file.Open() 成功")
	_, err2 := minioClient.PutObject(bucketName, videoname+".mp4", video, -1, minio.PutObjectOptions{ContentType: "Progress"})
	if err != nil {
		log.Printf("upload video failed", err)
		return err2
	}

	return nil

}
