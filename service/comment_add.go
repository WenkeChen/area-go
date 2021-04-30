package service

import (
	"AreaGo/model"
	"time"
)

type CommentAddService struct {
	Uid       uint      `json:"uid"`
	Pid       uint      `json:"pid" binding:"required"`
	PostId    uint      `json:"post_id" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Ip        string    `json:"ip"`
	Useragent string    `json:"useragent"`
	Approved  uint      `json:"approved"`
	CreatedAt time.Time `json:"created_at"`
}

func (cas *CommentAddService) Add() bool {
	comment := model.Comment{
		Uid:       cas.Uid,
		Pid:       cas.Pid,
		PostId:    cas.PostId,
		Content:   cas.Content,
		Ip:        cas.Ip,
		Useragent: cas.Useragent,
		Approved:  cas.Approved,
	}
	if err := model.Db.Create(&comment).Error; err != nil {
		return false
	}
	return true
}
