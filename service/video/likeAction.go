package video

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	dao "github.com/neverTanking/TiktokByGo/model/dao"
)

const (
	LIKE    = 1
	DISLIKE = 2
)

type LikeState struct {
	UserId     uint
	VideoId    uint
	actionType int
}

func LikeAction(userid uint, videoid uint, actiontype int) error {
	return NewLikeState(userid, videoid, actiontype).Finish()
}

func NewLikeState(userid uint, videoid uint, actiontype int) *LikeState {
	return &LikeState{
		UserId:     userid,
		VideoId:    videoid,
		actionType: actiontype,
	}
}

func (u *LikeState) Finish() error {

	if err := u.ParameterValid(); err != nil {
		return err
	}
	//因为前面已经判断了,只能是LIKE or UNLIKE
	if u.actionType == LIKE {
		if err := u.LikeVideo(); err != nil {
			return err
		}
	} else {
		if err := u.UnLikeVideo(); err != nil {
			return err
		}
	}
	return nil
}

func (u *LikeState) ParameterValid() error {
	//根据UserId查询用户是否存在，需要用户组写
	exists := dao.NewUserInfoDAO().IsUserExistById(int64(u.UserId))
	if !exists {
		return errors.New("User Not Exists")
	}
	//判断actionType是否合法
	if u.actionType != LIKE && u.actionType != DISLIKE {
		return errors.New("actionType illegal")
	}
	return nil
}

// 点击喜欢
func (u *LikeState) LikeVideo() error {
	//需要判断UserId是否存在，VideoId是否存在
	//需要判断这个记录是否已经存在
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		//在Redis中没找到
		//去数据库中找like表
		ok = dao.NewLikeDAO().IsLikeByUserIdAndVideoId(u.UserId, u.VideoId)
	}
	if ok {
		return errors.New("you can't like again after you've already liked it")
	}
	if err := dao.NewLikeDAO().AddOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}
	//如果不存在会先赋值为0然后加1
	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, true); err != nil {
		return err
	}
	//给视频的喜欢总数加1
	if err := Redis.NewRedisDao().AddOneLikeNumByVideoId(u.VideoId); err != nil {
		return err
	}
	return nil
}

// 点击取消喜欢
func (u *LikeState) UnLikeVideo() error {
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		//在Redis中没找到
		//去数据库中找like表
		ok = dao.NewLikeDAO().IsLikeByUserIdAndVideoId(u.UserId, u.VideoId)
	}
	if !ok {
		return errors.New("you can't cancel like again after you've already dislike it")
	}

	if err := dao.NewLikeDAO().SubOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}
	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, false); err != nil {
		return err
	}
	if err := Redis.NewRedisDao().SubOneLikeNumByVideoId(u.VideoId); err != nil {
		return err
	}
	return err
}
