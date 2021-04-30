package middleware

import (
	"AreaGo/formater"
	"AreaGo/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckLogin(c *gin.Context) {
	// Token from another example.  This token is expired
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" || strings.HasPrefix(tokenString, "Bearer") {
		formater.Send(200, 40100, "token有误", "", c)
		c.Abort()
		return
	}

	tokenString = tokenString[7:]
	token, claims, err := utils.ParseToken(tokenString)

	if err != nil || !token.Valid {
		formater.Send(200, 40100, "unauthorized", "", c)
		c.Abort()
		return
	}

	uid := claims.Uid
	c.Set("uid", uid)
	c.Next()
}
