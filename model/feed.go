package model

import (
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/middleware/JWT"
	"strconv"
	"time"
)

// GetVideoList generate VideoList from database by time sorted
func GetVideoList(lastTimeUnix string, token string, videoNumber int) []db.Video {
	var videoList []db.Video
	//claim token, not use token_info now
	lastTimeUnixInt, err := strconv.Atoi(lastTimeUnix)
	//TODO: need determine whether lastTimeUnix valid
	if err != nil {
	}
	var lastTime = time.Unix(int64(lastTimeUnixInt), 0).UTC()

	_, isValid := JWT.TokenToClaim(token)

	//not login
	if isValid == false {
		//get video from the latest time
	}
	//TODO:personal feed for login user use user follow?
	res := db.DB.Where("CreatedAt > ?", lastTime).Order("CreatedAt DESC").Limit(videoNumber).Find(&videoList)
	fmt.Print(res.Error)
	fmt.Print(res.RowsAffected)

	return videoList
}
