package rpc

import (
	"context"
	orderTicketRpc "github.com/yongsuha/train-proto/order_ticket"
	"github.com/yongsuha/train-rpc/source/service"
)

type OrderTicketServer struct {
	orderTicketCen *service.OrderTicketService
}

func NewOrderTicketService() *OrderTicketServer {
	return &OrderTicketServer{
		orderTicketCen: service.NewOrderTicketService(),
	}
}

func (ot *OrderTicketServer) AddOrderTicket(ctx context.Context, req *orderTicketRpc.AddOrderTicketReq) (*orderTicketRpc.AddOrderTicketResp, error) {
	return ot.orderTicketCen.AddOrderTicket(ctx, req)
}

func (ot *OrderTicketServer) GetOTDetail(ctx context.Context, req *orderTicketRpc.GetOTDetailReq) (*orderTicketRpc.GetOTDetailResp, error) {
	return ot.orderTicketCen.GetOTDetail(ctx, req)
}

func (ot *OrderTicketServer) DelOrderTicket(ctx context.Context, req *orderTicketRpc.DelOrderTicketReq) (*orderTicketRpc.DelOrderTicketResp, error) {
	return ot.orderTicketCen.DelOrderTicket(ctx, req)
}
