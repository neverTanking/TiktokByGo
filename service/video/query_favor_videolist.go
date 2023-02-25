package video

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/neverTanking/TiktokByGo/model/dao"
)

type FavorList struct {
	Videos []*model.Video `json:"video_list"`
}

func QueryFavorVideoList(userId int64) (*FavorList, error) {
	return NewQueryFavorVideoListFlow(userId).Do()
}

type QueryFavorVideoListFlow struct {
	userId int64

	videos    []*model.Video
	likes     []*db.Like
	videoList *FavorList
}

func NewQueryFavorVideoListFlow(userId int64) *QueryFavorVideoListFlow {
	return &QueryFavorVideoListFlow{userId: userId}
}

func (q *QueryFavorVideoListFlow) Do() (*FavorList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	//fmt.Println("Failed6666")
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	//fmt.Println(q.videos)
	return q.videoList, nil
}

func (q *QueryFavorVideoListFlow) checkNum() error {
	if !dao.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户状态异常")
	}
	return nil
}

func (q *QueryFavorVideoListFlow) prepareData() error {
	//查likes表里看看这个UserId喜欢的videoId
	//这边其实可以用Redis-set优化
	if err := dao.NewVideoDAO().QueryFavorVideoListByUserId(q.userId, &q.likes); err != nil {
		return err
	}

	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
	for i := range q.likes {
		//有了videoId,现在要在videos中查这个视频作者是谁
		var like db.Like
		//获取单个videoId
		if err := dao.NewUserInfoDAO().QueryUserIdByVideoIdInVideos(int64(q.likes[i].VideoID), &like); err != nil {
			return err
		}

		var db_video db.Video
		if err := dao.NewVideoDAO().QueryVideoInformationByVideoId(like.VideoID, &db_video); err != nil {
			return err
		}
		//fmt.Println("888888", like.UserID)
		//查询的是对的
		//return nil
		//作者信息查询
		var OneVideo model.Video
		OneVideo.Video = db_video
		var db_userInfo db.User
		var model_userInfo model.User
		var err error
		model_userInfo.FavoriteCount, err = Redis.NewRedisDao().GetUserFavoriteCount(uint(q.userId))
		if err != nil {
			return err
		}
		model_userInfo.FollowCount = 0
		model_userInfo.FollowerCount = 0
		model_userInfo.WorkCount, err = Redis.NewRedisDao().GetUserWorkCount(uint(q.userId))
		if err != nil {
			return err
		}

		model_userInfo.TotalFavorited = "0"
		err = dao.NewUserInfoDAO().QueryUserInfoById(int64(like.UserID), &db_userInfo)
		//return nil
		//更新videos里
		if err != nil {
			return err
		}
		model_userInfo.User = db_userInfo
		OneVideo.Author = model_userInfo
		OneVideo.FavoriteCount, err = Redis.NewRedisDao().GetLikeNumByVideoId(OneVideo.Video.ID)
		if err != nil { //如果找不到就是0
			OneVideo.FavoriteCount = 0
		}
		OneVideo.CommentCount = 0
		OneVideo.IsFavorite = true

		q.videos = append(q.videos, &OneVideo)

		fmt.Println(q.videos[0])
	}
	//fmt.Println(q.videos)
	return nil
}
func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}

	//fmt.Println(q.videoList)
	return nil
}
