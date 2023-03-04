package config

// return status_code
const (
	SUCCESS       = 0
	USERNOTFOUND  = 401
	TOKENOUT      = 402
	TOKENNOTRIGHT = 403
)

// Redis
var (
	RedisIP       = "localhost"
	RedisPort     = "6379"
	RedisAddr     = RedisIP + ":" + RedisPort
	RedisPassword = ""
	RedisDB       = 0
)

// MySql
var ()

// minio
var (
	Miniourl       = "192.168.199.129:9000"             //搭建容器地址
	MinioaccessKey = "62N5uF7SDDD6Bjhq"                 //minioadmin
	MiniosecretKey = "2z2zqK34guD01OSd2edvm3gR6hZnhdws" //minioadmin //key
	HeartbeatTime  = 2 * 60

	BucketName   = "short_videos"
	BucketFfmpeg = "cover"
	Location     = "cn-north-1"
)

// PlayUrlPrefix 存储的图片和视频的链接
const PlayUrlPrefix = "http://192.168.199.129/short_video"
const CoverUrlPrefix = "http://192.168.199.129/cover"
