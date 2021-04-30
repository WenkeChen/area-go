package adminservice

import (
	"AreaGo/model"
	"AreaGo/utils"
)

type AdminLoginService struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (as *AdminLoginService) CheckLoginInfo() bool {
	var user model.User
	model.Db.Where("user_name = ?", as.Username).First(&user)
	return user.UserPwd == utils.EncryptPassword(as.Password, user.Salt)
}
