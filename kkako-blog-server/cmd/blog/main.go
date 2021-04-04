//+build !wireinject

package main

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"kkako-blog/internal/router"
	"kkako-blog/log"
	"kkako-blog/pkg/cache"
	"kkako-blog/pkg/mysql"
	"kkako-blog/pkg/oauth"
	"kkako-blog/pkg/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	setting.InitViper()
	log.InitLogger(viper.Sub("app"))
	cache.InitRedis(viper.Sub("redis"))
	oauth.InitOauth(viper.Sub("oauth"))
	defer cache.Redis.Close()
	sqldb := mysql.InitMysql(viper.Sub("mysql"))
	defer sqldb.Close()
	engine := router.InitRouter()

	appViper := viper.Sub("app")
	srv := &http.Server{
		Addr: ":" + appViper.GetString("port"),
		Handler:engine,
	}
	log.L().Info("server run")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.L().Fatal("listen err: " + err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.L().Error("server shutdown",
			zap.String("err", err.Error()))
	}
	log.L().Info("server exit")
}
