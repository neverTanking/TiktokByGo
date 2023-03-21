package video

import (
	"errors"
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
	if err := dao.NewLikeDAO().QueryFavorVideoListByUserId(q.userId, &q.likes); err != nil {
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
		var err error
		var model_userInfo model.User
		model_userInfo, err = model.SearchUserByID(uint(q.userId))
		if err != nil {
			return err
		}
		var OneVideo model.Video

		OneVideo, err = model.SearchVideoByID(like.VideoID, model_userInfo)
		model_userInfo, err = model.SearchUserByID(uint(q.userId))
		if err := dao.NewVideoDAO().QueryVideoInformationByVideoId(like.VideoID, &db_video); err != nil {
			return err
		}
		//更新videos里
		if err != nil {
			return err
		}
		model_userInfo.FavoriteCount, err = Redis.NewRedisDao().GetUserFavoriteCount(uint(q.userId))
		if err != nil {
			cnt, err := dao.NewLikeDAO().QueryLenFavorVideoListByUserId(q.userId)
			if err != nil {
				model_userInfo.FavoriteCount = 0
			} else {
				model_userInfo.FavoriteCount = int64(cnt)
			}
		}
		model_userInfo.FollowCount = 0
		model_userInfo.FollowerCount = 0
		model_userInfo.WorkCount, err = Redis.NewRedisDao().GetUserWorkCount(uint(q.userId))
		if err != nil {
			cnt, err := dao.NewUserInfoDAO().QueryLenUserInfoById(q.userId)
			if err != nil {
				model_userInfo.WorkCount = 0
			} else {
				model_userInfo.WorkCount = int64(cnt)
				//找到了给Redis设置这个cnt
				Redis.NewRedisDao().SetUserWorkCount(uint(q.userId), int64(cnt))
			}
		}
		OneVideo.FavoriteCount, err = Redis.NewRedisDao().GetLikeNumByVideoId(OneVideo.Id)
		if err != nil {
			cnt, err := dao.NewLikeDAO().QueryLenFavorVideoListByVideoId(int64(OneVideo.Id))
			if err != nil {
				OneVideo.FavoriteCount = 0
			} else {
				model_userInfo.FavoriteCount = int64(cnt)
				//找到了给Redis设置这个cnt
				Redis.NewRedisDao().SetUserFavoriteCount(OneVideo.Id, int64(cnt))
			}
		}
		OneVideo.CommentCount, err = Redis.NewRedisDao().GetCommentByVideoId(OneVideo.Id)
		if err != nil {
			cnt, err := dao.NewCommentDAO().QueryLenCommentByVideoId(OneVideo.Id)
			if err != nil {
				OneVideo.CommentCount = 0
			} else {
				OneVideo.CommentCount = cnt
				Redis.NewRedisDao().SetCommentByVideoId(OneVideo.Id, cnt)
			}
		}
		OneVideo.IsFavorite = true
		model_userInfo.TotalFavorited = "0"

		OneVideo.Author = model_userInfo

		q.videos = append(q.videos, &OneVideo)
	}
	return nil
}
func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
