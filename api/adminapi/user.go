package adminapi

import (
	"AreaGo/formater"
	"AreaGo/model"
	"AreaGo/service/adminservice"
	"AreaGo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAdmin(c *gin.Context) {
	if adminservice.AdminAllowRegister() {
		formater.Send(200, 40300, "不允许注册为管理员", "", c)
		return
	}
	var admin adminservice.AdminRegService
	if err := c.ShouldBindJSON(&admin); err != nil {
		formater.Send(200, 40300, "参数错误", gin.H{"error": err.Error()}, c)
		return
	}
	if err := admin.Register(); err != nil {
		formater.Send(200, 50000, "注册失败", gin.H{"error": err.Error()}, c)
		return
	}
	formater.Send(200, 20000, "注册成功", "", c)
}

func LoginAdmin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model.User
	model.Db.Where("user_name = ?", username).First(&user)
	if user.ID == 0 || user.UserPwd != utils.EncryptPassword(password, user.Salt) {
		formater.Send(http.StatusUnprocessableEntity, 42201, "用户名或密码不正确", "", c)
		return
	}
	tokenString, err := utils.ReleaseToken(user.ID)
	if err != nil {
		formater.Send(http.StatusInternalServerError, 50001, "系统异常", "", c)
		return
	}

	formater.Send(200, 20000, "", gin.H{"token": tokenString}, c)
	return
}
