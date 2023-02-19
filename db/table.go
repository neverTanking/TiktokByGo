package db

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID      uint
	User        User
	VideoID     uint
	CommentText string
}
type Like struct {
	gorm.Model
	UserID  uint
	VideoID uint
}
type User struct {
	gorm.Model
	Name            string `json:"name"`
	Password        string `json:"password,omitempty"`
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介

	Videos []Video
	Likes  []Like
}

type Video struct {
	gorm.Model
	UserID   uint
	PlayUrl  string
	CoverUrl string
	Title    string

	Likes    []Like
	Comments []Comment
}
