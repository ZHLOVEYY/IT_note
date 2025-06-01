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
- `docker exec -it fortest-mysql-1 mysql -u root -p` 
	- 然后输入密码，检查数据库  `use posts_db`  , `select * from posts`
