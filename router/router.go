package router

import (
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/controller"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/middleware/JWT"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	// public directory is used to serve static resources
	db.Init()
	r.Static("/static", "./static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", JWT.JWTMiddleware(), controller.Feed)
	apiRouter.GET("/user/", JWT.JWTMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register) //
	apiRouter.POST("/user/login/", controller.Login)       //
	apiRouter.POST("/publish/action/", JWT.JWTMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", JWT.JWTMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", JWT.JWTMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", JWT.JWTMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", JWT.JWTMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", JWT.JWTMiddleware(), controller.CommentList)

	// extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction)
	// apiRouter.GET("/relation/follow/list/", controller.FollowList)
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	// apiRouter.GET("/relation/friend/list/", controller.FriendList)
	// apiRouter.GET("/message/chat/", controller.MessageChat)
	// apiRouter.POST("/message/action/", controller.MessageAction)
	return r
}
