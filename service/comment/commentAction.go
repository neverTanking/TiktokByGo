package comment

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/neverTanking/TiktokByGo/model/dao"
)

type Response struct {
	MyComment *model.Comment `json:"comment"`
}

func QueryCommentAction(userId uint, videoId uint, actionType int, commentText string, commentId uint) (*Response, error) {
	return NewCommentAction(userId, videoId, actionType, commentText, commentId).Finish()
}

type commentAction struct {
	userId      uint
	videoId     uint
	actionType  int
	commentText string
	commentId   uint

	*Response
	comment *model.Comment
}

func NewCommentAction(userId uint, videoId uint, actionType int, commentText string, commentId uint) *commentAction {
	return &commentAction{userId: userId, videoId: videoId, actionType: actionType, commentText: commentText, commentId: commentId}
}

func (u *commentAction) Finish() (*Response, error) {
	if err := u.checkNum(); err != nil {
		return nil, err
	}
	if err := u.prepareData(); err != nil {
		return nil, err
	}
	if err := u.packData(); err != nil {
		return nil, err
	}
	return u.Response, nil
}

func (u *commentAction) checkNum() error {
	if !dao.NewUserInfoDAO().IsUserExistById(int64(u.userId)) {
		return errors.New("用户状态异常")
	}
	if !dao.NewVideoDAO().IsVideoExistById(int64(u.videoId)) {
		return errors.New("视频状态异常")
	}
	return nil
}

// 操作数据库
func (u *commentAction) prepareData() error {
	var OneCommentAction model.Comment
	var model_userInfo model.User
	var err error
	model_userInfo, err = model.SearchUserByID(u.userId)
	if err != nil {
		return err
	}
	model_userInfo.FavoriteCount, err = Redis.NewRedisDao().GetUserFavoriteCount(u.userId)
	model_userInfo.WorkCount, err = Redis.NewRedisDao().GetUserWorkCount(u.userId)
	model_userInfo.TotalFavorited = "0"
	if u.actionType == 1 { //用户插入评论内容
		if err := dao.NewCommentDAO().InsertCommentByUserIdAndVideoIdAndCommentText(u.userId, u.videoId, u.commentText); err != nil {
			return err
		}
		OneCommentAction, err = model.SearchCommentByID(u.commentId, model_userInfo, u.actionType)
		if err != nil {
			return err
		}
	} else { //用户删除评论内容
		//先判断这个CommentId在不在
		if ok := dao.NewCommentDAO().IsExistsCommentId(u.commentId); !ok {
			return errors.New("CommentId 不存在")
		}
		OneCommentAction, err = model.SearchCommentByID(u.commentId, model_userInfo, u.actionType)
		if err != nil {
			return err
		}
		if err := dao.NewCommentDAO().DeleteCommentByCommentId(u.commentId); err != nil {
			return err
		}
	}
	u.comment = &OneCommentAction
	return nil
}

func (u *commentAction) packData() error {
	u.Response = &Response{MyComment: u.comment}
	return nil
}
