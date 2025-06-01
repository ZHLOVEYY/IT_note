package main

import (
	"encoding/json"
	"net/http"
	"os"
	"test/db"
	"test/rabbitmq"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func main() {
	// 初始化 zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 初始化数据库
	db := db.InitDB()

	// 连接 RabbitMQ （这个用于生产者的发送）
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		log.Fatal().Msg("RABBITMQ_URL environment variable not set")
	}

	var conn *amqp.Connection
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(rabbitMQURL)
		if err == nil {
			log.Info().Msg("Successfully connected to RabbitMQ")
			break
		}
		log.Info().Msgf("Failed to connect to RabbitMQ, retrying in 5 seconds... (attempt %d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	// 检查最终连接状态
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to RabbitMQ after all retries")
		return
	}

	if conn == nil {
		log.Fatal().Msg("Connection is nil after successful connection attempt")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("post_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to declare a queue")
	}

	// 启动 RabbitMQ 消费者（在 goroutine 中）
	consumerLogger := log.With().Str("component", "consumer").Logger()
	go rabbitmq.StartConsumer(db, consumerLogger)

	// 初始化 Gin
	r := gin.Default()

	// 提交帖子接口
	apiLogger := log.With().Str("component", "api").Logger()
	r.POST("/posts", func(c *gin.Context) {
		var post struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&post); err != nil {
			apiLogger.Error().Err(err).Msg("Failed to bind JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 序列化帖子为 JSON
		postData, err := json.Marshal(post)
		if err != nil {
			apiLogger.Error().Err(err).Msg("Failed to serialize post")
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
			apiLogger.Error().Err(err).Msg("Failed to publish to RabbitMQ")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to RabbitMQ"})
			return
		}

		apiLogger.Info().Str("title", post.Title).Msg("Post submitted")
		c.JSON(http.StatusOK, gin.H{"message": "Post submitted successfully"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
		apiLogger.Info().Msg("Ping endpoint hit")
	})

	// 启动 Gin 服务
	apiLogger.Info().Msg("Starting Gin server on :8080")
	if err := r.Run(":8080"); err != nil {
		apiLogger.Fatal().Err(err).Msg("Failed to start Gin server")
	}
}
