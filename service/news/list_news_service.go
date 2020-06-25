package news

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListNewsService 新闻列表服务
type ListNewsService struct {
}

// List 文章列表
func (service *ListNewsService) List() serializer.Response {
	var newC model.News
	var news []model.News

	rows, err := database.DB.Query(`SELECT * FROM news ORDER BY id`)
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
		newC.ID = id
		newC.Title = title
		newC.Content = content
		newC.CreatedAt = createdAt

		news = append(news, newC)
	}

	return serializer.Response{
		Data: serializer.BuildNews(news),
	}
}
