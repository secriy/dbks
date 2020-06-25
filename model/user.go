package model

import (
	"time"

	"server/database"
)

// 用户模型
type User struct {
	ID        uint64
	UserName  string
	Password  string
	Authority uint8
	CreatedAt time.Time
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := database.DB.QueryRow(`SELECT * FROM user WHERE ID = ?`, ID).Scan(&user.ID, &user.UserName, &user.Password, &user.Authority, &user.CreatedAt)
	return user, result
}
