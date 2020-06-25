package model

import "time"

// 企业招聘模型
type Offers struct {
	ID        uint64
	Title     string
	Content   string
	CreatedAt time.Time
}
