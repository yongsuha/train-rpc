package biz

import (
	"context"
	"time"
)

type SeatsConf struct {
	SeatId      int64     `json:"seat_id" gorm:"column:seat_id"`
	TrainId     int64     `json:"train_id" gorm:"column:train_id"`
	SeatNumber  string    `json:"seat_number" gorm:"column:seat_number"`
	SeatType    string    `json:"seat_type" gorm:"column:seat_type"`
	IsAvailable int64     `json:"is_available" gorm:"column:is_available"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

type SeatsConfRepo interface {
	AddSeat(ctx context.Context, seat *SeatsConf) (int64, error)
	DelSeat(ctx context.Context, trainId int64) error
	UpdateSeat(ctx context.Context, seat *SeatsConf) error
	GetSeatDetail(ctx context.Context, getDetailParam *GetSeatDetailParam) (*SeatsConf, error)
}

type (
	GetSeatDetailParam struct {
		TrainId    int64  `json:"train_id"`
		SeatNumber string `json:"seat_number"`
	}
)
