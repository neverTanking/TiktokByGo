package comment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/service/comment"
	"net/http"
	"strconv"
)

type ListResponse struct {
	db.CommonResponse
	*comment.List
}

type CommentListController struct {
	*gin.Context

	videoId uint
	userId  uint
}

func NewCommentListController(c *gin.Context) *CommentListController {
	return &CommentListController{Context: c}
}

func QueryCommentListController(c *gin.Context) {
	NewCommentListController(c).Finish()
}

func (u *CommentListController) Finish() {
	if err := u.checkNum(); err != nil {
		u.ReturnError(err.Error())
		return
	}
	commentlist, err := comment.QueryCommentList(u.userId, u.videoId)
	if err != nil {
		u.ReturnError(err.Error())
		return
	}
	u.ReturnSuccess(commentlist)

}

func (u *CommentListController) checkNum() error {
	strUserId, _ := u.Get("UserId")
	userId, ok := strUserId.(uint)
	if !ok {
		return errors.New("UserId解析错误")
	}
	u.userId = userId
	strVideoId := u.Query("video_id")
	videoId, err := strconv.ParseInt(strVideoId, 10, 64)
	if err != nil {
		return err
	}
	u.videoId = uint(videoId)
	return nil
}

func (u *CommentListController) ReturnError(msg string) {
	u.JSON(http.StatusOK, db.CommonResponse{StatusCode: 1, StatusMsg: msg})
}
func (u *CommentListController) ReturnSuccess(commentList *comment.List) {
	u.JSON(http.StatusOK, ListResponse{
		CommonResponse: db.CommonResponse{StatusCode: 0, StatusMsg: "Success"},
		List:           commentList,
	})
}
