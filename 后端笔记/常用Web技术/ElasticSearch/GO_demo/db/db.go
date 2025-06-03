package db

import (
    "log"
    "esdemo/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:1234@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    }
	//填写对应的root 后面的密码，以及 database 的名字（我的是 test）
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    err = db.AutoMigrate(&models.Post{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    return db
}

