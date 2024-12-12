package service

import (
	"context"
	"errors"
	ticketsRpc "github.com/yongsuha/train-proto/ticket"
	"github.com/yongsuha/train-rpc/source/biz"
	"github.com/yongsuha/train-rpc/source/data/mysql"
	"log"
	"time"
)

type TicketsService struct {
	ticketsConfModel biz.TicketsConfRepo
}

func NewTicketsService() *TicketsService {
	return &TicketsService{
		ticketsConfModel: mysql.NewTicketsConfModel(),
	}
}

func (t *TicketsService) AddTicket(ctx context.Context, req *ticketsRpc.AddTicketReq) (*ticketsRpc.AddTicketResp, error) {
	tag := "TrainsService:AddTicket"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	if req.UserId == 0 {
		return nil, errors.New("UserId is empty")
	}
	if req.SeatId == 0 {
		return nil, errors.New("SeatId is empty")
	}
	if req.Price == 0 {
		return nil, errors.New("Price is empty")
	}
	if req.PurchaseTime == "" {
		return nil, errors.New("PurchaseTime is empty")
	}
	ticketMessage := &biz.TicketsConf{
		TrainId:      req.TrainId,
		UserId:       req.UserId,
		SeatId:       req.SeatId,
		Price:        req.Price,
		PurchaseTime: req.PurchaseTime,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	id, err := t.ticketsConfModel.AddTicket(ctx, ticketMessage)
	if err != nil {
		log.Println(ctx, tag, "AddTicket", "AddTicket err:%v", err)
		return nil, err
	}
	resp := &ticketsRpc.AddTicketResp{TicketId: id}
	return resp, nil
}

func (t *TicketsService) DelTicket(ctx context.Context, req *ticketsRpc.DelTicketReq) (*ticketsRpc.DelTicketResp, error) {
	tag := "TicketsService:DelTicket"
	if req.TicketId == 0 {
		return nil, errors.New("TicketId is empty")
	}
	err := t.ticketsConfModel.DelTicket(ctx, req.TicketId)
	if err != nil {
		log.Println(ctx, tag, "DelTicket", "DelTicket err:%v", err)
		return nil, err
	}
	resp := &ticketsRpc.DelTicketResp{Success: 1}
	return resp, nil
}

func (t *TicketsService) GetTicketDetail(ctx context.Context, req *ticketsRpc.GetTicketDetailReq) (*ticketsRpc.GetTicketDetailResp, error) {
	tag := "TicketsService:GetTicketDetail"
	reqInfo := false
	if req.TicketId != 0 {
		reqInfo = true
	}
	if req.UserId != 0 {
		reqInfo = true
	}
	if reqInfo == false {
		return nil, errors.New("reqInfo is empty")
	}
	getDetailParam := &biz.GetTicketDetailParam{
		TicketId: req.TicketId,
		UserId:   req.UserId,
	}
	tickets, err := t.ticketsConfModel.GetTicketDetail(ctx, getDetailParam)
	if err != nil {
		log.Println(ctx, tag, "GetTicketDetail", "GetTicketDetail err:%v", err)
		return nil, err
	}
	if tickets == nil {
		return nil, errors.New("tickets is empty")
	}
	ticketLists := []*ticketsRpc.TicketInfo{}
	for _, ticket := range tickets {
		ticketInfo := &ticketsRpc.TicketInfo{
			TicketId:     ticket.TicketId,
			UserId:       ticket.UserId,
			TrainId:      ticket.TrainId,
			SeatId:       ticket.SeatId,
			Price:        ticket.Price,
			PurchaseTime: ticket.PurchaseTime,
			CreateTime:   ticket.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:   ticket.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		ticketLists = append(ticketLists, ticketInfo)
	}
	resp := &ticketsRpc.GetTicketDetailResp{TicketDetail: ticketLists}
	return resp, nil
}

func (t *TicketsService) UpdateTicket(ctx context.Context, req *ticketsRpc.UpdateTicketReq) (*ticketsRpc.UpdateTicketResp, error) {
	tag := "TicketsService:UpdateTicket"
	if req.TicketId == 0 {
		return nil, errors.New("TicketId is empty")
	}
	getTicketparam := &biz.GetTicketDetailParam{
		TicketId: req.TicketId,
	}
	ticketInfos, err := t.ticketsConfModel.GetTicketDetail(ctx, getTicketparam)
	if err != nil {
		log.Println(ctx, tag, "UpdateTicket", "GetTicketDetail err:%v", err)
		return nil, err
	}
	if ticketInfos == nil {
		return nil, errors.New("ticketInfos is empty")
	}
	ticketMessage := ticketInfos[0]
	ticketInfo := &biz.TicketsConf{
		TicketId:     req.TicketId,
		UserId:       req.UserId,
		TrainId:      req.TrainId,
		SeatId:       req.SeatId,
		Price:        req.Price,
		PurchaseTime: req.PurchaseTime,
		CreateTime:   ticketMessage.CreateTime,
		UpdateTime:   time.Now(),
	}
	err = t.ticketsConfModel.UpdateTicket(ctx, ticketInfo)
	if err != nil {
		log.Println(ctx, tag, "UpdateTicket", "UpdateTicket err:%v", err)
		return nil, err
	}
	resp := &ticketsRpc.UpdateTicketResp{
		TicketId:     req.TicketId,
		UserId:       req.UserId,
		SeatId:       req.SeatId,
		TrainId:      req.TrainId,
		Price:        req.Price,
		PurchaseTime: req.PurchaseTime,
		CreateTime:   ticketMessage.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   ticketInfo.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}
