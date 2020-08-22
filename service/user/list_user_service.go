package user

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListUsersService 用户列表服务
type ListUsersService struct {
}

// List 用户列表
func (service *ListUsersService) List() serializer.Response {
	var userVar model.User
	var users []model.User

	rows, err := database.DB.Query(`SELECT * FROM user ORDER BY id`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		var (
			id        uint64
			userName  string
			password  string
			authority uint8
			createdAt time.Time
		)

		err = rows.Scan(&id, &userName, &password, &authority, &createdAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		userVar.ID = id
		userVar.UserName = userName
		userVar.Authority = authority
		userVar.CreatedAt = createdAt

		users = append(users, userVar)
	}

	return serializer.Response{
		Data: serializer.BuildUsers(users),
	}
}
