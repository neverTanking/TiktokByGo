package video

import (
	"errors"
	"fmt"
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

	//测试ParameterVaild正确性
	//正确
	/*
		{
			fmt.Println("6666666666666666", u.UserId, u.VideoId, u.actionType)
		}
	*/

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
	//需要判断UserId是否存在，VideoId是否存在
	//需要判断这个是否已经存在
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		return err
	}
	if ok {
		fmt.Println("ERROR666666666!")
		return errors.New("you can't like again after you've already liked it")
	}
	if err := db.NewVideoDao().AddOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}

	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, true); err != nil {
		return err
	}
	return nil
}

// 点击取消喜欢
func (u *LikeState) UnLikeVideo() error {
	ok, err := Redis.NewRedisDao().GetLikeState(u.UserId, u.VideoId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("you can't cancel like again after you've already dislike it")
	}

	if err := db.NewVideoDao().SubOneLikeByUserIdAndVideoId(u.UserId, u.VideoId); err != nil {
		return err
	}
	if err := Redis.NewRedisDao().UpdatePostLike(u.UserId, u.VideoId, false); err != nil {
		return err
	}
	return err
}
