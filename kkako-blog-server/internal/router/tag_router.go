package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/router/api"
)

func initTagRouter(e *gin.RouterGroup) {
	tagG := e.Group("/tag")
	tagApi := api.NewTagApi()
	{
		tagG.GET("", api.Wrap(tagApi.GetList))
		tagG.POST("", api.Wrap(tagApi.Add))
		tagG.PUT("", api.Wrap(tagApi.Edit))
		tagG.DELETE("", api.Wrap(tagApi.Delete))
		//tagG.POST("/register", api.Wrap(userApi.Register))
	}
}
