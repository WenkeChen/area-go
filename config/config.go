package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func New() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	gin.SetMode(viper.GetString("app.mode"))
}
