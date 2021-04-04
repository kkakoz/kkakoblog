package router

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/middleware"
	"kkako-blog/internal/router/api"
)

func initUserRouter(e *gin.RouterGroup) {
	userG := e.Group("/user")
	userApi := api.NewUserApi()
	{
		userG.POST("/login", api.Wrap(userApi.Login))
		userG.POST("/admin_login", api.Wrap(userApi.AdminLogin))
		userG.GET("/info", api.Wrap(userApi.Info))
		userG.POST("/logout", api.Wrap(userApi.Logout))
		//userG.POST("/register", api.Wrap(userApi.Register))
	}
	adminG := e.Group("/admin", middleware.UserAuth())
	{
		adminG.GET("/:id", api.Wrap(userApi.Info))
	}
}
