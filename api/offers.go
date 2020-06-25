package api

import (
	"github.com/gin-gonic/gin"
	"server/service/offers"
)

// CreateOffers 招聘投稿
func CreateOffers(c *gin.Context) {
	createOffersService := offers.AddOffersService{}
	if err := c.ShouldBind(&createOffersService); err == nil {
		res := createOffersService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListOffers 招聘列表接口
func ListOffers(c *gin.Context) {
	listArticleService := offers.ListOffersService{}
	if err := c.ShouldBind(&listArticleService); err == nil {
		res := listArticleService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateOffers 招聘更新
func UpdateOffers(c *gin.Context) {
	updateOffersService := offers.UpdateOffersService{}
	if err := c.ShouldBind(&updateOffersService); err == nil {
		res := updateOffersService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteOffers 招聘删除接口
func DeleteOffers(c *gin.Context) {
	deleteOffersService := offers.DeleteOffersService{}
	res := deleteOffersService.Delete(c.Param("id"))
	c.JSON(200, res)
}
