package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/middleware"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.CORS(), gin.Recovery(), middleware.GinLog())
	v1 := engine.Group("/blog")
	initUserRouter(v1)
	initArticleRouter(v1)
	initOauthRouter(v1)
	initCategoryRouter(v1)
	initTagRouter(v1)
	return engine
}