package video

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/controller"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/service/video"
	"net/http"
	"strconv"
)

type LikeController struct {
	*gin.Context
	UserId     uint
	VideoId    uint
	actionType int
}

func LikeActionController(c *gin.Context) {
	NewLikeController(c).Finish()
}

func NewLikeController(c *gin.Context) *LikeController {
	return &LikeController{Context: c}
}

func (u *LikeController) Finish() {
	//Get Parameter
	if err := u.ParseParameter(); err != nil {
		u.ReturnError(err.Error()) //自定义的ErrorString,将err转化成string
		return
	}
	//use Parameter
	if err := video.LikeAction(u.UserId, u.VideoId, u.actionType); err != nil {
		u.ReturnError(err.Error())
	}
	u.ReturnSuccess()
}

func (u *LikeController) ParseParameter() error {
	curUserId, _ := u.Get("UserId") //curUserId 是 interface类型的
	UserId, ok := curUserId.(uint)  //转uint
	if !ok {
		return errors.New("ParseUserId Failed") //创建错误
	}
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
	u.VideoId = uint(VideoId)
	u.UserId = UserId
	return nil
}

func (u *LikeController) ReturnError(message string) {
	u.JSON(http.StatusOK, db.CommonResponse{
		StatusCode: controller.Fail,
		StatusMsg:  message,
	})
}
func (u *LikeController) ReturnSuccess() {
	u.JSON(http.StatusOK, db.CommonResponse{
		StatusCode: controller.Success,
	})
}
