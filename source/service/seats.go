package service

import (
	"context"
	"errors"
	seatsRpc "github.com/yongsuha/train-proto/seat"
	"github.com/yongsuha/train-rpc/source/biz"
	"github.com/yongsuha/train-rpc/source/data/mysql"
	"log"
	"time"
)

type SeatsService struct {
	seatsConfModel biz.SeatsConfRepo
}

func NewSeatsService() *SeatsService {
	return &SeatsService{
		seatsConfModel: mysql.NewSeatsConfModel(),
	}
}

func (s *SeatsService) AddSeat(ctx context.Context, req *seatsRpc.AddSeatReq) (*seatsRpc.AddSeatResp, error) {
	tag := "SeatsService:AddSeat"
	// 参数校验
	if req.SeatNumber == "" {
		return nil, errors.New("SeatNumber is empty")
	}
	if req.SeatType == "" {
		return nil, errors.New("SeatType is empty")
	}
	if req.IsAvailable == 0 {
		return nil, errors.New("IsAvailable is empty")
	}
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	// 填充数据库需要的各个参数
	seat := &biz.SeatsConf{
		SeatNumber:  req.SeatNumber,
		SeatType:    req.SeatType,
		IsAvailable: req.IsAvailable,
		TrainId:     req.TrainId,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	// 写入数据库
	id, err := s.seatsConfModel.AddSeat(ctx, seat)
	if err != nil {
		log.Println(ctx, tag, "AddSeat", "AddSeat err:%v", err)
		return nil, err
	}
	ret := &seatsRpc.AddSeatResp{Id: id}
	return ret, nil
}

func (s *SeatsService) DelSeat(ctx context.Context, req *seatsRpc.DelSeatReq) (*seatsRpc.DelSeatResp, error) {
	tag := "SeatsService:DelSeat"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	err := s.seatsConfModel.DelSeat(ctx, req.TrainId)
	if err != nil {
		log.Println(ctx, tag, "DelSeat", "DelSeat err:%v", err)
		return nil, err
	}
	resp := &seatsRpc.DelSeatResp{Success: 1}
	return resp, nil
}

func (s *SeatsService) UpdateSeat(ctx context.Context, req *seatsRpc.UpdateSeatReq) (*seatsRpc.UpdateSeatResp, error) {
	tag := "SeatsService:UpdateSeat"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	if req.SeatNumber == "" {
		return nil, errors.New("SeatNumber is empty")
	}
	if req.IsAvailable == 0 {
		return nil, errors.New("IsAvailable is empty")
	}
	// 判断是否存在该车次的该座位号信息
	getSeatDetailReq := &biz.GetSeatDetailParam{
		TrainId:    req.TrainId,
		SeatNumber: req.SeatNumber,
	}
	seatMessage, err := s.seatsConfModel.GetSeatDetail(ctx, getSeatDetailReq)
	if err != nil {
		log.Println(ctx, tag, "UpdateSeat", "GetSeatDetail err:%v", err)
		return nil, err
	}
	if seatMessage == nil {
		return nil, errors.New("seatMessage is empty")
	}
	// 修改数据库中该条记录
	seatInfo := &biz.SeatsConf{
		SeatId:      seatMessage.SeatId,
		TrainId:     seatMessage.TrainId,
		SeatNumber:  seatMessage.SeatNumber,
		SeatType:    seatMessage.SeatType,
		IsAvailable: req.IsAvailable,
		CreateTime:  seatMessage.CreateTime,
		UpdateTime:  time.Now(),
	}
	err = s.seatsConfModel.UpdateSeat(ctx, seatInfo)
	if err != nil {
		log.Println(ctx, tag, "UpdateSeat", "UpdateSeat err:%v", err)
		return nil, err
	}
	resp := &seatsRpc.UpdateSeatResp{
		TrainId:    req.TrainId,
		SeatNumber: req.SeatNumber,
	}
	return resp, nil
}

func (s *SeatsService) GetSeatDetail(ctx context.Context, req *seatsRpc.GetSeatDetailReq) (*seatsRpc.GetSeatDetailResp, error) {
	tag := "SeatsService:GetSeatDetail"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	if req.SeatNumber == "" {
		return nil, errors.New("SeatNumber is empty")
	}
	getDetailParam := &biz.GetSeatDetailParam{
		TrainId:    req.TrainId,
		SeatNumber: req.SeatNumber,
	}
	trainMessage, err := s.seatsConfModel.GetSeatDetail(ctx, getDetailParam)
	if err != nil {
		log.Println(ctx, tag, "GetSeatDetail", "GetSeatDetail err:%v", err)
		return nil, err
	}
	resp := &seatsRpc.GetSeatDetailResp{
		SeatId:      trainMessage.SeatId,
		TrainId:     trainMessage.TrainId,
		SeatNumber:  trainMessage.SeatNumber,
		SeatType:    trainMessage.SeatType,
		IsAvailable: trainMessage.IsAvailable,
		CreateTime:  trainMessage.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  trainMessage.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}
