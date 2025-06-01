package rabbitmq

import (
	"encoding/json"
	"os"
	"test/model"
	"time"

	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func StartConsumer(db *gorm.DB, logger zerolog.Logger) {
	// 连接 RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	var conn *amqp.Connection
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(rabbitMQURL)
		if err == nil {
			logger.Info().Msg("Successfully connected to RabbitMQ")
			break
		}
		logger.Info().Msgf("Failed to connect to RabbitMQ, retrying in 5 seconds... (attempt %d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	// 检查最终连接状态
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to RabbitMQ after all retries")
		return
	}

	if conn == nil {
		logger.Fatal().Msg("Connection is nil after successful connection attempt")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("post_queue", false, false, false, false, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to declare a queue")
	}

	// 消费消息
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to register a consumer")
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var post model.Post
			if err := json.Unmarshal(d.Body, &post); err != nil {
				logger.Error().Err(err).Msg("Failed to unmarshal post")
				continue
			}

			// 存储到数据库
			if err := db.Create(&post).Error; err != nil {
				logger.Error().Err(err).Msg("Failed to save post to database")
				continue
			}
			//消费 info
			logger.Info().Str("title", post.Title).Str("content", post.Content).Msg("Saved post")
		}
	}()

	logger.Info().Msg("Waiting for posts")
	<-forever
}
