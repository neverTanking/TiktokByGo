package Redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/neverTanking/TiktokByGo/config"
	"sync"
	"time"
)

var (
	CommentNum    = "commentNum"
	LIKENUM       = "likeNum"
	LIKE          = "like"
	FavoriteCount = "favoriteCount"
	FollowCount   = "followCount"
	FollowerCount = "followerCount"
	WorkCount     = "workCount"
)

var (
	Rdb               *redis.Client
	ctx               = context.Background()
	defaultExpireTime = 2 * time.Hour
)

type RedisDao struct {
}

var ExecRedis bool

var (
	redisDao  *RedisDao
	redisOnce sync.Once
)

func NewRedisDao() *RedisDao {
	redisOnce.Do(func() {
		ExecRedis = true
		redisDao = new(RedisDao)
		InitClientRdb()
	})
	return redisDao
}

func InitClientRdb() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := Rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}
