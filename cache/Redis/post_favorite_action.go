package Redis

import "fmt"

// 点赞，更新喜欢列表，喜欢列表是所有点赞的视频
func (u *RedisDao) UpdatePostLike(userId uint, vedioId uint, state bool) error {
	strKey := fmt.Sprintf("%s-%d", LIKE, userId)
	if state {
		if err := Rdb.SAdd(ctx, strKey, vedioId).Err(); err != nil {
			fmt.Println("UpdatePostLike SAdd Failed:", err)
			return err
		}
	} else {
		if err := Rdb.SRem(ctx, strKey, vedioId).Err(); err != nil {
			fmt.Println("UpdatePostLike SRem Failed:", err)
			return err
		}
	}
	return nil
}

// 查询用户对该视频是否点赞
func (u *RedisDao) GetLikeState(userId uint, vedioId uint) (bool, error) {
	strKey := fmt.Sprintf("%s-%d", LIKE, userId)
	ok, err := Rdb.SIsMember(ctx, strKey, vedioId).Result()
	if err != nil {
		fmt.Println("RedisContains SIsMember Failed:", err)
		return false, err
	}
	return ok, nil
}
