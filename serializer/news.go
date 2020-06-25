package serializer

import "server/model"

// New 新闻序列化器
type New struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNews 序列化新闻
func BuildNew(item model.News) New {
	return New{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildNews 序列化新闻列表
func BuildNews(items []model.News) []New {
	var news []New
	for _, item := range items {
		newV := BuildNew(item)
		news = append(news, newV)
	}
	return news
}
