package api

import (
	"github.com/gin-gonic/gin"
	"server/service/news"
)

// CreateNews 新闻投稿
func CreateNews(c *gin.Context) {
	createNewsService := news.AddNewsService{}
	if err := c.ShouldBind(&createNewsService); err == nil {
		res := createNewsService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListNews 新闻列表接口
func ListNews(c *gin.Context) {
	listArticleService := news.ListNewsService{}
	if err := c.ShouldBind(&listArticleService); err == nil {
		res := listArticleService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateNews 新闻更新
func UpdateNews(c *gin.Context) {
	updateNewsService := news.UpdateNewsService{}
	if err := c.ShouldBind(&updateNewsService); err == nil {
		res := updateNewsService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// DeleteNews 新闻删除接口
func DeleteNews(c *gin.Context) {
	deleteNewsService := news.DeleteNewsService{}
	res := deleteNewsService.Delete(c.Param("id"))
	c.JSON(200, res)
}
