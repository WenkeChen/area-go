package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Owner       string `json:"owner"`
	OwnerEmail  string `gorm:"index:link_email" json:"owner_email"`
	Image       string `json:"image"`
	Target      string `gorm:"index:link_target" json:"target"`
	Description string `json:"description"`
	Visible     string `json:"visible"`
	Rss         string `gorm:"index:link_rss" json:"rss"`
	Note        string `json:"note"`
}
