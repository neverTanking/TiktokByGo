package video

import (
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/neverTanking/TiktokByGo/service/video"
	"net/http"
	"strconv"
)

type FavorVideoListResponse struct {
	model.Response
	*video.FavorList
}

func FavoriteListController(c *gin.Context) {
	NewProxyFavoriteList(c).Do()
}

type ProxyFavorVideoListController struct {
	*gin.Context
	userId int64
}

func NewProxyFavoriteList(c *gin.Context) *ProxyFavorVideoListController {
	return &ProxyFavorVideoListController{Context: c}
}

func (p *ProxyFavorVideoListController) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	//正式调用
	favorVideoList, err := video.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}
	//成功返回
	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListController) parseNum() error {
	//rawUserId, _ := p.Get("UserId")
	//_, ok := rawUserId.(int64)
	//if !ok {
	//	return errors.New("token中UserId解析出错")
	//}
	strBeSeenUserId := p.Query("user_id")
	BeSeenUserId, err := strconv.ParseInt(strBeSeenUserId, 10, 64)
	if err != nil {
		return err
	}
	p.userId = BeSeenUserId
	return nil
}

func (p *ProxyFavorVideoListController) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		Response: model.Response{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListController) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		FavorList: favorList,
	})
}
