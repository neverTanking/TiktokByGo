# TiktokByGo

抖音简版——大作业

## 文档
[NeverTanking青训营后端结业项目答辩汇报文档](https://lgmz502nro.feishu.cn/docx/S9CldZF8MousnmxVSGact5CanCg)

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
