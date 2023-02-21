package video

import (
	"errors"
	"fmt"
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
	//fmt.Println("99999999", q.likes[0].VideoID)
	//return nil
	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
	for i := range q.likes {
		//有了videoId,现在要在videos中查这个视频作者是谁
		var like db.Like
		//获取单个video
		if err := dao.NewUserInfoDAO().QueryUserIdByVideoIdInVideos(int64(q.likes[i].VideoID), &like); err != nil {
			return err
		}

		//fmt.Println("888888", like.UserID)
		//查询的是对的
		//return nil
		//作者信息查询
		var db_userInfo db.User
		var model_userInfo model.User
		model_userInfo.FavoriteCount = 0
		model_userInfo.FollowCount = 0
		model_userInfo.FollowerCount = 0
		model_userInfo.WorkCount = 0
		model_userInfo.TotalFavorited = 0
		err := dao.NewUserInfoDAO().QueryUserInfoById(int64(like.UserID), &userInfo)
		fmt.Println(userInfo)
		//return nil
		//更新videos里
		if err == nil { //若查询未出错则更新，否则不更新作者信息
			q.videos[i].User = userInfo
		}
		q.videos[i].IsFavorite = true
		return nil
	}
	return nil
}
func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
