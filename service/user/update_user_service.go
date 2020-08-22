package user

import (
	"strconv"
	"time"

	"server/database"
	"server/model"

	"server/serializer"
)

// UpdateUserService 用户更新服务
type UpdateUserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=2,max=10"`
	Authority uint8  `form:"authority" json:"authority" binding:"required"`
}

// Update 用户更新
func (service *UpdateUserService) Update(id string) serializer.Response {
	var userVar model.User
	var count = 0
	var createAt time.Time
	if service.Authority > 127 {
		return serializer.Response{
			Code: 40007,
			Msg:  "Authority只能为小于128的正整数",
		}
	}
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM user WHERE id = ?`, id).Scan(&count)
	_ = database.DB.QueryRow(`SELECT create_at FROM user WHERE id = ?`, id).Scan(&createAt)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "用户不存在",
		}
	}
	if id == "1" && service.Authority != 1 {
		return serializer.Response{
			Code: 50003,
			Msg:  "默认管理员权限不允许更改",
		}
	}

	intNum, _ := strconv.Atoi(id)
	userVar.ID = uint64(intNum)
	userVar.UserName = service.UserName
	userVar.Authority = service.Authority
	userVar.CreatedAt = createAt

	_, err := database.DB.Exec(`
				UPDATE user SET username = ?,  authority=?
				WHERE id = ?`, userVar.UserName, userVar.Authority, id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "用户更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(userVar),
		Msg:  "用户更新成功",
	}
}
