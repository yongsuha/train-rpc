package rpc

import (
	"context"
	ticketsRpc "github.com/yongsuha/train-proto/ticket"
	"github.com/yongsuha/train-rpc/source/service"
)

type TicketsServer struct {
	ticketsCen *service.TicketsService
}

func NewTicketsService() *TicketsServer {
	return &TicketsServer{
		ticketsCen: service.NewTicketsService(),
	}
}

func (t *TicketsServer) AddTicket(ctx context.Context, req *ticketsRpc.AddTicketReq) (*ticketsRpc.AddTicketResp, error) {
	return t.ticketsCen.AddTicket(ctx, req)
}

func (t *TicketsServer) DelTicket(ctx context.Context, req *ticketsRpc.DelTicketReq) (*ticketsRpc.DelTicketResp, error) {
	return t.ticketsCen.DelTicket(ctx, req)
}

func (t *TicketsServer) GetTicketDetail(ctx context.Context, req *ticketsRpc.GetTicketDetailReq) (*ticketsRpc.GetTicketDetailResp, error) {
	return t.ticketsCen.GetTicketDetail(ctx, req)
}

func (t *TicketsServer) UpdateTicket(ctx context.Context, req *ticketsRpc.UpdateTicketReq) (*ticketsRpc.UpdateTicketResp, error) {
	return t.ticketsCen.UpdateTicket(ctx, req)
}
