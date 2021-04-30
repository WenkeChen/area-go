package model

type Option struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	OptionKey string `gorm:"index:option_key" json:"option_key"`
	OptionVal string `json:"option_val"`
}
