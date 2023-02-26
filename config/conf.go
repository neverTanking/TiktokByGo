package config

// return status_code
const (
	SUCCESS       = 0
	USERNOTFOUND  = 401
	TOKENOUT      = 402
	TOKENNOTRIGHT = 403
	SUCCESS_MSG   = "Successful"
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
