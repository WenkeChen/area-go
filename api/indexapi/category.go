package indexapi

import (
	"AreaGo/formater"
	"AreaGo/model"
	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	var categories []model.Category
	model.Db.Find(&categories)
	formater.Send(200, 20000, "", categories, c)
	return
}
