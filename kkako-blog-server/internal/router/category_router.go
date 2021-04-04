package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/router/api"
)

func initCategoryRouter(e *gin.RouterGroup) {
	categoryG := e.Group("/category")
	CategoryApi := api.NewCategoryApi()
	{
		categoryG.GET("", api.Wrap(CategoryApi.GetList))
		categoryG.POST("", api.Wrap(CategoryApi.Add))
		categoryG.PUT("/enable", api.Wrap(CategoryApi.Enable))
		categoryG.PUT("", api.Wrap(CategoryApi.Edit))
		categoryG.DELETE("", api.Wrap(CategoryApi.Delete))
		//categoryG.POST("/register", api.Wrap(userApi.Register))
	}
}
