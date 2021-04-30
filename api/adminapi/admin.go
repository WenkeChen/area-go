package adminapi

import (
	"AreaGo/formater"
	"AreaGo/service"
	"github.com/gin-gonic/gin"
	"os"
)

//检查程序是否安装过
func Install(c *gin.Context) {
	_, err := os.Open("./config/install.lock")
	if os.IsNotExist(err) {
		//todo go to register
		formater.Send(301, 30101, "", gin.H{"url": "/register"}, c)
		return
	}
	formater.Send(200, 2000, "ok", "", c)
}

func Info(c *gin.Context) {
	info := service.GetOptions([]string{"username", "avatar"})
	service.BuildUpOptions(info)
	formater.Send(200, 20000, "", service.BuildUpOptions(info), c)
}
