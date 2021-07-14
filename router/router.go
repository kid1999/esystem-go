package routes

import (
	"esystem/middleware"
	"github.com/gin-gonic/gin"
	"esystem/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 全局中间件
	router.Use(middleware.Logger())

	v1 := router.Group("/api/v1")
	{
		adminUser := new(controller.AdminUser)
		v1.GET("/admin-users", adminUser.Index)
		v1.POST("/admin-users", adminUser.Store)
		v1.PATCH("/admin-users/:id", adminUser.Update)
		v1.DELETE("/admin-users/:id", adminUser.Destroy)
		v1.GET("/admin-users/:id", adminUser.Show)

	}

	v2 := router.Group("/api/v2")
	{
		//v2.Use(middleware.JWT())
		user := new(controller.User)
		v2.POST("/test", user.Test)
		v2.POST("/register", user.Register)
		v2.POST("/login", user.Login)

		goods := new(controller.Goods)
		v2.GET("/goods/:id", goods.GoodsInfo)
		v2.GET("/goods", goods.FindAllGoods)
		v2.POST("/goods", goods.AddGoods)
		v2.PUT("/goods/:id", goods.UpdateGoods)
		v2.DELETE("/goods/:id", goods.DelGoods)
		v2.GET("/goods/search/:goodsName", goods.FindGoodsByName)
		v2.GET("/goods/user/:studentID", goods.FindGoodsByStudentID)

		admin := new(controller.Admin)
		v2.GET("/admin/users", user.FindAllUsers)
		v2.PUT("/admin/checkGoods", admin.CheckGoods)
		v2.PUT("/admin/checkUser", admin.CheckUser)

	}

	return router

}