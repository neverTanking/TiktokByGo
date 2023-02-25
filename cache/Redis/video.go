package Redis

import (
	"fmt"
	"strconv"
)

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
func (u *RedisDao) GetLikeState(userId uint, videoId uint) (bool, error) {
	strKey := fmt.Sprintf("%s-%d", LIKE, userId)
	ok, err := Rdb.SIsMember(ctx, strKey, videoId).Result()
	if err != nil {
		fmt.Println("RedisContains SIsMember Failed:", err)
		return false, err
	}
	return ok, nil
}

// 根据VideoId查询该视频被点赞的次数
func (u *RedisDao) GetLikeNumByVideoId(videoId uint) (int, error) {
	strKey := fmt.Sprintf("%s-%d", LIKENUM, videoId)
	strRes, err := Rdb.Get(ctx, strKey).Result()
	if err != nil {
		return 0, err
	}
	res, _ := strconv.ParseInt(strRes, 10, 64)
	return int(res), nil
}

// 给VideoId增加一个赞
func (u *RedisDao) AddOneLikeNumByVideoId(videoId uint) error {
	strKey := fmt.Sprintf("%s-%d", LIKENUM, videoId)
	if err := Rdb.Incr(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

// 给VideoId减少一个赞
func (u *RedisDao) SubOneLikeNumByVideoId(videoId uint) error {
	strKey := fmt.Sprintf("%s-%d", LIKENUM, videoId)
	if err := Rdb.Decr(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

// 给VideoId增加一个评论数
func (u *RedisDao) AddOneCommentNumByVideoId(videoId uint) error {
	strKey := fmt.Sprintf("%s-%d", CommentNum, videoId)
	if err := Rdb.Incr(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

// 给VideoId减少一个评论数
func (u *RedisDao) SubOneCommentByVideoId(videoId uint) error {
	strKey := fmt.Sprintf("%s-%d", CommentNum, videoId)
	if err := Rdb.Decr(ctx, strKey).Err(); err != nil {
		return err
	}
	return nil
}

// 根据VideoId查询评论数
func (u *RedisDao) GetCommentByVideoId(videoId uint) (int, error) {
	strKey := fmt.Sprintf("%s-%d", CommentNum, videoId)
	strRes, err := Rdb.Get(ctx, strKey).Result()
	if err != nil {
		return 0, err
	}
	res, _ := strconv.ParseInt(strRes, 10, 64)
	return int(res), nil
}
