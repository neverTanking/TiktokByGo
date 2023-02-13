package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/model"
)

// token -> user_id
var UsersLoginInfo = map[string]int64{
	"zhanglei---douyin": 1,
}

type UserLoginResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

/**
 * @description: 创建用户 token
 * @param {string} username 用户名
 * @param {string} password 用户密码
 * @return {string} 生成的 token
 */
func createToken(username string, password string) string {
	// TODO 加密
	return strings.Join([]string{username, password}, "---")
}

func decodeToken(token string) ([]string, error) {
	// TODO 解密
	res := strings.Split(token, "---")
	return res, nil
}

/**
 * @description: 获取用户信息
 * @param {uint} id	用户id
 * @return {*}	用户信息
 */
func getUserInfo(id uint) User {
	user := model.SearchUser(id)
	return user
}

/**
 * @description: 创建新用户，并存储
 * @param {uint} id 用户id
 * @param {string} name 用户名
 * @param {string} password 用户密码
 */
func newUser(name string, password string) bool {
	return model.CreatUser(name, password)
}

func searchUser(username string) (User, bool) {
	for _, user := range usersDate {
		if user.Name == username {
			return user, true
		}
	}
	return User{}, false
}

// 验证密码
func verifyPassword(input string, password string) bool {
	return input == password
}

/**
 * @description: 验证用户登录中间件
 * @param {string} token 请求参数
 */
func AuthMiddleware(c *gin.Context) {
	// 获取token
	token := c.Query("token")
	// 验证token
	if _, exist := usersLoginInfo[token]; exist {
		c.Next()
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: NotLogin},
		})
		c.Abort()
	}
}

/**
 * @description: 用户登录处理函数，返回用户登录信息
 */
func Register(c *gin.Context) {
	// 1. http请求中获取用户信息
	username := c.Query("username")
	password := c.Query("password")

	// 2. 查询用户是否存在，并返回用户信息
	user, exist := searchUser(username)
	// 3.1 存在则验证密码
	if exist {
		user.FollowCount = 0
		// verifyPassword(password, user.Password)
	}
	// 3.2 不存在则创建用户，并返回用户信息
	model.CreatUser(username, password)
	token := createToken(username, password)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success, StatusMsg: SignUpOk},
		UserId:   id,
		Token:    token,
	})

	// token := createToken(username, password)

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: Fail, StatusMsg: Existed},
	// 	})
	// } else {
	// 	// newUser := User{
	// 	// 	Id:   calcUserId(&userIdSequence),
	// 	// 	Name: username,
	// 	// }
	// 	newUser(calcUserId(&userIdSequence), username, password)

	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: Success},
	// 		UserId:   userIdSequence,
	// 		Token:    token,
	// 	})
	// }
}

/**
 * @description: 用户登录处理函数，返回用户登录信息
 */
func Login(c *gin.Context) {
	// 1. http请求中获取用户信息
	username := c.Query("username")
	password := c.Query("password")

	// 2. 查询用户是否存在，并返回用户信息
	user, exist := searchUser(username)
	if !exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: Fail, StatusMsg: NotExisted},
		})
		return
	}
	// 3. 验证密码
	// if !verifyPassword(password, user.Password) {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: Fail, StatusMsg: WrongPassword},
	// 	})
	// 	return
	// }

	// 4. 生成token，存储在服务器
	token := createToken(username, password)
	usersLoginInfo[token] = user.Id // 服务器存储token

	// 5. 返回用户信息和token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success},
		UserId:   user.Id,
		Token:    token,
	})

	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: Success},
	// 		UserId:   user.Id,
	// 		Token:    token,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: Fail, StatusMsg: NotExisted},
	// 	})
	// }
}

/**
 * @description: 获取用户信息处理函数，返回用户信息
 */
func UserInfo(c *gin.Context) {

	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	user := getUserInfo(user_id)
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: Success},
		User:     user,
	})
	// if user, exist := usersDate[user_id.(uint)]; exist {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: Success},
	// 		User:     user,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: Fail, StatusMsg: NotExisted},
	// 	})
	// }
}
