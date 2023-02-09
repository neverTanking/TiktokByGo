# TiktokByGo

抖音简版——大作业

## 使用

### 配置数据库

自行在 db 文件夹下创建本地文件配置 config.go，作为配置 mysql 的文件。

```LUA
| - db
    | - config.go
```

```go
package db

var Database = map[string]string{
	"user":     "root",   // mysql 用户名
	"password": "",       // mysql 密码
	"db_name":  "tiktok", // mysql 数据库名
}

```
