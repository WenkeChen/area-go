package adminapi

import (
	"AreaGo/formater"
	"AreaGo/model"
	"AreaGo/service"
	"AreaGo/utils"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

func ListPost(c *gin.Context) {
	var posts []model.Post
	var count int64
	model.Db.Scopes(service.Paginate(c)).Find(&posts).Count(&count)

	formater.Send(http.StatusOK, 20000, "ok", formater.PostListFormat{
		List:  formater.BuildPostList(posts),
		Count: count,
	}, c)
	return
}

func GetPostBySlug(c *gin.Context) {
	var post model.Post
	slugStr := c.Param("slug")
	model.Db.First(&post, "slug = ?", slugStr)
	if post.ID == 0 {
		formater.Send(http.StatusNotFound, 40000, "Not found", "", c)
		return
	}
	formater.Send(http.StatusOK, 20000, "", post, c)
	return
}

type PostForm struct {
	Title       string `form:"title" binding:"required" json:"title"`
	Description string `form:"description" binding:"required" json:"description"`
	CategoryId  uint   `form:"category_id" binding:"required" json:"category_id"`
	Markdown    string `form:"markdown" binding:"required" json:"markdown"`
	Html        string `form:"html" binding:"required" json:"html"`
}

func AddPost(c *gin.Context) {
	var postForm PostForm
	if err := c.BindJSON(&postForm); err != nil {
		panic(err)
	}
	var postData model.Post
	_ = utils.Form2model(postForm, &postData)
	postData.Slug = slug.Make(postData.Title)
	postData.Html = bluemonday.UGCPolicy().Sanitize(postData.Html)
	if err := model.Db.Create(&postData).Error; err != nil {
		formater.Send(400, 20000, err.Error(), nil, c)
		c.Abort()
		return
	}

	formater.Send(http.StatusCreated, 20000, "Success", nil, c)
	return
}

func EditPost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	model.Db.First(&post, id)
	if post.ID == 0 {
		formater.Send(http.StatusNotFound, 40000, "Post not found", nil, c)
		return
	}

	var postForm PostForm
	if err := c.BindJSON(&postForm); err != nil {
		panic(err)
	}
	var postData model.Post
	_ = utils.Form2model(postForm, &postData)
	postData.Slug = slug.Make(postData.Title)
	postData.Html = bluemonday.UGCPolicy().Sanitize(postData.Html)
	if err := model.Db.Model(&post).Updates(postData).Error; err != nil {
		formater.Send(http.StatusConflict, 40000, "Bad request", nil, c)
		return
	}

	formater.Send(http.StatusCreated, 20000, "Success", nil, c)
	return
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := model.Db.Delete(model.Post{}, "id = ?", id).Error; err != nil {
		formater.Send(200, 50001, "删除失败", gin.H{"error": err.Error()}, c)
		return
	}
	formater.Send(200, 20000, "删除成功", "", c)
	return
}
