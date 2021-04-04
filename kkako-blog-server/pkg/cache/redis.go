package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)


var Redis *redis.Client

func InitRedis(viper *viper.Viper) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("host") + ":" + viper.GetString("port"),
		Password: viper.GetString("password"),
		DB:       viper.GetInt("mysql"),
		PoolSize: viper.GetInt("poolSize"),
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}


