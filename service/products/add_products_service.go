package products

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// AddProductsService 产品投稿服务
type AddProductsService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=20"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Create 产品投稿
func (service *AddProductsService) Create() serializer.Response {
	Products := model.Products{
		Title:     service.Title,
		Content:   service.Content,
		CreatedAt: time.Now(),
	}

	// 创建产品
	_, err := database.DB.Exec(`INSERT INTO dbks.products(title,content,create_at)
		VALUES (?,?,?)`, Products.Title, Products.Content, Products.CreatedAt)
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "产品失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildProduct(Products),
		Msg:  "产品创建成功",
	}
}
