package service

//
//import (
//	"AreaGo/model"
//	"AreaGo/utils"
//	"encoding/json"
//	"github.com/pkg/errors"
//	"time"
//)
//
//type PostService struct {
//	CreatedAt   time.Time `form:"created_at"`
//	Category        uint    `form:"type"`
//	Title       string    `form:"title" binding:"required"`
//	CateId      uint      `json:"cate_id" form:"cate_id" binding:"required"`
//	Cover       string    `form:"cover"`
//	Description     string    `form:"summary"`
//	Markdown    string    `form:"markdown" binding:"required"`
//	Html        string    `form:"html" binding:"required"`
//	Pwd         string    `form:"pwd"`
//	Show      uint      `json:"is_show" form:"is_show"`
//	Commentable uint      `form:"commentable"`
//}
//
//func (p *PostService) Add() error {
//	if img := utils.GetFirstImg(p.Html); img == "" {
//		p.Cover = GetOption("cover")
//	}else {
//		p.Cover = img
//	}
//	p.Description = utils.SubStr(p.Html, 30)
//	post := model.Post{
//		CateId:      p.CateId,
//		Title:       p.Title,
//		Cover:       p.Cover,
//		Description:     p.Description,
//		Markdown:    p.Markdown,
//		Html:        p.Html,
//		Pwd:         p.Pwd,
//		Show:      p.Show,
//		Commentable: p.Commentable,
//	}
//	err := model.Db.Create(&post).Error
//	if err != nil {
//		return errors.Wrap(err, "add post error")
//	}
//	return nil
//}
//
//func (p *PostService) Edit(id int) error {
//	var post model.Post
//	model.Db.Find(&post, id)
//
//	var data map[string]interface{}
//	marshalJson, err := json.Marshal(p)
//	if err != nil {
//		return errors.Wrap(err, "json marshal error")
//	}
//	err = json.Unmarshal(marshalJson, data)
//	if err != nil {
//		return errors.Wrap(err, "json unmarshal error")
//	}
//	errs := model.Db.Model(&post).Updates(data).GetErrors()
//	if errs != nil {
//		return errors.Wrap(err, "update post error")
//	}
//	return nil
//}
//
