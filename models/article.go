package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title   string `binding:"required"`
	Content string `binding:"required"`
	Author  string `binding:"required"`
	Preview string `binding:"required"`
	Date    time.Time
	likes   int `gorm:"default:0"`
}
