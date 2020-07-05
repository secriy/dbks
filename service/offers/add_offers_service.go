package offers

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// AddOffersService 招聘投稿服务
type AddOffersService struct {
	Title   string `form:"title" json:"title" binding:"required,min=1,max=20"`
	Content string `form:"content" json:"content" binding:"required"`
}

// Create 招聘投稿
func (service *AddOffersService) Create() serializer.Response {
	Offers := model.Offers{
		Title:     service.Title,
		Content:   service.Content,
		CreatedAt: time.Now(),
	}

	// 创建招聘
	_, err := database.DB.Exec(`INSERT INTO dbks.offers(title,content,create_at)
		VALUES (?,?,?)`, Offers.Title, Offers.Content, Offers.CreatedAt)
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "招聘创建失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildOffer(Offers),
		Msg:  "招聘创建成功",
	}
}
