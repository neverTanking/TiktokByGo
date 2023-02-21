package Redis

import (
	"errors"
	"fmt"
	"strconv"
)

// 判断键是否存在，可能失效了
func (u *RedisDao) ExistUserFollowCount(userId uint) (bool, error) {
	strKey := fmt.Sprintf("%s-%d", FollowCount, userId)
	ok, err := Rdb.Exists(ctx, strKey).Result()
	if err != nil {
		return false, err
	}
	exists := false
	if ok == 1 {
		exists = true
	}
	return exists, nil
}

// 设定用户FollowCount的数据，直接设定
func (u *RedisDao) SetUserFollowCount(userId uint, cnt int64) error {
	strKey := fmt.Sprintf("%s-%d", FollowCount, userId)
	if err := Rdb.Set(ctx, strKey, cnt, defaultExpireTime).Err(); err != nil {
		return err
	}
	return nil
}

// 得到用户FollowCount的数据
func (u *RedisDao) GetUserFollowCount(userId uint) (int64, error) {
	strKey := fmt.Sprintf("%s-%d", FollowCount, userId)
	res, err := Rdb.Get(ctx, strKey).Result()
	if err != nil {
		return 0, err
	}
	cnt, _ := strconv.ParseInt(res, 10, 64)
	return cnt, nil
}

// 增加用户FollowCount数量一个
func (u *RedisDao) AddOneUserFollowCount(userId uint) error {
	strKey := fmt.Sprintf("%s-%d", FollowCount, userId)
	if err := Rdb.Incr(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

// 减少用户FollowCount数量一个
func (u *RedisDao) SubOneUserFollowCount(userId uint) error {
	strKey := fmt.Sprintf("%s-%d", FollowCount, userId)
	if err := Rdb.Decr(ctx, strKey).Err(); err != nil {
		return err
	}
	res, err := Rdb.Get(ctx, strKey).Result()
	if err != nil {
		return err
	}
	cnt, _ := strconv.ParseInt(res, 10, 64)
	if cnt <= -1 {
		return errors.New("The Key Not Found, You Can't Decrease it")
	}
	return nil
}
