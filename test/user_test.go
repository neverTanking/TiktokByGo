package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testUserAA = "douyinTestUserA"

func TestMain(m *testing.M) {
	db.Init()
	model.CreatUser(testUserAA, "111")
	user := db.User{}
	db.DB.Where("name = ?", testUserAA).Find(&user)
	code := m.Run()
	db.DB.Unscoped().Delete(&user)
	os.Exit(code)
}

// Will cause id increase
func TestCreateUser(t *testing.T) {
	userId, err := model.CreatUser(testUserAA, "password11")
	if err != nil {
		fmt.Errorf("CreateUser fail:%v", err)
	}
	user := db.User{}
	res := db.DB.Find(&user, userId)
	assert.Equal(t, int64(1), res.RowsAffected, "user table should have only one line meet the criteria")
	assert.Equal(t, userId, user.ID, "user id should equal")
	assert.Equal(t, testUserAA, user.Name, "user name should equal")
	assert.Equal(t, "password11", user.Password, "user name should equal")
	res = db.DB.Unscoped().Delete(&user, userId)
	assert.Equal(t, int64(1), res.RowsAffected, "only one line should be deleted from database")
}

func TestSearchUserByName(t *testing.T) {
	//Success case
	funcUser, err := model.SearchUserByName(testUserAA)
	if err != nil {
		errors.New("user not found")
	}
	realUser := db.User{}
	//直接查询数据库方式验证
	res := db.DB.Find(&realUser, funcUser)
	assert.Equal(t, int64(1), res.RowsAffected, "One user data found")
	assert.Equal(t, funcUser.ID, realUser.ID, "user id should equal")
	assert.Equal(t, testUserAA, realUser.Name, "user name should equal")
}

func TestSearchUserByID(t *testing.T) {
	funcUser, err := model.SearchUserByID(1)
	if err != nil {
		errors.New("user not found")
	}
	realUser := db.User{}
	//直接查询数据库方式验证
	res := db.DB.Find(&realUser, funcUser.ID)
	assert.Equal(t, int64(1), res.RowsAffected, "One user data found")
	assert.Equal(t, funcUser.ID, realUser.ID, "user id should equal")
	assert.Equal(t, "duplicateUsername", realUser.Name, "user name should equal")
	//以json 格式打印user
	b, err := json.Marshal(funcUser)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func TestVerify(t *testing.T) {
	var errNotFound = errors.New("user not found")
	var errWrongPassword = errors.New("wrong password")

	existUser, _ := model.SearchUserByName(testUserAA)
	//existUser(id) with correct password
	var err error
	err = model.SearchUserForVerify(existUser.ID, "111")
	assert.Empty(t, err, "password check pass")

	//existUser(id) with incorrect password
	err = model.SearchUserForVerify(existUser.ID, "wrong password")
	assert.NotEmpty(t, err, "err should work")
	assert.Equal(t, err, errWrongPassword, "should tell me wrong password")

	//not exist user with any password
	err = model.SearchUserForVerify(99999, "any password")
	assert.NotEmpty(t, err, "err should work")
	assert.Equal(t, err, errNotFound, "should tell me not found user")

}
