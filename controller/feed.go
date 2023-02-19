package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.DefaultQuery("next_time", strconv.FormatInt(time.Now().Add(time.Minute*-30).Unix(), 10))
	token := c.DefaultQuery("token", "noToken")
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoConvert(model.GetVideoList(lastTime, token, 5)),
		NextTime:  time.Now().Unix(),
	})
}

func videoConvert(video []db.Video) []Video {
	var videoList []Video
	return videoList
}
