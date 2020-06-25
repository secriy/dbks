package model

import "time"

// 产品模型
type Products struct {
	ID        uint64
	Title     string
	Content   string
	CreatedAt time.Time
}
