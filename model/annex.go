package model

type Annex struct {
	ID     uint `gorm:"primary_key" json:"id"`
	PostId uint `json:"post_id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Src    string `json:"src"`
}
