package biz

import (
	"context"
	"time"
)

type TicketsConf struct {
	TicketId     int64     `json:"ticket_id" gorm:"column:ticket_id"`
	UserId       int64     `json:"user_id" gorm:"column:user_id"`
	TrainId      int64     `json:"train_id" gorm:"column:train_id"`
	SeatId       int64     `json:"seat_id" gorm:"column:seat_id"`
	Price        int64     `json:"price" gorm:"column:price"`
	PurchaseTime string    `json:"purchase_time" gorm:"column:purchase_time"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
}

type TicketsConfRepo interface {
	AddTicket(ctx context.Context, ticket *TicketsConf) (int64, error)
	DelTicket(ctx context.Context, id int64) error
	GetTicketDetail(ctx context.Context, param *GetTicketDetailParam) ([]*TicketsConf, error)
	UpdateTicket(ctx context.Context, ticket *TicketsConf) error
}

type (
	GetTicketDetailParam struct {
		TicketId int64 `json:"ticket_id"`
		UserId   int64 `json:"user_id"`
	}
)
