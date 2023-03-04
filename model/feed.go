package model

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	"github.com/neverTanking/TiktokByGo/db"
	"strconv"
	"time"
)

// GetVideoList generate VideoList from database by time sorted
func GetVideoList(lastTimeUnix string, curUser User, videoNumber int) ([]Video, error) {
	var dbVideoList []db.Video
	var videoList []Video

	//convert time
	lastTimeUnixInt, _ := strconv.Atoi(lastTimeUnix)
	lastTime := time.Unix(int64(lastTimeUnixInt), 0).UTC()

	//TODO:personal feed for login user use user follow?
	res := db.DB.Where("created_at> ?", lastTime).Order("created_at DESC").Limit(videoNumber).Find(&dbVideoList)
	if res.Error != nil {
		return videoList, fmt.Errorf("database error:%v", res.Error)
	}
	if int64(videoNumber) != res.RowsAffected {
		if res.RowsAffected == 0 {
			res1 := db.DB.Order("created_at DESC").Limit(videoNumber).Find(&dbVideoList)
			if res1.Error != nil {
				return videoList, fmt.Errorf("database error:%v", res.Error)
			}
			if res1.RowsAffected == 0 {
				return videoList, errors.New("no new videos")
			}
		}
	}
	for _, video := range dbVideoList {
		author, _ := SearchUserByID(video.UserID)
		favoriteCount, err := Redis.NewRedisDao().GetUserFavoriteCount(video.UserID)
		if err != nil {
			favoriteCount = 0
		}
		commentCount, err := Redis.NewRedisDao().GetCommentByVideoId(video.ID)
		if err != nil {
			commentCount = 0
		}
		isFavorite, err := Redis.NewRedisDao().GetLikeState(curUser.ID, video.ID)
		if err != nil {
			isFavorite = false
		}
		videoList = append(videoList, Video{
			Id:            video.ID,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int(favoriteCount),
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoList, nil
}
