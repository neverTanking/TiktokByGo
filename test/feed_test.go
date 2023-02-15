package test

import (
	"github.com/neverTanking/TiktokByGo/model"
	"strconv"
	"testing"
	"time"
)

func TestGetVideoList(t *testing.T) {
	model.GetVideoList(strconv.FormatInt(time.Now().Unix(), 10), "noToken", 5)
}
