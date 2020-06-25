package api

import (
	"github.com/gin-gonic/gin"
	"server/service/products"
)

// CreateProducts 产品投稿
func CreateProducts(c *gin.Context) {
	createProductsService := products.AddProductsService{}
	if err := c.ShouldBind(&createProductsService); err == nil {
		res := createProductsService.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListProducts 产品列表接口
func ListProducts(c *gin.Context) {
	listProductsService := products.ListProductsService{}
	if err := c.ShouldBind(&listProductsService); err == nil {
		res := listProductsService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateProducts 产品更新
func UpdateProducts(c *gin.Context) {
	updateProductsService := products.UpdateProductsService{}
	if err := c.ShouldBind(&updateProductsService); err == nil {
		res := updateProductsService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteProducts 产品删除接口
func DeleteProducts(c *gin.Context) {
	deleteProductsService := products.DeleteProductsService{}
	res := deleteProductsService.Delete(c.Param("id"))
	c.JSON(200, res)
}
