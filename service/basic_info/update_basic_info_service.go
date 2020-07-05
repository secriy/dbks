package basic_info

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// AddBasicInfoService 基础数据投稿服务
type AddBasicInfoService struct {
	Name       string `form:"name" json:"name"`
	Address    string `form:"address" json:"address"`
	Department string `form:"department" json:"department"`
	Phone      string `form:"phone" json:"phone" `
	Email      string `form:"email" json:"email"`
	Url        string `form:"url" json:"url"`
}

// Create 基础数据投稿
func (service *AddBasicInfoService) Create() serializer.Response {
	BasicInfo := model.BasicInfo{
		Name:       service.Name,
		Address:    service.Address,
		Department: service.Department,
		Phone:      service.Phone,
		Email:      service.Email,
		Url:        service.Url,
		CreatedAt:  time.Now(),
	}

	// 创建基础数据
	_, err := database.DB.Exec(`
				REPLACE INTO basic_info(name, address, department, phone, email, url, create_at)
			  	VALUES (?,?,?,?,?,?,?)`,
		BasicInfo.Name, BasicInfo.Address, BasicInfo.Department, BasicInfo.Phone, BasicInfo.Email, BasicInfo.Url, BasicInfo.CreatedAt)
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "基础数据创建失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildBasicInfo(BasicInfo),
		Msg:  "基础数据创建成功",
	}
}
