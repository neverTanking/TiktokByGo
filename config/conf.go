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
	Miniourl       = "192.168.199.129:9000"
	MinioaccessKey = "CEETOJG1955MURS4GKRU"                     //minioadmin
	MiniosecretKey = "f1sPI0nZuept9sp2Ndqu+73vQ+30yeYxEsQ9YfHf" //minioadmin //key
	HeartbeatTime  = 2 * 60
)
