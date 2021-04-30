package router

import (
	"github.com/gin-gonic/gin"
)

var apiGroup *gin.RouterGroup

func New() *gin.Engine {
	router := gin.New()

	//router.LoadHTMLGlob("dist/*.html")    // 添加入口index.html
	//router.LoadHTMLFiles("static/*/*")	// 添加资源路径
	//router.Static("/static", "./dist/static") 	// 添加资源路径
	//router.StaticFile("/", "dist/index.html")  //前端入口

	apiGroup = router.Group("/api", gin.Logger())

	newIndex()
	newAdmin()

	return router
}
