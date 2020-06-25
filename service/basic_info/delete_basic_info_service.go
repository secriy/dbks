package basic_info

import (
	"fmt"

	"server/database"
	"server/serializer"
)

// DeleteBasicInfosService 基础数据删除服务
type DeleteBasicInfoService struct {
}

// Delete 基础数据删除
func (service *DeleteBasicInfoService) Delete(option string) serializer.Response {
	str := `UPDATE basic_info SET ` + option + ` = '' LIMIT 1`
	fmt.Println(str)
	_, err := database.DB.Exec(str)
	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "基础数据删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Msg: "基础数据删除成功",
	}
}
