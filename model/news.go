package model

import "time"

// 新闻模型
type News struct {
	ID        uint64
	Title     string
	Content   string
	CreatedAt time.Time
}
