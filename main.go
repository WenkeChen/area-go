package main

import (
	"AreaGo/config"
	"AreaGo/mailer"
	"AreaGo/model"
	"AreaGo/router"
)

func main() {
	//Initialize configuration
	config.New()

	//Connect to database
	model.New()

	//Connect to redis
	//redisClient := cache.New()
	//defer redisClient.Close()

	mailCh := mailer.New()
	defer close(mailCh)

	//Load routes
	r := router.New()

	//upload.New()

	//Run application
	_ = r.Run(":8888")
}
