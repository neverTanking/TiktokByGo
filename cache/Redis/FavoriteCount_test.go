package Redis

import (
	"fmt"
	"testing"
)

func TestNewRedisDao_GetUserFavoriteCount(t *testing.T) {
	userId := uint(121)
	cnt := int64(10)
	if err := NewRedisDao().SetUserFavoriteCount(userId, cnt); err != nil {
		panic(err)
	}
	nowCnt, err := NewRedisDao().GetUserFavoriteCount(userId)
	if err != nil {
		panic(err)
	}
	if nowCnt != cnt {
		panic("GetUserFavoriteCount Has Problem")
	}
	fmt.Println("GetUserFavoriteCount Success")
}

func TestRedisDao_SetUserFavoriteCount(t *testing.T) {
	userId := uint(780)
	cnt := int64(5)
	//先设置key的值
	if err := NewRedisDao().SetUserFavoriteCount(userId, cnt); err != nil {
		panic(err)
	}
	//查询key的值
	nowCnt, err := NewRedisDao().GetUserFavoriteCount(userId)
	if err != nil {
		panic(err)
	}
	if nowCnt != cnt {
		panic("SetUserFavoriteCount Has Problem")
	}
	fmt.Println("Set UserFavoriteCount Success")
}

func TestRedisDao_ExistUserFavoriteCount(t *testing.T) {
	userId := uint(10086)

	//设置这个key
	err := NewRedisDao().SetUserFavoriteCount(userId, 10)
	//测试函数ExistUserFavoriteCount，返回值应该是yes
	ok, err := NewRedisDao().ExistUserFavoriteCount(uint(userId))
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("ExistUserFavoriteCount Has Problem1")
	}

	//把key删了
	strKey := fmt.Sprintf("%s-%d", FavoriteCount, userId)
	if err := DelOneKeyByKey(strKey); err != nil {
		fmt.Println(err)
	}

	//测试函数ExistUserFavoriteCount，返回值应该是no
	ok, err = NewRedisDao().ExistUserWorkCount(uint(userId))
	if err != nil {
		fmt.Println(err)
	}
	if ok {
		panic("ExistUserFavoriteCount Has Problem2")
	}
	fmt.Println("ExistUserFavoriteCount is Right")
}

func DelOneKeyByKey(strKey string) error {
	if err := Rdb.Del(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

func TestNewRedisDao_AddOneUserFavoriteCount(t *testing.T) {
	userId := uint(678)
	cnt := int64(45)
	if err := NewRedisDao().SetUserFavoriteCount(userId, cnt); err != nil {
		panic(err)
	}
	if err := NewRedisDao().AddOneUserFavoriteCount(userId); err != nil {
		panic(err)
	}
	nowCnt, err := NewRedisDao().GetUserFavoriteCount(userId)
	if err != nil {
		panic(err)
	}
	if nowCnt != cnt+1 {
		panic("AddOneUserFavoriteCount Has Problem")
	}
	fmt.Println("AddOneUserFavoriteCount Success")
}

func TestNewRedisDao_SubOneUserFavoriteCount(t *testing.T) {
	userId := uint(999)
	cnt := int64(78)
	if err := NewRedisDao().SetUserFavoriteCount(userId, cnt); err != nil {
		panic(err)
	}
	if err := NewRedisDao().SubOneUserFavoriteCount(userId); err != nil {
		panic(err)
	}
	nowCnt, err := NewRedisDao().GetUserFavoriteCount(userId)
	if err != nil {
		panic(err)
	}
	if nowCnt != cnt-1 {
		panic("SubOneUserFavoriteCount Has Problem")
	}
	fmt.Println("SubOneUserFavoriteCount Success")
}
