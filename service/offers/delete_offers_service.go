package offers

import (
	"server/database"
	"server/serializer"
)

// DeleteOffersService 招聘删除服务
type DeleteOffersService struct {
}

// Delete 招聘删除
func (service *DeleteOffersService) Delete(id string) serializer.Response {
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM offers WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "招聘不存在",
		}
	}
	_, err := database.DB.Exec(`DELETE FROM offers WHERE id = ?`, id)
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "招聘删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "招聘删除成功",
	}
}
