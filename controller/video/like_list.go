package video

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/service/video"
	"net/http"
)

type FavorVideoListResponse struct {
	db.CommonResponse
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
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyFavorVideoListController) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: db.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListController) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{CommonResponse: db.CommonResponse{StatusCode: 0},
		FavorList: favorList,
	})
}
