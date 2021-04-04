package middleware

import (
	"github.com/gin-gonic/gin"
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
)

func UserAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userDao := dao.NewUserDao()
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(200, model.Res{
				Code:  400,
			})
			ctx.Abort()
			return
		}
		user, err := userDao.UserCache(token)
		if err != nil {
			ctx.JSON(200, model.Res{
				Code:  400,
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userDao := dao.NewUserDao()
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(200, model.Res{
				Code:  400,
			})
			ctx.Abort()
			return
		}
		user, err := userDao.UserCache(token)
		if err != nil {
			ctx.JSON(200, model.Res{
				Code:  400,
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

