package mysql

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitMysql(viper *viper.Viper) *sql.DB {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?" +
		"charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("user"), viper.GetString("password"),
		viper.GetString("host"), viper.GetString("port"),
		viper.GetString("name"))
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(viper.GetInt("max-open-conn"))
	sqlDB.SetMaxIdleConns(viper.GetInt("max-idle-conn"))
	return sqlDB
}


func DB() *gorm.DB {
	return db
}