package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/router/api"
)

func initFileRouter(v1 *gin.RouterGroup) {
	oauthG := v1.Group("/file")
	fileApi := api.NewFileApi()
	{
		oauthG.POST("/upload", api.Wrap(fileApi.UploadFile))
	}
}
