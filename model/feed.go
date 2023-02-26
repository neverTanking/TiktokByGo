package model

import (
	"github.com/neverTanking/TiktokByGo/db"
	"strconv"
	"time"
)

// GetVideoList generate VideoList from database by time sorted
func GetVideoList(lastTimeUnix string, curUser User, videoNumber int) ([]Video, error) {
	var dbVideoList []db.Video
	var videoList []Video
	//claim token, not use token_info now
	lastTimeUnixInt, err := strconv.Atoi(lastTimeUnix)
	//TODO: need determine whether lastTimeUnix valid
	if err != nil {
	}
	var lastTime = time.Unix(int64(lastTimeUnixInt), 0).UTC()

	//not login
	if isValid == false {
		//get video from the latest time
	}
	//TODO:personal feed for login user use user follow?
	res := db.DB.Where("created_at> ?", lastTime).Order("created_at DESC").Limit(videoNumber).Find(&dbVideoList)
	if res.Error != nil {

	}
	return videoList, nil
}
