GO 使用 Gin + Gorm + ES 简单示例

## 目录
``` text
post-platform/
├── main.go           # 主程序（Gin 服务）
├── db/
│   └── db.go         # MySQL 连接
├── models/
│   └── post.go       # 帖子模型
├── elasticsearch/
│   └── es.go         # Elasticsearch 操作
├── go.mod
└── go.sum
```

## 准备
- docker
- docker-compose
gin，gorm，es 相关 go 包

## 运行
`docker-compose up --build`

## 测试
 - http://localhost:8080/posts   POST
	发送
``` json
{
    "title": "My first post",
    "content": "我爱死go 了!"
}
```
	多发送几次，修改不同内容
在 sql 的 docker 操作界面 exec 中，mysql -u root -p  登录 sql，查询结果是否真实存入

- http://localhost:8080/search?q=go  GET
	q 代表查询的字段内容
	可以查看到返回的结果

