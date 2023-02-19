package JWT

// CommonResponse copy from db/common.go
type CommonResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"` //为空时序列化不会赋值
}
