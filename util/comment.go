package util

import (
	"errors"
	"github.com/neverTanking/TiktokByGo/model"
)

func FillCommentFields(comment *model.Comment) error {
	if comment == nil {
		return errors.New("FillCommentFields comment 为空")
	}
	return nil
}
