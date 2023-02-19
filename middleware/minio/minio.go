package myminio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/neverTanking/TiktokByGo/config"
)

func Init() *minio.Client { //a
	endpoint := config.Miniourl
	accessKeyID := config.MinioaccessKey
	secretAccessKey := config.MiniosecretKey
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("创建 MinIO 客户端失败", err)

	}
	log.Printf("创建 MinIO 客户端成功")
	// minioClient.PutObject()
	return minioClient
}
