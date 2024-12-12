package biz

import (
	"context"
	"time"
)

type OrderTicketConf struct {
	OrderId    int64     `json:"order_id" gorm:"column:order_id"`
	TicketId   int64     `json:"ticket_id" gorm:"column:ticket_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}

type OrderTicketConfRepo interface {
	AddOrderTicket(ctx context.Context, orderTicket *OrderTicketConf) error
	DelOrderTicket(ctx context.Context, param *OTParam) error
	GetOrderTicketDetail(ctx context.Context, orderId int64) ([]*OrderTicketConf, error)
}

type (
	OTParam struct {
		OrderId  int64 `json:"order_id"`
		TicketId int64 `json:"ticket_id"`
	}
)
