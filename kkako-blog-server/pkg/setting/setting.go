package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("app.port", 10003)
	viper.SetDefault("app.name", "blog")
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("mysql.user", "root")
	viper.SetDefault("mysql.password", "")
	viper.SetDefault("mysql.host", "127.0.0.1")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.name", "blog")
	viper.SetDefault("mysql.max-open-conn", 200)
	viper.SetDefault("mysql.max-idle-conn", 20)
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.poolSize", 10)
	viper.SetDefault("redis.mysql", 1)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err")
	}
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("配置文件修改", e.Name, "为", e.Op.String())
	//})
	//return viper.GetViper()
}