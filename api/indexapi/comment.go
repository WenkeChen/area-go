package indexapi

import (
	"AreaGo/formater"
	"AreaGo/model"
	"AreaGo/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	user, _ := c.Get("user")

	if err := c.ShouldBindJSON(new(validator.CommentCreateValidator)); err != nil {
		formater.Send(http.StatusOK, 40444, err.Error(), "", c)
		return
	}
	comment := new(model.Comment)
	comment.Ip = c.ClientIP()
	comment.Useragent = c.GetHeader("User-Agent")
	comment.Uid = user.(model.User).ID

	if err := model.Db.Create(comment).Error; err != nil {
		formater.Send(http.StatusOK, 50000, "数据添加失败", "", c)
		return
	}

	formater.Send(http.StatusOK, 20000, "添加成功", "", c)
}
