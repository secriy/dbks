package offers

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// UpdateOffersService 招聘更新服务
type UpdateOffersService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Update 招聘更新
func (service *UpdateOffersService) Update(id string) serializer.Response {
	var offerC model.Offers
	var count = 0
	_ = database.DB.QueryRow(`SELECT  COUNT(*) FROM offers WHERE id = ?`, id).Scan(&count)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "招聘不存在",
		}
	}
	offerC.Title = service.Title
	offerC.Content = service.Content

	_, err := database.DB.Exec(`UPDATE offers SET title = ?,content = ? WHERE id = ?`, service.Title, service.Content, id)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "招聘更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildOffer(offerC),
		Msg:  "招聘更新成功",
	}
}
