package ffm

import (
	"context"

	"bytes"
	"log"
	"os"
	"os/exec"

	"github.com/minio/minio-go/v7"
	"github.com/neverTanking/TiktokByGo/config"
)

// Ffmpeg 调用ffmpeg命令来创建视频截图
func Ffmpeg(minioclient *minio.Client, imageName string, videoName string) error {

	bucketName := config.BucketFfmpeg
	location := minio.MakeBucketOptions{Region: "cn-north-1"}

	minioc := minioclient
	err := minioc.MakeBucket(context.Background(), bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioc.BucketExists(context.Background(), bucketName)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", bucketName)
		} else {
			log.Fatalln("查询存储桶状态异常", err)
			return err
		}
	}
	log.Printf("创建存储桶 %s 成功", bucketName)
	//-i video.mp4  -y  -f  image2   -ss  60   -vframes  1  video.png
	//-ss 00:00:01 -i /home/ftpuser/video/" + videoName + ".mp4 -vframes 1 /home/ftpuser/images/" + imageName + ".jpg"
	args := []string{"-ss 00:00:01 -y -i ~/videorepo/" + videoName + ".mp4 -vframes 1 ~/ffmpgex/" + imageName + ".jpg"}
	cmd := exec.Command("ffmpeg", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
	// 使用PutObject上传一个文件
	ffmpeg, err := os.OpenFile("~/ffmpgex/"+imageName+".jpg", os.O_RDWR, 6)
	if err != nil {
		log.Printf("方法file.Open() 失败%v", err)
		return err
	}
	log.Printf("方法file.Open() 成功")
	_, err = minioc.PutObject(context.Background(), bucketName, imageName+".jpg", ffmpeg, -1, minio.PutObjectOptions{ContentType: ""})
	if err != nil {
		log.Printf("upload video failed", err)
		return err
	}

	return nil
}
