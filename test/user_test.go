package test

import (
	"encoding/json"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testUserAA = "douyinTestUserA"
var testUserBB = "douyinTestUserB"

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	userId, err := model.CreatUser(testUserAA, "password11")
	if err != nil {
	}
	user := db.User{}
	res := db.DB.Find(&user, userId)
	assert.Equal(t, int64(1), res.RowsAffected, "One user data found")
	assert.Equal(t, userId, user.ID, "user id should equal")
	assert.Equal(t, testUserAA, user.Name, "user name should equal")
	assert.Equal(t, "password11", user.Password, "user name should equal")
}

func TestSearchUserByName(t *testing.T) {
	userId, ok := model.SearchUserByName(testUserAA)
	if ok == false {

	}
	user := db.User{}
	res := db.DB.Find(&user, userId)
	assert.Equal(t, int64(1), res.RowsAffected, "One user data found")
	assert.Equal(t, userId, user.ID, "user id should equal")
	assert.Equal(t, testUserAA, user.Name, "user name should equal")
	assert.Equal(t, "password111", user.Password, "user name should equal")
}

func TestSearchUserById(t *testing.T) {
	user, _ := model.SearchUserByID(1)
	//以json 格式打印user
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
