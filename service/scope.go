package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize := viper.GetInt("app.pageSize")
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
