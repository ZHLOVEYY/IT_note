package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	Content string `json:"content" gorm:"type:text;not null"`
}
