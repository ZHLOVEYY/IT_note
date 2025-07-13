GO 使用 RabbitMQ+GIN+GORM+Docker 部署，实现帖子存储
## 目录
demo 目录
```
post-platform/
├── Dockerfile        # Go 应用 Dockerfile
├── docker-compose.yml # Docker Compose 配置
├── main.go           # 主程序（Gin + RabbitMQ 消费者）
├── db/
│   └── db.go         # 数据库连接
├── models/
│   └── post.go       # 帖子模型
├── rabbitmq/
│   └── consumer.go   # RabbitMQ 消费者逻辑
├── go.mod
└── go.sum
```

## 准备
- docker
- docker-compose
## 运行
`docker-compose up --build` 
然后在 docker 中可以检查容器情况
## 测试
#### postman
- `http://localhost:8080/ping`  GET
- `http://localhost:8080/posts` POST
``` json
	{
    "title": "My First post",
    "content": "Hello, world!"
}
```

#### 检查容器 
- `docker exec -it fortest-mysql-1(name) mysql -u root -p` 
	- 然后输入密码，检查数据库  `use posts_db`  , `select * from posts`
- 也可以在 docker客户端 的 Exec 中直接操作`mysql -u root -p`

## 对比功能
为了说明 RabbitMQ 的引入有什么好处！
- 下载wrk
`brew install wrk`
 - 测试指令：
	``` bash
	wrk -t 10 -c 100 -d 30s --latency http://localhost:8080/posts -s post.lua
	```
	-t：使用 ​**​10 个线程​**​ 并发发送请求。
	-c：总共建立并保持 ​**​100 个 TCP 连接​**​（即并发数）。
	-d 30s：测试持续 ​**​30 秒​**​。
	--latency：在测试结束后，输出 ​**​延迟统计信息**
	-s post.lua：指定 lua 脚本
- lua 脚本：
``` lua
wrk.method = "POST"
wrk.body   = '{"title":"Test Post","content":"This is a test"}'
wrk.headers["Content-Type"] = "application/json"
```
#### 没有rabbitMQ
![[../../../../attachments/Pasted image 20250601114039.png]]
#### 有rabbitMQ
![[../../../../attachments/Pasted image 20250601115515.png]]
#### wrk 测试分析

- 请求延迟
	有 RabbMQ整体效果更好，平均值低了 9倍
	有 RabbMQ 的里，50% 请求延迟：6.21ms，99% 请求延迟：74.02ms 都显著更好
- 吞吐量
	显然有 RabbitMQ 的总吞吐量更大（总的 request 和每秒的 request）
- 稳定性
	查看请求延迟百分比以及 stdev （越小越好）可以发现，有 RabbitMQ 更加稳定

补充：
- RabbitMQ 支持多个消费者并行处理消息，通过增加消费者实例（--scale app=3）可进一步提高吞吐量和降低延迟。
	- docker-compose up --scale app=3
- 同步版本受限于 MySQL 写入性能，扩展性较差（需优化数据库，如分片或读写分离）
- 停止消费者，发送 100 个请求然后查看 http://localhost:15672 可以检查数据的完整性（存在里面）

疑问：

- 问：为什么同样都是数据库读取，RabbitMQ 的话我的理解就是中间“多了一个管道”这样，但是数据库还是一个个读取的，理论上应该加了 rabbitMQ 的延迟更高么？
	答：虽然 RabbitMQ 增加了一个中间层（消息队列），但它带来的 异步处理 和 解耦 效应显著降低了 API 端的延迟
	- 没rabbitMQ 的同步版本流程：
		- 用户发送 POST /posts 请求。
		- Gin 服务接收请求，解析 JSON 数据。
		- Gin 直接调用 GORM 将帖子写入 MySQL（db.Create(&post)）。
		- 等待 MySQL 完成写入操作（包括磁盘 I/O、事务提交等）。
		- 返回响应给用户。
		延迟来自：MySQL 写入延迟：通常在 10-50ms 级别（取决于数据库负载、索引、锁等）。以及高并发下，MySQL 可能出现锁竞争或连接池瓶颈，进一步增加延迟
	- 异步版本有 RabbitMQ 流程：
		- 用户发送 POST /posts 请求。
		- Gin 服务接收请求，解析 JSON 数据。
		- Gin 将帖子序列化为 JSON，推送至 RabbitMQ 队列（ch.Publish）。
		- 立即返回响应给用户（不等待数据库写入）。
		- 消费者（后台 goroutine）从 RabbitMQ 队列读取消息，调用 GORM 写入 MySQL。
		延迟来自：推送消息到 RabbitMQ：通常在 1-2ms 级别（RabbitMQ 是内存操作，速度很快）。API 响应时间仅包含推送消息的耗时，不包括数据库写入。
	结论：RabbitMQ 并没有增加 API 请求的延迟，而是将数据库写入的延迟“转移”到消费者端
## 备忘
- docker 中的健康检查：容器启动顺序虽然 depend 确认了，但是内部的服务可能没有完全准备好，所以需要service_healthy 状态进行检查