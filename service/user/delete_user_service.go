package user

import (
	"server/database"
	"server/serializer"
)

// DeleteUserService 用户删除服务
type DeleteUserService struct {
}

// Delete 用户删除
func (service *DeleteUserService) Delete(id string) serializer.Response {
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM user WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "用户不存在",
		}
	}
	if id == "1" {
		return serializer.Response{
			Code: 50003,
			Msg:  "默认管理员不允许删除",
		}
	}
	_, err := database.DB.Exec(`DELETE FROM user WHERE id = ?`, id)
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "用户删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "用户删除成功",
	}
}
