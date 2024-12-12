package biz

import (
	"context"
	"time"
)

type UsersConf struct {
	UserId      int64     `json:"user_id" gorm:"column:user_id"`
	UserName    string    `json:"user_name" gorm:"column:user_name"`
	PassWord    string    `json:"pass_word" gorm:"column:pass_word"`
	Email       string    `json:"email" gorm:"column:email"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

type UsersConfRepo interface {
	AddUser(ctx context.Context, user *UsersConf) (int64, error)
	DelUser(ctx context.Context, userId int64) error
	UpdateUser(ctx context.Context, user *UsersConf) error
	GetUserDetail(ctx context.Context, getDetailParam *GetUserDetailParam) (*UsersConf, error)
}

type (
	GetUserDetailParam struct {
		UserId      int64  `json:"user_id"`
		UserName    string `json:"user_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
)
