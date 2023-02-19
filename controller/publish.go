package controller

import (
	"log"
	"net/http"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish1(c *gin.Context) {
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
	err = service.Publish_up(data, userId, title)

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
// func PublishList1(c *gin.Context) {
// 	c.JSON(http.StatusOK, VideoListResponse{
// 		Response: Response{
// 			StatusCode: 0,
// 		},
// 		VideoList: DemoVideos,
// 	})
// }
