package comment

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/cache/Redis"
	"github.com/neverTanking/TiktokByGo/db"
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
	var db_userInfo db.User
	var model_userInfo model.User
	var err error
	model_userInfo.FavoriteCount, err = Redis.NewRedisDao().GetUserFavoriteCount(u.userId)
	if err != nil {
		return err
	}
	model_userInfo.FollowCount = 0
	model_userInfo.FollowerCount = 0
	model_userInfo.WorkCount, err = Redis.NewRedisDao().GetUserWorkCount(u.userId)
	if err != nil {
		return err
	}
	model_userInfo.TotalFavorited = "0"
	err = dao.NewUserInfoDAO().QueryUserInfoById(int64(u.userId), &db_userInfo)
	if err != nil {
		return err
	}
	model_userInfo.User = db_userInfo

	//fmt.Println(model_userInfo)

	//u.comment.User = model_userInfo
	OneCommentAction.User_ = model_userInfo

	//fmt.Println(OneCommentAction)
	if u.actionType == 1 { //用户填写评论内容
		if err := dao.NewCommentDAO().InsertCommentByUserIdAndVideoIdAndCommentText(u.userId, u.videoId, u.commentText); err != nil {
			return err
		}
		var comment db.Comment
		if err := dao.NewCommentDAO().QueryCommentId(&comment); err != nil {
			return err
		}
		OneCommentAction.Comment.ID = comment.ID

	} else { //用户删除评论内容
		//先判断这个CommentId在不在
		if ok := dao.NewCommentDAO().IsExistsCommentId(u.commentId); !ok {
			return errors.New("CommentId 不存在")
		}
		if err := dao.NewCommentDAO().DeleteCommentByCommentId(u.commentId); err != nil {
			return err
		}
		OneCommentAction.Comment.ID = u.commentId
	}
	OneCommentAction.CommentText = u.commentText
	u.comment = &OneCommentAction
	fmt.Println(u.comment)
	return nil
}

func (u *commentAction) packData() error {
	//_ = util.FillCommentFields(u.comment)
	u.Response = &Response{MyComment: u.comment}
	return nil
}
