package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func New() *gorm.DB {
	if Db == nil {
		var err error
		dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.database"), viper.GetString("database.charset"))
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 禁用彩色打印
			},
		)
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "ar_",
			},
			Logger: newLogger,
		})
		if err != nil {
			panic(err)
		}
	}

	err := Db.AutoMigrate(&Annex{}, &Comment{}, &Link{}, &Option{}, &Post{}, &User{}, &Category{}, &Tag{})
	if err != nil {
		panic(err)
	}
	return Db
}
