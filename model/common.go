package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type User struct {
	//<<<<<<< HEAD
	//	db.User
	//	FavoriteCount  int64  `json:"favorite_count"`  // 喜欢数
	//	FollowCount    int64  `json:"follow_count"`    // 关注总数
	//	FollowerCount  int64  `json:"follower_count"`  // 粉丝总数
	//	TotalFavorited string `json:"total_favorited"` // 获赞数量
	//	WorkCount      int64  `json:"work_count"`      // 作品数
	//	IsFollow       bool   `json:"is_follow"`
	//=======
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
}

// 单个视频
//
//	type Video struct {
//		db.Video
//		Author        User  `json:"author"` //author信息都在这里面
//		FavoriteCount int64 `json:"favorite_count"`
//		CommentCount  int64 `json:"comment_count"`
//		IsFavorite    bool  `json:"is_favorite"`
//	}
type Video struct {
	Id            uint   `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

//type Comment struct {
//	db.Comment
//	User_ User `json:"user"`
//}

type Comment struct {
	Id          uint   `json:"id"`
	User_       User   `json:"user"`
	Content     string `json:"content"`
	Create_Date string `json:"create_date"`
}
