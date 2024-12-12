package rpc

import (
	"context"
	seatsRpc "github.com/yongsuha/train-proto/seat"
	"github.com/yongsuha/train-rpc/source/service"
)

type SeatsServer struct {
	seatsCen *service.SeatsService
}

func NewSeatsService() *SeatsServer {
	return &SeatsServer{
		seatsCen: service.NewSeatsService(),
	}
}

func (s *SeatsServer) AddSeat(ctx context.Context, req *seatsRpc.AddSeatReq) (*seatsRpc.AddSeatResp, error) {
	return s.seatsCen.AddSeat(ctx, req)
}

func (s *SeatsServer) DelSeat(ctx context.Context, req *seatsRpc.DelSeatReq) (*seatsRpc.DelSeatResp, error) {
	return s.seatsCen.DelSeat(ctx, req)
}

func (s *SeatsServer) UpdateSeat(ctx context.Context, req *seatsRpc.UpdateSeatReq) (*seatsRpc.UpdateSeatResp, error) {
	return s.seatsCen.UpdateSeat(ctx, req)
}

func (s *SeatsServer) GetSeatDetail(ctx context.Context, req *seatsRpc.GetSeatDetailReq) (*seatsRpc.GetSeatDetailResp, error) {
	return s.seatsCen.GetSeatDetail(ctx, req)
}
