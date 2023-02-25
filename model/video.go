package model

import (
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
)

var fakeVideo = Video{}

func SearchVideoByID(videoId uint, user User) (video Video, err error) {
	var curVideo db.Video
	res := db.DB.First(&curVideo, videoId)
	if res.Error != nil {
		return fakeVideo, fmt.Errorf("database err:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return fakeVideo, errNotFound
	}
	return Video{
		Id:            videoId,
		Author:        user,
		PlayUrl:       curVideo.PlayUrl,
		CoverUrl:      curVideo.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         curVideo.Title,
	}, nil
}
