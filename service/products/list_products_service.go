package products

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListProductsService 新闻列表服务
type ListProductsService struct {
}

// List 文章列表
func (service *ListProductsService) List() serializer.Response {
	var productC model.Products
	var products []model.Products

	rows, err := database.DB.Query(`SELECT * FROM products ORDER BY id`)
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
			title     string
			content   string
			createdAt time.Time
		)
		err = rows.Scan(&id, &title, &content, &createdAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		productC.ID = id
		productC.Title = title
		productC.Content = content
		productC.CreatedAt = createdAt

		products = append(products, productC)
	}

	return serializer.Response{
		Data: serializer.BuildProducts(products),
	}
}
