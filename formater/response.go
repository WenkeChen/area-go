package formater

import (
	"github.com/gin-gonic/gin"
)

func Send(httpStatus int, code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(httpStatus, map[string]interface{}{
		"code":    code,
		"message": msg,
		"data":    data,
	})
	c.Abort()
}
