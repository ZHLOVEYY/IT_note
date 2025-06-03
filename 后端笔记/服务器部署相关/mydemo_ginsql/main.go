package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100)"`
}

func main() {
	// 连接 MySQL 数据库
	// dsn := "root:1234@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"  //用于本地测试
	dsn := "user:password@tcp(mysql:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local" //连接 docker 中
	//注意 3306 是 mysql 的 docker 中的内部端口

	// 尝试连接数据库，最多尝试30次，每次间隔2秒  （docker-compose 中添加 health-check 也行）
	var db *gorm.DB
	var err error

	for attempts := 0; attempts < 30; attempts++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("数据库连接失败，2秒后重试: %v", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("无法连接到数据库，超过最大重试次数")
	}

	log.Println("成功连接到数据库")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// 自动迁移，创建表
	db.AutoMigrate(&User{})

	// 初始化 Gin
	r := gin.Default()

	// 简单 API：添加用户
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"message": "user created", "user": user})
	})

	// 简单 API：查询所有用户
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	// 启动服务
	r.Run(":8080")
}
