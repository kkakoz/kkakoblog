package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/middleware"
	"kkako-blog/internal/router/api"
)

func initArticleRouter(v1 *gin.RouterGroup) {
	articleG := v1.Group("/article")
	articleApi := api.NewArticleApi()
	{
		articleG.GET("", api.Wrap(articleApi.ArticlePage))
		articleG.GET("/:id", api.Wrap(articleApi.Info))
	}
	adminG := v1.Group("/article", middleware.UserAuth())
	adminG.POST("/add", api.Wrap(articleApi.ArticleAdd))
}
