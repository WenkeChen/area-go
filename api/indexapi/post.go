package indexapi

import (
	"AreaGo/formater"
	"AreaGo/model"
	"AreaGo/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"net/http"
	"strconv"
	"time"
)

// 获取文章列表
// todo Done
func GetPostList(c *gin.Context) {
	var posts []model.Post
	var count int64
	model.Db.Preload("Metas").Scopes(service.Paginate(c)).Find(&posts).Count(&count)

	formater.Send(http.StatusOK, 20000, "ok", formater.PostListFormat{
		List:  formater.BuildPostList(posts),
		Count: count,
	}, c)
}

// 通过类型slug获取文章列表
// todo Done
func GetPostsByCategory(c *gin.Context) {
	categorySlug := c.Param("slug")

	var category model.Category
	model.Db.Where("slug = ?", categorySlug).First(&category)
	var posts []model.Post
	_ = model.Db.Model(&category).Preload("Category").Preload("Tags").Association("Posts").Find(&posts)

	formater.Send(http.StatusOK, 20000, "ok", formater.PostListFormat{
		Name:  category.Name,
		Slug:  category.Slug,
		List:  formater.BuildPostList(posts),
		Count: category.Count,
	}, c)
	return
}

// 通过日期获取文章列表
// todo Done
func GetPostsByDate(c *gin.Context) {
	date := c.Param("slug")
	dateTime, _ := time.Parse("2006年01月", date)
	referenceTime := now.New(dateTime)
	beginTime := referenceTime.BeginningOfMonth()
	endTime := referenceTime.EndOfMonth()

	var posts []model.Post
	var count int64
	model.Db.Model(&model.Post{}).Where("type = ?", model.TypeArticle).
		Where("created_at BETWEEN ? AND ?", beginTime, endTime).
		Preload("Metas").Count(&count).Order("is_sticky DESC, created_at DESC, id DESC").
		Scopes(service.Paginate(c)).Find(&posts)
	formater.Send(http.StatusOK, 20000, "", formater.PostListFormat{
		Type:  "Date",
		Slug:  date,
		List:  formater.BuildPostList(posts),
		Count: count,
	}, c)
}

// 获取文章详情
// todo Done
func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		formater.Send(400, 40001, "请求参数错误", "", c)
		return
	}
	var post model.Post
	model.Db.Preload("Metas").First(&post, id)

	if post.Pwd != "" && post.Pwd != c.PostForm("password") {
		formater.Send(200, 40001, "密码错误", "", c)
		return
	}
	formater.Send(200, 20000, "", formater.BuildPostMetas(post), c)
	return
}

func GetPostBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var post model.Post
	model.Db.Preload("Metas").Where("slug = ?", slug).First(&post)

	formater.Send(200, 20000, "", formater.BuildPostMetas(post), c)
	return
}

func GetPostComments(c *gin.Context) {
	articleId := c.Query("aid")

	var post model.Post
	model.Db.Scopes(service.Paginate(c)).First(&post, articleId)
	if post.ID == 0 {
		formater.Send(http.StatusNotFound, 40001, "Not found", nil, c)
		return
	}

	if post.Pwd != "" && post.Pwd != c.PostForm("password") {
		formater.Send(200, 40001, "密码错误", "", c)
		return
	}

	var comments []model.Comment
	model.Db.Where("post_id = ?", post.ID).Where("approved = ?", true).Scopes(service.Paginate(c)).Find(&comments)

	var parentIds []uint
	for _, comment := range comments {
		parentIds = append(parentIds, comment.Pid)
	}
	var parents []model.Comment
	model.Db.Find(&parents, parentIds)

	formater.Send(http.StatusOK, 20000, "ok", formater.CommentListFormat{
		Id:   post.ID,
		List: formater.BuildCommentList(comments, parents),
	}, c)
	return
}
