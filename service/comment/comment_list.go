package comment

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/neverTanking/TiktokByGo/model/dao"
)

type List struct {
	Comments []*model.Comment `json:"comment_list"`
}

func QueryCommentList(userId uint, videoId uint) (*List, error) {
	return NewQueryCommentListFlow(userId, videoId).Finish()
}

type QueryCommentListFlow struct {
	userId  uint
	videoId uint

	comments []*model.Comment

	commentList *List
}

func NewQueryCommentListFlow(userId uint, videoId uint) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

func (u *QueryCommentListFlow) Finish() (*List, error) {
	if err := u.checkNum(); err != nil {
		return nil, err
	}
	if err := u.prepareData(); err != nil {
		return nil, err
	}
	if err := u.packData(); err != nil {
		return nil, err
	}
	return u.commentList, nil
}

func (u *QueryCommentListFlow) checkNum() error {
	if !dao.NewUserInfoDAO().IsUserExistById(int64(u.userId)) {
		return errors.New("u.userId 用户不存在")
	}
	if !dao.NewVideoDAO().IsVideoExistById(int64(u.videoId)) {
		return errors.New("u.videoId 视频不存在")
	}
	return nil
}

func (u *QueryCommentListFlow) prepareData() error {
	var db_comments []*db.Comment
	if err := dao.NewCommentDAO().QueryCommentListByVideoId(u.videoId, &db_comments); err != nil {
		return err
	}
	for i := range db_comments {
		var comment model.Comment
		var err error
		comment.User_, err = model.SearchUserByID(u.userId)
		if err != nil {
			return err
		}
		comment, err = model.SearchCommentByID(db_comments[i].ID, comment.User_, 2)
		if err != nil {
			return err
		}
		comment.User_.FavoriteCount, err = Redis.NewRedisDao().GetUserFavoriteCount(comment.User_.ID)
		if err != nil {
			return err
		}
		comment.User_.WorkCount, err = Redis.NewRedisDao().GetUserWorkCount(comment.User_.ID)
		if err != nil {
			return err
		}
		u.comments = append(u.comments, &comment)
	}
	return nil
}

func (u *QueryCommentListFlow) packData() error {
	u.commentList = &List{Comments: u.comments}
	return nil
}
