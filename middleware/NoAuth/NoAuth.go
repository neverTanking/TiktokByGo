package NoAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
	"net/http"
	"strconv"
)

func NoAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		strUserId := c.Query("user_id")
		if strUserId == "" {
			strUserId = c.PostForm("user_id")
		}
		//用户不存在
		if strUserId == "" {
			c.JSON(http.StatusOK, db.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "请求user_id失败",
			})
			c.Abort()
		}
		userId, err := strconv.ParseInt(strUserId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, db.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "请求user_id失败",
			})
			c.Abort()
		}
		c.Set("UserId", userId)
		c.Next()
	}
}
