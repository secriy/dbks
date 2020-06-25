package user

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// UpdateUserService 用户更新服务
type UpdateUserService struct {
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=2,max=10"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=16"`
	Authority uint8  `form:"authority" json:"authority" binding:"required,numeric"`
}

// Update 用户更新
func (service *UpdateUserService) Update(id string) serializer.Response {
	var userC model.User
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM user WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "用户不存在",
		}
	}
	userC.UserName = service.UserName
	userC.Password = service.Password
	userC.Authority = service.Authority

	_, err := database.DB.Exec(`
				UPDATE user SET username = ?, password = ?, authority=?
				WHERE id = ?`, service.UserName, service.Password, service.Authority, id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "用户更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(userC),
		Msg:  "用户更新成功",
	}
}