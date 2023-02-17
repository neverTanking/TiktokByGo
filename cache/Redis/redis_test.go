package Redis

import (
	"fmt"
	"testing"
)

func TestInitClientRdb(t *testing.T) {
	NewRedisDao()
}

func TestUpdatePostLike(t *testing.T) {
	TestInitClientRdb(t)
	user_id := uint(1)
	vedio_id := uint(10)
	state := true
	if err := NewRedisDao().UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike Failed:", err)
		return
	}
	//再关注一次看看会不会报错
	if err := NewRedisDao().UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike1 Failed:", err)
		return
	}

	user_id = uint(1)
	vedio_id = uint(2)
	state = false
	if err := NewRedisDao().UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike2 Failed:", err)
		return
	}

	fmt.Println("UpdatePostLike Successful")
}

func TestGetLikeState(t *testing.T) {
	TestInitClientRdb(t)
	user_id := uint(1)
	vedio_id := uint(10)
	state, err := NewRedisDao().GetLikeState(user_id, vedio_id)
	if err != nil {
		fmt.Println("GetLikeState Failed:", err)
		return
	}
	fmt.Println(state)

	user_id = uint(1)
	vedio_id = uint(100)
	state, err = NewRedisDao().GetLikeState(user_id, vedio_id)
	if err != nil {
		fmt.Println("GetLikeState Failed:", err)
		return
	}
	fmt.Println(state)
}
