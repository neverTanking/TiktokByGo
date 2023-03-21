package comment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/neverTanking/TiktokByGo/service/comment"
	"net/http"
	"strconv"
)

type commentResponse struct {
	model.Response
	*comment.Response1
}
type commentController struct {
	*gin.Context
	userId      uint
	videoId     uint
	actionType  int
	commentText string
	commentId   uint
}

func CommentActionController(c *gin.Context) {
	NewCommentController(c).Finish()
}

func NewCommentController(c *gin.Context) *commentController {
	return &commentController{Context: c}
}

func (u *commentController) Finish() {
	//Get parameter
	if err := u.ParseParameter(); err != nil {
		u.ReturnError(err.Error())
		return
	}
	commentRes, err := comment.QueryCommentAction(u.userId, u.videoId, u.actionType, u.commentText, u.commentId)
	//fmt.Println(commentAction)
	if err != nil {
		u.ReturnError(err.Error())
		return
	}
	u.ReturnSuccess(commentRes)
}

func (u *commentController) ParseParameter() error {
	rawUserId, _ := u.Get("UserId")
	UserId, ok := rawUserId.(uint)
	if !ok {
		return errors.New("ParseUserId Failed") //创建错误
	}
	u.userId = UserId
	//获取VideoId
	strVideoId := u.Query("video_id")
	VideoId, err := strconv.ParseInt(strVideoId, 10, 32)
	if err != nil {
		return errors.New("ParseVideo Failed")
	}
	//获取actionType
	strActionType := u.Query("action_type")
	ActionType, err := strconv.ParseInt(strActionType, 10, 32)
	if err != nil {
		return errors.New("ParseActionType Failed")
	}
	u.actionType = int(ActionType)
	u.videoId = uint(VideoId)
	if u.actionType == 1 {
		strCommentText := u.Query("comment_text")
		u.commentText = strCommentText
		u.commentId = 0 //默认不存在
	} else if u.actionType == 2 {
		strCommentId := u.Query("comment_id")
		commentId, err := strconv.ParseInt(strCommentId, 10, 64)
		if err != nil {
			return errors.New("ParseCommentId Failed")
		}
		u.commentId = uint(commentId)
		u.commentText = "" //默认不存在
	} else {
		return errors.New("action_type invalid")
	}

	return nil
}

func (u *commentController) ReturnError(msg string) {
	u.JSON(http.StatusOK, commentResponse{
		Response: model.Response{StatusCode: 1, StatusMsg: msg}})
}

func (u *commentController) ReturnSuccess(comment *comment.Response1) {
	u.JSON(http.StatusOK, commentResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		Response1: comment,
	})
}
