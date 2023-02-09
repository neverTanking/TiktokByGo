package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/neverTanking/TiktokByGo/config"
)

var Rdb *redis.Client
var ctx = context.Background()
var LIKE = "like"

func InitClientRdb() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := Rdb.Ping(ctx).Err(); err != nil {
		return err
	} else {
		return nil
	}
}

// 点赞，更新喜欢列表，喜欢列表是所有点赞的视频
func UpdatePostLike(user_id int64, vedio_id int64, state bool) error {
	str_key := fmt.Sprintf("%s-%d", LIKE, user_id)
	if state {
		if err := Rdb.SAdd(ctx, str_key, vedio_id).Err(); err != nil {
			fmt.Println("UpdatePostLike SAdd Failed:", err)
			return err
		}
	} else {
		if err := Rdb.SRem(ctx, str_key, vedio_id).Err(); err != nil {
			fmt.Println("UpdatePostLike SRem Failed:", err)
			return err
		}
	}
	return nil
}

// 查询用户对该视频是否点赞
func GetLikeState(user_id int64, vedio_id int64) (bool, error) {
	str_key := fmt.Sprintf("%s-%d", LIKE, user_id)
	ok, err := Rdb.SIsMember(ctx, str_key, vedio_id).Result()
	if err != nil {
		fmt.Println("RedisContains SIsMember Failed:", err)
		return false, err
	}
	return ok, nil
}
