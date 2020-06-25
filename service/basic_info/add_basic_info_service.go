package basic_info

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// AddBasicInfoService 基础数据投稿服务
type AddBasicInfoService struct {
	Name       string `form:"name" json:"name" binding:"required,min=1,max=20"`
	Address    string `form:"address" json:"address" binding:"required"`
	Department string `form:"department" json:"department" binding:"required"`
	Phone      string `form:"phone" json:"phone" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required"`
	Url        string `form:"url" json:"url" binding:"required"`
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
	}

	// 创建基础数据
	_, err := database.DB.Exec(`
				REPLACE INTO basic_info(name, address, department, phone, email, url, create_at)
			  	VALUES (?,?,?,?,?,?,?)`,
		BasicInfo.Name, BasicInfo.Address, BasicInfo.Department, BasicInfo.Phone, BasicInfo.Email, BasicInfo.Url, time.Now())
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
