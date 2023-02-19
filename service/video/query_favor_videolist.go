package video

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model/dao"
)

type FavorList struct {
	Videos []*db.Video `json:"video_list"`
}

func QueryFavorVideoList(userId int64) (*FavorList, error) {
	return NewQueryFavorVideoListFlow(userId).Do()
}

type QueryFavorVideoListFlow struct {
	userId int64

	videos []*db.Video

	videoList *FavorList
}

func NewQueryFavorVideoListFlow(userId int64) *QueryFavorVideoListFlow {
	return &QueryFavorVideoListFlow{userId: userId}
}

func (q *QueryFavorVideoListFlow) Do() (*FavorList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	/*if err := q.prepareData(); err != nil {
		return nil, err
	}*/
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

/*
	func (q *QueryFavorVideoListFlow) prepareData() error {
		err := dao.NewVideoDAO().QueryFavorVideoListByUserId(q.userId, &q.videos)
		if err != nil {
			return err
		}
		//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
		for i := range q.videos {
			//作者信息查询
			var userInfo db.User
			err = dao.NewUserInfoDAO().QueryUserInfoById(int64(q.videos[i].ID), &userInfo)
			if err == nil { //若查询未出错则更新，否则不更新作者信息
				q.videos[i].User = userInfo
			}
			q.videos[i].IsFavorite = true
		}
		return nil
	}
*/
func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
