package test

import (
	"encoding/json"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.Init()
	fmt.Println(model.CreatUser(testUserA, "password111"))
}

func TestSearchUserByName(t *testing.T) {
	db.Init()
	fmt.Println(model.SearchUserByName(testUserA))
}
func TestSearchUserById(t *testing.T) {
	db.Init()
	user, _ := model.SearchUserByID(1)
	fmt.Println(user)
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
