## 作用
类似一个“中间过渡器”，应对突发流量导致数据库连接池耗尽或者请求导致服务崩溃
- **流量洪峰​**​：促销活动时，前置 Nginx 将请求写入 RabbitMQ，后端服务按能力消费
- **容灾恢复​**​：数据库故障期间，消息持久化在队列；恢复后继续消费  （消费指的是 Mysql 取出数据然后存起来）
- 将任务分发到多个消费者实例，确保高负载下任务均匀分配。这就可以实现负载均衡 （比如多个 worker 处理帖子审核）

需要考虑如果用户的申请不是很多情况下，多引入一层 RabbitMQ 其实会导致实际的速度变慢（毕竟多加了一层）


## 经典例子
GO 语言相关库：`go get -u github.com/streadway/amqp`

docker 快速部署 rabbitMQ：`docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management` 
- 5672：AMQP 端口
- 15672：管理界面端口，访问   http://localhost:15672 （ 默认用户/密码：guest/guest）

`http://localhost:15672` 利用 guest/guest 登录后可以查看队列中的内容
#### 生产者（发送端）
创建 producer文件夹下创建producer.go ，然后单独 go run（同时 go run 后面的消费者记得）
``` GO
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// 统一错误输出
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close() //关闭连接

	ch, err := conn.Channel() //建立通道，通过 conn 建立的，可以调用 amqp 中的函数
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"post_queue", // 指定创建或引用的队列名称
		false,        // 持久化  false 表示队列不会持久化到磁盘，重启 RabbitMQ 后会丢失。true 的话重启后就还在
		false,        // 自动删除   设置为 false 表示队列不会自动删除，如果 true，最后一个消费者断开后队列删除
		false,        // 独占   设置为 true 表示该队列只供一个消费者使用，当连接关闭后，队列会自动删除。false表示队列可以被多个连接使用
		false,        // 无等待  false 表示需要服务器确认队列创建，true表示客户端不会等待服务器的确认响应，如果操作失败也不会收到错误通知
		nil,          // 额外参数  额外参数可以用来设置队列的特殊属性，如消息TTL、队列最大长度、死信队列等
	)
	failOnError(err, "Failed to declare a queue")

	// 设置定时器，每5秒发送一次消息
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 创建一个函数用于发送消息，这样循环调用函数就是发送多次消息
	sendMessage := func(msgContent string) {
		err = ch.Publish(
			"",     // 交换机名称   这里是默认交换机，能够将消息直接路由到与路由键同名的队列
			q.Name, // 路由键   也就是队列名称，路由键应该与目标队列名称一致，消息才能被正确路由
			false,  // mandatory标志  false 表示消息无法路由到队列，则消息会被丢弃  如果是 true 就是当消息不能路由到队列时，RabbitMQ会返回一个Basic.Return命令给生产者
			false,  // immediate 标志   false 表示如果队列中没有消费者，消息会被存入队列等待消费， true表示当没有消费者能够立即消费该消息时，消息不会入队而是被丢弃
			amqp.Publishing{ //消息内容和性质
				ContentType: "text/plain",       //制定为 MIME 类型
				Body:        []byte(msgContent), //转换为字节类型
			})
		if err != nil {
			log.Printf("Failed to publish a message: %s", err)
			return
		}
		log.Printf(" [x] Sent %s", msgContent)
	}

	count := 1
	log.Println("Starting periodic message sending. Press Ctrl+C to exit.")

	// 等待定时器触发，定期发送消息
	for range ticker.C {
		sendMessage(fmt.Sprintf("Hello, RabbitMQ! Message #%d", count))
		count++
	}
}

```
- 这里我将函数设置为每间隔 1s 就发送消息，同时记录数据
- 如果运行后，隔一段时间再启动消费者，或者说运行中途关闭消费者，过一段时间再启动消费者，会发现中间发出的信号也会打印出来，这说明实际上是有存储在 RabbitMQ 中的（运行的时候，关闭后存储就需要看上面的设置了）


#### 消费者（接收端）
 consumer 文件夹下创建 consumer.go 然后单独一个终端 go run
``` GO
package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//连接 channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("post_queue", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // 队列
		"",     // 消费者标签
		true,   // 自动确认
		false,  // 独占
		false,  // 无本地
		false,  // 无等待
		nil,    // 额外参数
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

```




## gin框架例子

``` 
post-platform/
├── main.go           # Gin 服务（生产者）
├── rabbitmq.go       # RabbitMQ 操作
├── models/
│   └── post.go       # 帖子模型
├── db/
│   └── db.go         # 数据库连接和操作
├── consumer/
│   └── main.go       # 消费者（存储到 MySQL）
├── go.mod
└── go.sum
```

#### 服务端生产者
- 定义 post.go
``` GO

package models

type Post struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

```
gin 框架：`"go get github.com/gin-gonic/gin"`
- main.go:
``` GO

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"test/db"
	"test/rabbitmq"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 初始化数据库
	db := db.InitDB()
	// 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("post_queue", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	// 启动 RabbitMQ 消费者（在 goroutine 中运行）
	go rabbitmq.StartConsumer(db)
	// 初始化 Gin
	r := gin.Default()

	// 提交帖子接口
	r.POST("/posts", func(c *gin.Context) {
		var post struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 序列化帖子为 JSON
		postData, err := json.Marshal(post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize post"})
			return
		}

		// 发送到 RabbitMQ
		err = ch.Publish(
			"",     // 交换机
			q.Name, // 队列名称
			false,  // 强制
			false,  // 立即
			amqp.Publishing{
				ContentType: "application/json",
				Body:        postData,
			})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to RabbitMQ"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post submitted successfully"})
	})

	r.Run(":8081")
}


```

#### 数据库存储
gorm 框架，需要 go get：
``` go
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
```

- db.go
``` GO
package db

import (
	"log"
	"test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//根据情况填写password 和 dbname（具体的数据库和密码），这里用的本地 sql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移，创建 posts 表
	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

```
#### 客户端消费者
- consumer.go
``` GO
package rabbitmq

import (
	"encoding/json"
	"log"
	"test/model"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func StartConsumer(db *gorm.DB) {
	// 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("post_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 消费消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var post model.Post
			if err := json.Unmarshal(d.Body, &post); err != nil {
				log.Printf("Failed to unmarshal post: %v", err)
				continue
			}

			// 存储到数据库
			if err := db.Create(&post).Error; err != nil {
				log.Printf("Failed to save post to database: %v", err)
				continue
			}

			log.Printf("Saved post: Title=%s, Content=%s", post.Title, post.Content)
		}
	}()

	log.Printf(" [*] Waiting for posts. To exit press CTRL+C")
	<-forever
}



```
gorm 中的 Create 是只要结构体的名字一样就会找对应的表，所以结构体命名为 Post/Posts都可以，虽然和 model 中的不一样，但是如果名字不一样，Create 函数就“找不到”
#### 访问测试
分别终端运行程序后：
地址：http://localhost:8081/posts
发送内容：
``` json
{
    "title": "My First Post",
    "content": "Hello, world!"
}
```

可以发现能正确送达，同时能存储到数据库中


