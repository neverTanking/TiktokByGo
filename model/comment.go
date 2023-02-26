package model

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
)

var fakeComment = Comment{}

func SearchCommentByID(CommentId uint, User1 User, actionType int) (comment Comment, err error) {
	var curComment db.Comment

	res := db.DB
	if actionType == 2 { //根据commentId查
		res = db.DB.First(&curComment, CommentId)
	} else { //插入，查最后一个
		res = db.DB.Last(&curComment)
	}

	if res.Error != nil {
		return fakeComment, fmt.Errorf("database err:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return fakeComment, errNotFound
	}
	return Comment{
		Id:          curComment.ID,
		User_:       User1,
		Content:     curComment.CommentText,
		Create_Date: curComment.CreatedAt.Format("1-2"),
	}, nil
}

func FillCommentFields(comment *Comment) error {
	if comment == nil {
		return errors.New("FillCommentFields comment 为空")
	}
	return nil
}
