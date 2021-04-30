package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware(c *gin.Context) {
	//TODO 当`Access-Control-Allow-Credentials`为`true`时，`Access-Control-Allow-Origin`的值必须为具体的域名，而不能为`*`
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}
