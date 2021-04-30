package adminservice

import (
	"AreaGo/model"
	"AreaGo/utils"
)

type AdminRegService struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func AdminAllowRegister() bool {
	var count int64
	model.Db.Model(&model.User{}).Where("role = ?", "admin").Count(&count)
	return count > 0
}

func (a *AdminRegService) Register() error {
	salt := utils.GetMixStr(6)
	userPwd := utils.EncryptPassword(a.Password, salt)
	admin := model.User{
		UserName: a.Username,
		UserPwd:  userPwd,
		Salt:     salt,
		Role:     "admin",
		Email:    a.Email,
		Source:   "local",
	}
	return model.Db.Create(&admin).Error
}
