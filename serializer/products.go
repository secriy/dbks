package serializer

import "server/model"

// Product 招聘序列化器
type Product struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// BuildProducts 序列化招聘
func BuildProduct(item model.Products) Product {
	return Product{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildProducts 序列化招聘列表
func BuildProducts(items []model.Products) []Product {
	var products []Product
	for _, item := range items {
		productV := BuildProduct(item)
		products = append(products, productV)
	}
	return products
}
