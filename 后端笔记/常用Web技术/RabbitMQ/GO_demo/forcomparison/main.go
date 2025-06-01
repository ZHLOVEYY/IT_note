package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "net/http"
    "os"
    "compare/db"
    "compare/model"
)

func main() {
    // 初始化 zerolog
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

    // 初始化数据库
    db := db.InitDB()

    // 初始化 Gin
    r := gin.Default()

    // 提交帖子接口
    apiLogger := log.With().Str("component", "api").Logger()
    r.POST("/posts", func(c *gin.Context) {
        var post model.Post
        if err := c.ShouldBindJSON(&post); err != nil {
            apiLogger.Error().Err(err).Msg("Failed to bind JSON")
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // 直接写入数据库
        if err := db.Create(&post).Error; err != nil {
            apiLogger.Error().Err(err).Msg("Failed to save post to database")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post to database"})
            return
        }

        apiLogger.Info().Str("title", post.Title).Msg("Post saved")
        c.JSON(http.StatusOK, gin.H{"message": "Post saved successfully"})
    })

    // 测试接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
    // 启动 Gin 服务
    apiLogger.Info().Msg("Starting Gin server on :8080")
    if err := r.Run(":8080"); err != nil {
        apiLogger.Fatal().Err(err).Msg("Failed to start Gin server")
    }
}