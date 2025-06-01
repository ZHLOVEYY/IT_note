package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段，标准化一些
	Title      string `json:"title" gorm:"type:varchar(255);not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
}
