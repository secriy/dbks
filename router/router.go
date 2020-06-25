package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.Cors())
	router.Use(middleware.CurrentUser())

	r := router.Group("/api")
	{
		// 公开接口
		r.POST("user/login", api.UserLogin)
		r.GET("news/all", api.ListNews)
		r.GET("offers/all", api.ListOffers)
		r.GET("products/all", api.ListProducts)
		r.GET("info/all", api.ListBasicInfo)
		// 登录保护接口
		login := r.Group("")
		login.Use(middleware.AuthRequired())
		{
			login.GET("user/me", api.UserMe)
			login.DELETE("user/logout", api.UserLogout)
		}
		// 管理员权限接口
		admin := r.Group("")
		admin.Use(middleware.AdminRequired())
		{
			// 用户CRUD
			admin.GET("user/all", api.UsersList)
			admin.POST("user/add", api.UserAdd)
			admin.DELETE("user/del/:id", api.UserDelete)
			admin.PUT("user/update/:id", api.UserUpdate)
			// 新闻CRUD
			admin.POST("news/add", api.CreateNews)
			admin.DELETE("news/del/:id", api.DeleteNews)
			admin.PUT("news/update/:id", api.UpdateNews)
			// 招聘CRUD
			admin.POST("offers/add", api.CreateOffers)
			admin.DELETE("offers/del/:id", api.DeleteOffers)
			admin.PUT("offers/update/:id", api.UpdateOffers)
			// 产品CRUD
			admin.POST("products/add", api.CreateProducts)
			admin.DELETE("products/del/:id", api.DeleteProducts)
			admin.PUT("products/update/:id", api.UpdateProducts)
			// 基础数据CRUD
			admin.POST("info/add", api.CreateBasicInfo)
			admin.DELETE("info/del/:option", api.DeleteBasicInfo)
			admin.PUT("info/update", api.UpdateBasicInfo)
		}
	}

	return router
}
