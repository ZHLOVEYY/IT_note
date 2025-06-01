package db

import (
    "github.com/rs/zerolog/log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "compare/model"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        log.Fatal().Msg("MYSQL_DSN environment variable not set")
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to connect to database")
    }

    // 自动迁移，创建 posts 表
    err = db.AutoMigrate(&model.Post{})
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to migrate database")
    }

    return db
}