package news

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// UpdateNewsService 新闻更新服务
type UpdateNewsService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Update 新闻更新
func (service *UpdateNewsService) Update(id string) serializer.Response {
	var newC model.News
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM news WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "新闻不存在",
		}
	}
	newC.Title = service.Title
	newC.Content = service.Content

	_, err := database.DB.Exec(`UPDATE news SET title = ?,content = ? WHERE id = ?`, service.Title, service.Content, id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "新闻更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildNew(newC),
		Msg:  "新闻更新成功",
	}
}
