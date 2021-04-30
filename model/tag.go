package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string
	Slug  string
	Count uint
}
