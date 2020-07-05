package products

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// UpdateProductsService 产品更新服务
type UpdateProductsService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=20"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Update 产品更新
func (service *UpdateProductsService) Update(id string) serializer.Response {
	var productC model.Products
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM products WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "产品不存在",
		}
	}
	productC.Title = service.Title
	productC.Content = service.Content

	_, err := database.DB.Exec(`UPDATE products SET title = ?,content = ? WHERE id = ?`, service.Title, service.Content, id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "产品更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildProduct(productC),
		Msg:  "产品更新成功",
	}
}
