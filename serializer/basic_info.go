package serializer

import "server/model"

// BasicInfo 基础数据序列化器
type BasicInfo struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Department string `json:"department"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Url        string `json:"url"`
	CreatedAt  int64  `json:"created_at"`
}

// BuildBasicInfos 序列化基础数据
func BuildBasicInfo(item model.BasicInfo) BasicInfo {
	return BasicInfo{
		Name:       item.Name,
		Address:    item.Address,
		Department: item.Department,
		Phone:      item.Phone,
		Email:      item.Email,
		Url:        item.Url,
		CreatedAt:  item.CreatedAt.Unix(),
	}
}
