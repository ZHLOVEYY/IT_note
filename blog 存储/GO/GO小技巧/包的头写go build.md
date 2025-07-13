更多个人笔记见[github个人笔记仓库](https://github.com/ZHLOVEYY/IT_note)
## 问题描述
在包package 的前面写//go build可以方便在执行go build的时候指定特定的环境
这点主要应用在k8s打包生成镜像的时候

比如，执行镜像打包的命令：`GOOS=linux GOARCH=arm go build -o myproject .` 
但是如果在主程序中，我们的端口需要经常切换，比如redis，sql对应的端口，本地测试和镜像测试等使用的端口不会都是3306，6379等，如下两部分：

``` Go
db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3309"))
//gorm的连接
......
redisClient := redis.NewClient(&redis.Options{  
    Addr: "localhost:6379",  
})
//redis的创建
```
如果要测试docker环境的端口地址比如sql为"root:root@tcp(project-live-sql:3309)/mydata"，就要频繁更改
那么我们就需要// go build 将环境分类！

## 解决方法
思路如下：（注意看代码一开始）
比如，我们可以新建一个config文件夹，包含：
- types.go
``` GO
package config

type config struct {
	DB    DBConfig
	Redis RedisConfig
}

type RedisConfig struct {
	Addr string
}

type DBConfig struct {
	DSN string
}
```
这里定义DB和redis对应的地址存储结构体

- local.go
``` GO
//go:build !k8s

package config
// 代表本地
var Config = config{
	DB: DBConfig{
		// 本地连接
		DSN: "root:root@tcp(localhost:13316)/mydata",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}

```
上面的//go:build !k8s代表非k8s环境时候本地测试的配置，不是注释o，虽然有//
那么执行打包镜像的指令时候我们就可以：`GOOS=linux GOARCH=arm go build -tags=!k8s -o myproject .`  这样来指定

- k8s.go
```go
//go:build k8s  
  
package config  
  
var Config = config{  
    DB: DBConfig{  
       DSN: "root:root@tcp(project-live-sql:3309)/mydata",  
    },  
    //这里的命名就是和k8s的
    Redis: RedisConfig{  
       Addr: "preoject-live-redis:11479",  
    },  
}
```
同样我们可以：`GOOS=linux GOARCH=arm go build -tags=k8s -o myproject .`  这样来指定打包docker为k8s的环境


- 这样我们的主程序部分指令就统一变成：
``` Go
db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
//第一个config是包名
redisClient := redis.NewClient(&redis.Options{  
    Addr: config.Config.Redis.Addr,  
})
```

就不用频繁改主程序的代码啦～ 只用在执行打包镜像的时候通过 -tags指定对应 go build方式啦

更多不同端口的配置在config包（文件夹下）完成就可以了