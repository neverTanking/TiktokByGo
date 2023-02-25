package controller

const (
	//StatusCode
	Success = 0
	Fail    = 1

	//StatusMsg
	NotExisted    = "User not exist"             // 用户不存在
	Existed       = "User exist"                 // 用户已存在
	WrongPassword = "Wrong password"             // 登录密码错误
	SignUpOk      = "Sign up ok"                 // 注册成功
	LoginOk       = "Login ok"                   // 登录成功
	NotLogin      = "User not login"             // 用户未登录
	UnknownReason = "Unknown reason"             // 未知原因
	TokenFail     = "Token problem"              // Token异常
	Valid         = "Username or password valid" // 用户名或密码非法
)
