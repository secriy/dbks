package news

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// AddNewsService 新闻投稿服务
type AddNewsService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Create 新闻投稿
func (service *AddNewsService) Create() serializer.Response {
	News := model.News{
		Title:   service.Title,
		Content: service.Content,
	}

	// 创建新闻
	_, err := database.DB.Exec(`INSERT INTO dbks.news(title,content,create_at)
		VALUES (?,?,?)`, News.Title, News.Content, time.Now())
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "新闻创建失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildNew(News),
		Msg:  "新闻创建成功",
	}
}
