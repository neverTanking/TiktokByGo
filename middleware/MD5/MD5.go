package MD5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 不区分大小写判断两个字符串是否相等
// content是数据库里的，encrypted是新生成MD5的
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}

func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
