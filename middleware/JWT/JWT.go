package JWT

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/config"
	"github.com/neverTanking/TiktokByGo/db"
	"net/http"
	"strconv"
	"time"
)

var jwtKey = []byte("uuuuuu")

var cnt = 0

type MyClaims struct {
	UserId   uint
	UserName string
	PassWord string
	jwt.StandardClaims
}

// GetToken从username和password变成Token

func GetToken(userid uint, username string, password string) (string, error) {
	cnt++
	//创建一个我们自己的声明
	myClaims := &MyClaims{
		UserId:   userid,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,               //什么时间生效
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), //只有2小时的有效时间
			Issuer:    "neverTaking",                        //签发者
			Subject:   "TiktokByGo",                         //主题
			Id:        strconv.Itoa(cnt),                    //编号
		},
	}
	//NewWithClaims使用HS256指定签名和声明生成一个令牌token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	//SignedString创建并返回一个完整有符号的JWT
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("GetToken Failed:", err)
		return "", err
	}
	return tokenStr, nil

}

// TokenToClaim 解析Token变成Claim
func TokenToClaim(tokenStr string) (*MyClaims, bool) {
	token, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key1, ok := token.Claims.(*MyClaims); ok {
			if token.Valid {
				return key1, true
			} else {
				return key1, false
			}
		}
	}
	return nil, false
}

// 下面这个要用apifox测

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token") //请求url里的token参数值
		if tokenStr == "" {
			tokenStr = c.PostForm("token") //请求表单里的token参数值
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, db.CommonResponse{
				config.USERNOTFOUND,
				"用户不存在",
			})
			c.Abort()
			return
		}
		//验证token是否正确
		claim, ok := TokenToClaim(tokenStr)
		if !ok {
			//错误403有很多种情况，最常见的是无权限访问
			c.JSON(http.StatusOK, db.CommonResponse{
				config.TOKENNOTRIGHT,
				"token不正确",
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			c.JSON(http.StatusOK, db.CommonResponse{
				config.TOKENOUT,
				"token过期",
			})
			c.Abort()
			//Abort不再执行后面的Handeler
			return
		}
		if time.Now().Unix() < claim.NotBefore {
			c.JSON(http.StatusOK, db.CommonResponse{
				config.TOKENOUT,
				"token还没生效",
			})
			c.Abort()
			return
		}
		//apifox测试用，后面补充完，要删掉，因为后面还有response
		{
			c.JSON(http.StatusOK, db.CommonResponse{
				config.SUCCESS,
				"token正确",
			})
		}
		c.Set("UserId", claim.UserId)
		c.Set("UserName", claim.UserName)
		c.Set("UserPassword", claim.PassWord)
		c.Next()
		//Next继续执行后面的Handeler
	}
}
