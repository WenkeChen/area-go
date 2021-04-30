package model

import (
	"gorm.io/gorm"
)

const (
	TypeArticle = iota + 1
	TypeMoment
	TypeStatic
)

type Post struct {
	gorm.Model
	Title        string    `json:"title" gorm:"index:post_title;default:''"`
	Slug         string    `json:"slug" gorm:"index:post_slug"`
	Cover        string    `json:"cover" gorm:"default:''"`
	Description  string    `json:"description" gorm:"default:''"`
	CategoryId   uint      `json:"category_id" gorm:"not null"`
	Category     Category  `json:"category"`
	Markdown     string    `json:"markdown" gorm:"type:longtext;not null;"`
	Html         string    `json:"html" gorm:"type:longtext;not null;"`
	Pwd          string    `json:"pwd" gorm:"default:''"`
	Sticky       uint      `json:"sticky" gorm:"default:0"`
	Show         uint      `json:"show" gorm:"default:1"`
	Commentable  uint      `json:"commentable" gorm:"default:1"`
	CommentCount uint      `json:"comment_count" gorm:"default:1"`
	Comments     []Comment `json:"comments"`
	Tags         []*Tag    `json:"tags" gorm:"many2many:posts_tags;"`
	//Annexes      []Annex
}
