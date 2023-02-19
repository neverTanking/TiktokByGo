package router

import (
	"JWT"

	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/db"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	// public directory is used to serve static resources
	db.Init()
	r.Static("/static", "./static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", JWT.WTMiddleware(), controller.Feed)
	apiRouter.GET("/user/", JWT.WTMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register) //
	apiRouter.POST("/user/login/", controller.Login)       //
	apiRouter.POST("/publish/action/", JWT.WTMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", JWT.WTMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", JWT.WTMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", JWT.WTMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", JWT.WTMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", JWT.WTMiddleware(), controller.CommentList)

	// extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction)
	// apiRouter.GET("/relation/follow/list/", controller.FollowList)
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	// apiRouter.GET("/relation/friend/list/", controller.FriendList)
	// apiRouter.GET("/message/chat/", controller.MessageChat)
	// apiRouter.POST("/message/action/", controller.MessageAction)
	return r
}
