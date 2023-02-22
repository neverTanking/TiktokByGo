package comment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/service/comment"
	"net/http"
	"strconv"
)

type commentResponse struct {
	db.CommonResponse
	*comment.Comment
}
type commentController struct {
	*gin.Context
	userId      uint
	videoId     uint
	actionType  int
	commentText string
	commentId   int
}

func commentActionController(c *gin.Context) {
	NewCommentController(c).Finish()
}

func NewCommentController(c *gin.Context) *commentController {
	return &commentController{Context: c}
}

func (u *commentController) Finish() {
	//Get parameter
	if err := u.ParseParameter(); err != nil {
		return
	}
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
	strCommentText := u.Query("comment_text")
	strCommentId := u.Query("comment_id")
	commentId, err := strconv.ParseInt(strCommentId, 10, 64)
	if err != nil {
		return errors.New("ParseCommentId Failed")
	}
	u.commentText = strCommentText
	u.commentId = int(commentId)
	return nil
}

func (u *commentController) ReturnError(msg string) {
	u.JSON(http.StatusOK, commentResponse{
		CommonResponse: db.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (u *commentController) ReturnSuccess() {
	u.JSON(http.StatusOK, commentResponse{
		CommonResponse: db.CommonResponse{StatusCode: 0, StatusMsg: "success"},
	})
}
