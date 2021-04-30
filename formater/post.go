package formater

import (
	"AreaGo/model"
)

type PostListFormat struct {
	Type  string       `json:"type"`
	Name  string       `json:"name"`
	Slug  string       `json:"slug"`
	List  []PostFormat `json:"list"`
	Count int64        `json:"count"`
}

type PostFormat struct {
	ID           uint   `json:"id"`
	CreatedAt    string `json:"created_at"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Cover        string `json:"cover"`
	Sticky       bool   `json:"sticky"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Commentable  bool   `json:"commentable"`
	CommentCount uint   `json:"comment_count"`
}

func BuildPostItem(post model.Post) PostFormat {
	var stickyStatus bool
	if post.Sticky != 0 {
		stickyStatus = true
	}
	return PostFormat{
		ID:           post.ID,
		CreatedAt:    post.CreatedAt.Format("2006-01-02"),
		Title:        post.Title,
		Slug:         post.Slug,
		Cover:        post.Cover,
		Description:  post.Description,
		Category:     post.Category.Slug,
		Sticky:       stickyStatus,
		CommentCount: post.CommentCount,
	}
}

func BuildPostList(posts []model.Post) []PostFormat {
	var articles = make([]PostFormat, len(posts))
	for i, t := range posts {
		articles[i] = BuildPostItem(t)
	}
	return articles
}

type PageItem struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func BuildPageItem(post model.Post) PageItem {
	return PageItem{
		ID:    post.ID,
		Title: post.Title,
	}
}

func BuildPageList(posts []model.Post) []PageItem {
	var pages = make([]PageItem, len(posts))
	for i, t := range posts {
		pages[i] = BuildPageItem(t)
	}
	return pages
}

type BlackholeItem struct {
	ID        uint                   `json:"id"`
	CreatedAt string                 `json:"created_at"`
	Title     string                 `json:"title"`
	Html      string                 `json:"html"`
	Annexes   map[string][]AnnexItem `json:"annexes"`
}

type PostDetails struct {
	ID           uint   `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Cover        string `json:"cover"`
	Category     string `json:"category"`
	Html         string `json:"html"`
	Commentable  bool   `json:"commentable"`
	CommentCount uint   `json:"comment_count"`
}

//获取post及post所有的meta
func BuildPostMetas(post model.Post) PostDetails {
	var commentStatus bool
	if post.Commentable == 1 {
		commentStatus = true
	}
	return PostDetails{
		ID:           post.ID,
		CreatedAt:    post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    post.UpdatedAt.Format("2006-01-02 15:04:05"),
		Title:        post.Title,
		Slug:         post.Slug,
		Cover:        post.Cover,
		Html:         post.Html,
		Commentable:  commentStatus,
		CommentCount: post.CommentCount,
		Category:     post.Category.Slug,
	}
}

//加密文章详情结构体
type SecretPostDetails struct {
	ID           uint   `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	CommentCount uint   `json:"comment_count"`
	Title        string `json:"title"`
	Cover        string `json:"cover"`
}

//时间归档列表
type DateArchiveItem struct {
	CreatedAt string `json:"created_at"`
	Count     uint   `json:"count"`
}

//归档统一返回结构体
type ArchivePosts struct {
	Type  string       `json:"type"`
	Name  string       `json:"name"`
	Count uint         `json:"count"`
	Posts []PostFormat `json:"posts"`
}
