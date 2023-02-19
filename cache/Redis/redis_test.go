package Redis

import (
	"fmt"
	"testing"
)

func TestInitClientRdb(t *testing.T) {
	if err := InitClientRdb(); err != nil {
		fmt.Println("InitClientRdb Failed:", err)
		return
	}
	fmt.Println("InitClientRdb Successful")
}

func TestUpdatePostLike(t *testing.T) {
	TestInitClientRdb(t)
	user_id := int64(1)
	vedio_id := int64(10)
	state := true
	if err := UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike Failed:", err)
		return
	}
	//再关注一次看看会不会报错
	if err := UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike1 Failed:", err)
		return
	}

	user_id = int64(1)
	vedio_id = int64(2)
	state = false
	if err := UpdatePostLike(user_id, vedio_id, state); err != nil {
		fmt.Println("UpdatePostLike2 Failed:", err)
		return
	}

	fmt.Println("UpdatePostLike Successful")
}

func TestGetLikeState(t *testing.T) {
	TestInitClientRdb(t)
	user_id := int64(1)
	vedio_id := int64(10)
	state, err := GetLikeState(user_id, vedio_id)
	if err != nil {
		fmt.Println("GetLikeState Failed:", err)
		return
	}
	fmt.Println(state)

	user_id = int64(1)
	vedio_id = int64(100)
	state, err = GetLikeState(user_id, vedio_id)
	if err != nil {
		fmt.Println("GetLikeState Failed:", err)
		return
	}
	fmt.Println(state)
}
