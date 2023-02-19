package test

import (
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
	fmt.Println(model.SearchUserByID(1))
}
