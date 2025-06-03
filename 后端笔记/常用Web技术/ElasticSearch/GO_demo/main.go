package main

import (
	"esdemo/db"
	"esdemo/es"
	"esdemo/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// 初始化 zerolog 日志系统
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// 初始化 MySQL
	db := db.InitDB()

	// 初始化 Elasticsearch  es是我引入的包
	es.InitES()

	// 初始化 Gin
	r := gin.Default()

	// 提交帖子接口
	apiLogger := log.With().Str("component", "api").Logger()
	r.POST("/posts", func(c *gin.Context) {
		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			apiLogger.Error().Err(err).Msg("Failed to bind JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 写入 MySQL
		if err := db.Create(&post).Error; err != nil {
			apiLogger.Error().Err(err).Msg("Failed to save post to MySQL")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
			return
		}

		// 索引到 Elasticsearch
		if err := es.IndexPost(&post); err != nil {
			apiLogger.Error().Err(err).Msg("Failed to index post to Elasticsearch")
			// 继续返回成功，避免影响用户体验
		}

		apiLogger.Info().Str("title", post.Title).Msg("Post saved")
		c.JSON(http.StatusOK, gin.H{"message": "Post saved successfully"})
	})

	// 搜索帖子接口
	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
			return
		}

		posts, err := es.SearchPosts(query)
		if err != nil {
			apiLogger.Error().Err(err).Msg("Failed to search posts")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"posts": posts})
	})

	// 测试接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	
	// 启动 Gin 服务
	apiLogger.Info().Msg("Starting Gin server on :8080")
	if err := r.Run(":8080"); err != nil {
		apiLogger.Fatal().Err(err).Msg("Failed to start Gin server")
	}
}
