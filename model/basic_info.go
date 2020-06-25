package model

import "time"

// 基础信息模型
type BasicInfo struct {
	Name       string
	Address    string
	Department string
	Phone      string
	Email      string
	Url        string
	CreatedAt  time.Time
}
