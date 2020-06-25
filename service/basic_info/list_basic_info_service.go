package basic_info

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListBasicInfoService 基础数据查询服务
type ListBasicInfoService struct {
}

// List 基础数据列表
func (service *ListBasicInfoService) List() serializer.Response {
	var basicInfoC model.BasicInfo

	rows, err := database.DB.Query(`SELECT * FROM basic_info LIMIT 1`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		var (
			name       string
			address    string
			department string
			phone      string
			email      string
			url        string
			createdAt  time.Time
		)
		err = rows.Scan(&name, &address, &department, &phone, &email, &url, &createdAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		basicInfoC.Name = name
		basicInfoC.Address = address
		basicInfoC.Department = department
		basicInfoC.Phone = phone
		basicInfoC.Email = email
		basicInfoC.Url = url
		basicInfoC.CreatedAt = createdAt
	}

	return serializer.Response{
		Data: serializer.BuildBasicInfo(basicInfoC),
	}
}
