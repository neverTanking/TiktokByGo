package controller

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/neverTanking/TiktokByGo/service"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	data, err := c.FormFile("data")

	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	log.Printf("获取到用户id:%v\n", userId)
	title := c.PostForm("title")
	log.Printf("获取到视频title:%v\n", title)
	if err != nil {
		log.Printf("获取视频流失败:%v", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//
	//生成一个uuid作为视频的名字
	videoName := uuid.NewV4().String()
	//生成一个uuid作为图片的名字
	imageName := uuid.NewV4().String()
	ffmpegdst := path.Join("~/videorepo", videoName+".mp4")
	c.SaveUploadedFile(data, ffmpegdst)
	err = service.Publish_up(data, userId, title, videoName, imageName)

	if err != nil {
		log.Printf("方法videoService.Publish(data, userId) 失败：%v", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Printf("方法videoService.Publish(data, userId) 成功")

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// // PublishList all users have same publish video list
//
//	func PublishList1(c *gin.Context) {
//		c.JSON(http.StatusOK, VideoListResponse{
//			Response: Response{
//				StatusCode: 0,
//			},
//			VideoList: DemoVideos,
//		})
//	}
//
// PublishList /publish/list/
func PublishList(c *gin.Context) {
	user_Id, _ := c.GetQuery("user_id")
	userId, _ := strconv.ParseInt(user_Id, 10, 64)
	log.Printf("获取到用户id:%v\n", userId)
	curId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	log.Printf("获取到当前用户id:%v\n", curId)

	list, err := service.List(userId, curId)
	if err != nil {
		log.Printf("调用videoService.List(%v)出现错误：%v\n", userId, err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频列表失败"},
		})
		return
	}
	log.Printf("调用videoService.List(%v)成功", userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: list,
	})
}
