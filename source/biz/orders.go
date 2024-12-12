package biz

import (
	"context"
	"time"
)

type OrdersConf struct {
	OrderId     int64     `json:"order_id" gorm:"column:order_id"`
	UserId      int64     `json:"user_id" gorm:"column:user_id"`
	TotalPrice  int64     `json:"total_price" gorm:"column:total_price"`
	OrderStatus string    `json:"order_status" gorm:"column:order_status"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

type OrdersConfRepo interface {
	AddOrder(ctx context.Context, order *OrdersConf) (int64, error)
	DelOrder(ctx context.Context, id int64) error
	GetOrderDetail(ctx context.Context, param *GetOrderDetailParam) ([]*OrdersConf, error)
	UpdateOrder(ctx context.Context, order *OrdersConf) error
}

type (
	GetOrderDetailParam struct {
		OrderId int64 `json:"order_id"`
		UserId  int64 `json:"user_id"`
	}
)
