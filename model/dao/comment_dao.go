package dao

import (
	"errors"
	"fmt"
	"github.com/neverTanking/TiktokByGo/db"
	"gorm.io/gorm"
	"log"
	"sync"
)

type CommentDAO struct {
}

var (
	commentDAO  *CommentDAO
	commentOnce sync.Once
)

func NewCommentDAO() *CommentDAO {
	commentOnce.Do(func() {
		commentDAO = new(CommentDAO)
	})
	return commentDAO
}

// 根据userId、videoId、commentText插入一条评论数据
func (u *CommentDAO) InsertCommentByUserIdAndVideoIdAndCommentText(userId uint, videoId uint, commentText string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&db.Comment{UserID: userId, VideoID: videoId, CommentText: commentText}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据comment_id删除一条评论数据
func (u *CommentDAO) DeleteCommentByCommentId(commentId uint) error {
	fmt.Println("9999999", commentId)
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM `comments` WHERE `id` = ?", commentId).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据userId、videoId、commentText查询刚刚插入的一条评论数据
func (u *CommentDAO) QueryCommentId(comment *db.Comment) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Last(&comment).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *CommentDAO) IsExistsCommentId(id uint) bool {
	var commentInfo db.Comment
	if err := db.DB.Where("id=?", id).First(&commentInfo).Error; err != nil {
		log.Println(err)
	}
	if commentInfo.ID == 0 {
		return false
	}
	return true
}

func (u *CommentDAO) QueryCommentListByVideoId(videoId uint, comments *[]*db.Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId comments空指针")
	}
	if err := db.DB.Model(&db.Comment{}).Where("video_id=?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}

//videoId查询视频的评论总数
func (u *CommentDAO) QueryLenCommentByVideoId(videoId uint) (int, error) {
	var commentList *[]*db.Comment
	err := db.DB.Where("video_id=?", videoId).Find(&commentList).Error
	if err != nil || len(*commentList) == 0 {
		return 0, errors.New("该视频没有评论")
	}
	return len(*commentList), nil

}