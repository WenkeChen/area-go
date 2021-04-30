package router

import (
	"AreaGo/api/indexapi"
	"github.com/gin-gonic/gin"
)

func newIndex() {
	//查看文章列表
	apiGroup.GET("posts", indexapi.GetPostList)
	//文章详情
	apiGroup.POST("posts/:id", indexapi.GetPost)
	apiGroup.GET("posts/:slug", indexapi.GetPostBySlug)

	//文章评论
	apiGroup.GET("comments", indexapi.GetPostComments)

	//添加评论
	apiGroup.POST("comments", indexapi.AddComment)

	//分类列表
	apiGroup.GET("categories", indexapi.GetCategoryList)

	apiGroup.GET("archive/:type/:slug", func(context *gin.Context) {
		if context.Param("type") == "dates" {
			indexapi.GetPostsByDate(context)
		} else {
			indexapi.GetPostsByCategory(context)
		}
	})

	apiGroup.GET("moments", indexapi.GetPostComments)
}
