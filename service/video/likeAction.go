package video

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	db "github.com/neverTanking/TiktokByGo/db/video"
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
		u.LikeVideo()
	} else {
		u.UnLikeVideo()
	}
	return nil
}

func (u *LikeState) ParameterValid() error {
	//查询用户是否存在，需要用户组写
	//...
	//判断actionType是否合法
	if u.actionType != LIKE && u.actionType != DISLIKE {
		return errors.New("actionType illegal")
	}
	return nil
}

// 点击喜欢
func (u *LikeState) LikeVideo() error {
	if err := db.NewVideoDao().AddOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("you can't like again after you've already liked it")
	}
	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, true); err != nil {
		return err
	}
	return nil
}

// 点击取消喜欢
func (u *LikeState) UnLikeVideo() error {
	if err := db.NewVideoDao().SubOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("you can't cancel like again after you've already dislike it")
	}
	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, false); err != nil {
		return err
	}
	return err
}
