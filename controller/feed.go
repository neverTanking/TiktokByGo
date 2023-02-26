package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/middleware/JWT"
	"github.com/neverTanking/TiktokByGo/model"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	curUser := model.User{}
	var err error
	lastTime := c.DefaultQuery("next_time", strconv.FormatInt(time.Now().Add(time.Minute*-30).Unix(), 10))
	user, ok := JWT.TokenToClaim(c.DefaultQuery("token", "noToken"))
	if ok == true {
		curUser, err = model.SearchUserByID(user.UserId)
		if err != nil {
			if errors.Is(err, errNotFound) {
				curUser = model.User{}
			}
			fmt.Errorf("search user error:%v", err)
		}
	}
	videoList, err := model.GetVideoList(lastTime, curUser, 5)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
