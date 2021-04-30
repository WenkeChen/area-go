package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	UserPwd   string    `json:"user_pwd"`
	Salt      string    `json:"salt"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	Source    string    `json:"source"`
}
