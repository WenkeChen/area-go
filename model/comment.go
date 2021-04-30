package model

import "time"

type Comment struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Uid       uint      `json:"uid"`
	Pid       uint      `json:"pid"`
	PostId    uint      `json:"post_id"`
	Content   string    `json:"content"`
	Ip        string    `json:"ip"`
	Useragent string    `json:"useragent"`
	Approved  uint      `json:"approved" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignkey:Uid" json:"user"`
}
