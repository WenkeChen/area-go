package service

import (
	"AreaGo/model"
	"AreaGo/utils"
	"github.com/pkg/errors"
)

type PostService struct {
	ID          uint
	IsShow      uint
	Type        string
	WithAnnexes bool
	WithMetas   bool
	Fields      string
}

func (ps *PostService) List(pageStr string) ([]model.Post, int64, error) {
	var posts []model.Post
	var count int64

	size, offset, err := utils.GetOffset(pageStr)
	if err != nil {
		return nil, 0, errors.Wrap(err, "获取分页offset失败")
	}

	db := model.Db.Model(&model.Post{}).Where("is_show = ?", ps.IsShow)

	if ps.Type != "" {
		db = db.Where("type = ?", ps.Type)
	}

	if ps.WithAnnexes {
		db = db.Preload("Annexes")
	}
	if ps.WithMetas {
		db = db.Preload("Metas")
	}

	if err := db.Count(&count).Select(utils.BuildSelectFields(ps.Fields)).Order("is_sticky desc,created_at desc,id desc").Offset(offset).Limit(size).Find(&posts).Error; err != nil {
		return nil, 0, errors.Wrap(err, "查询post列表失败")
	}
	return posts, count, nil
}

func (ps *PostService) Details() model.Post {
	var post model.Post
	var db = model.Db
	if ps.WithAnnexes {
		db = db.Preload("Annexes")
	}
	if ps.WithMetas {
		db = db.Preload("Metas")
	}
	db.Select(utils.BuildSelectFields(ps.Fields)).First(&post, ps.ID)
	return post
}
