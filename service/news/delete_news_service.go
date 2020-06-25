package news

import (
	"server/database"
	"server/serializer"
)

// DeleteNewsService 新闻删除服务
type DeleteNewsService struct {
}

// Delete 新闻删除
func (service *DeleteNewsService) Delete(id string) serializer.Response {
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM news WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "新闻不存在",
		}
	}
	_, err := database.DB.Exec(`DELETE FROM news WHERE id = ?`, id)
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "新闻删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "新闻删除成功",
	}
}
