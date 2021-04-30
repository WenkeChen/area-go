package router

import (
	"AreaGo/api/adminapi"
)

func newAdmin() {

	admin := apiGroup.Group("/admin")

	//method login need not to check auth
	admin.GET("install", adminapi.Install)
	admin.POST("register", adminapi.RegisterAdmin)
	admin.POST("login", adminapi.LoginAdmin)

	//admin.Use(middleware.CheckLogin)
	//{
	admin.GET("info", adminapi.Info)
	//admin.POST("images", adminapi.UploadImg)
	admin.GET("posts", adminapi.ListPost)
	admin.GET("posts/:slug", adminapi.GetPostBySlug)
	admin.POST("posts", adminapi.AddPost)
	admin.PUT("posts/:id", adminapi.EditPost)
	admin.DELETE("posts/:id", adminapi.DeletePost)
	//}
}
