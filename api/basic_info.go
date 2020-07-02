package api

import (
	"github.com/gin-gonic/gin"
	"server/service/basic_info"
)

// ListBasicInfo 基础数据列表接口
func ListBasicInfo(c *gin.Context) {
	listArticleService := basic_info.ListBasicInfoService{}
	if err := c.ShouldBind(&listArticleService); err == nil {
		res := listArticleService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateBasicInfo 基础数据更新
func UpdateBasicInfo(c *gin.Context) {
	updateBasicInfoService := basic_info.AddBasicInfoService{}
	if err := c.ShouldBind(&updateBasicInfoService); err == nil {
		res := updateBasicInfoService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteBasicInfo 基础数据删除接口
func DeleteBasicInfo(c *gin.Context) {
	deleteBasicInfoService := basic_info.DeleteBasicInfoService{}
	res := deleteBasicInfoService.Delete(c.Param("option"))
	c.JSON(200, res)
}
