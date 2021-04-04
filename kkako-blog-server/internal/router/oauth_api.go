package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/router/api"
)

func initOauthRouter(v1 *gin.RouterGroup) {
	oauthG := v1.Group("/oauth")
	oauthApi := api.NewOauthApi()
	{
		// github相关配置
		oauthG.GET("/github_conf", api.Wrap(oauthApi.GithubConfig))
		oauthG.POST("/github_auth", api.Wrap(oauthApi.GithubAuth))
	}
}
