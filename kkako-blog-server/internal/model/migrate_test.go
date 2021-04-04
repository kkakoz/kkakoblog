package model

import (
	"github.com/spf13/viper"
	"kkako-blog/log"
	"kkako-blog/pkg/cache"
	"kkako-blog/pkg/mysql"
	"kkako-blog/pkg/setting"
	"testing"
)

func TestMigrate(t *testing.T)  {
	setting.InitViper()
	log.InitLogger(viper.Sub("app"))
	cache.InitRedis(viper.Sub("redis"))
	defer cache.Redis.Close()
	sqldb := mysql.InitMysql(viper.Sub("mysql"))
	defer sqldb.Close()
	db := mysql.DB()
	err := db.AutoMigrate(&UserAuth{}, &Category{}, &Tag{})
	if err != nil {
		log.L().Fatal(err.Error())
	}
}
