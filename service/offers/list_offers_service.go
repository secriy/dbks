package offers

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListOffersService 招聘列表服务
type ListOffersService struct {
}

// List 招聘列表
func (service *ListOffersService) List() serializer.Response {
	var offerC model.Offers
	var offers []model.Offers

	rows, err := database.DB.Query(`SELECT * FROM offers ORDER BY id`)
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
		offerC.ID = id
		offerC.Title = title
		offerC.Content = content
		offerC.CreatedAt = createdAt

		offers = append(offers, offerC)
	}

	return serializer.Response{
		Data: serializer.BuildOffers(offers),
	}
}
