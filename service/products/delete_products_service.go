package products

import (
	"server/database"
	"server/serializer"
)

// DeleteProductsService 产品删除服务
type DeleteProductsService struct {
}

// Delete 产品删除
func (service *DeleteProductsService) Delete(id string) serializer.Response {
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM products WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "产品不存在",
		}
	}
	_, err := database.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "产品删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "产品删除成功",
	}
}
