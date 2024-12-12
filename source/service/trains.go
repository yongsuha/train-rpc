package service

import (
	"context"
	"errors"
	"fmt"
	trainsRpc "github.com/yongsuha/train-proto/train"
	"github.com/yongsuha/train-rpc/source/biz"
	"github.com/yongsuha/train-rpc/source/data/mysql"
	"log"
	"time"
)

type TrainsService struct {
	trainsConfModel biz.TrainsConfRepo
	seatsConfModel  biz.SeatsConfRepo
}

func NewTrainsService() *TrainsService {
	return &TrainsService{
		trainsConfModel: mysql.NewTrainsConfModel(),
		seatsConfModel:  mysql.NewSeatsConfModel(),
	}
}

func (t *TrainsService) AddTrain(ctx context.Context, req *trainsRpc.AddTrainReq) (*trainsRpc.AddTrainResp, error) {
	// 参数校验
	tag := "TrainsService:AddTrain"
	if req.TrainNumber == "" {
		return nil, errors.New("TrainNumber is empty")
	}
	if req.ArrivalTime == "" {
		return nil, errors.New("ArrivalTime is empty")
	}
	if req.ArrivalStation == "" {
		return nil, errors.New("ArrivalStation is empty")
	}
	if req.DepartureStation == "" {
		return nil, errors.New("DepartureStation is empty")
	}
	if req.DepartureTime == "" {
		return nil, errors.New("DepartureTime is empty")
	}
	if req.Seats == nil {
		return nil, errors.New("Seats is empty")
	}
	// 填充数据库需要的各个字段
	var totalSeats int64 = 0
	for _, seat := range req.Seats {
		totalSeats += seat.SeatNum
	}
	train := &biz.TrainsConf{
		TrainNumber:      req.TrainNumber,
		DepartureTime:    req.DepartureTime,
		DepartureStation: req.DepartureStation,
		ArrivalStation:   req.ArrivalStation,
		ArrivalTime:      req.ArrivalTime,
		TotalSeat:        totalSeats,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
	}
	// 写入数据库
	id, err := t.trainsConfModel.AddTrain(ctx, train)
	if err != nil {
		log.Println(ctx, tag, "AddTrain", "AddTrain err:%v", err)
		return nil, err
	}
	// 初始化座位信息写入 seats 表
	for _, seat := range req.Seats {
		// 生成座位号，这里简单假设座位号从 1 开始递增
		for i := 1; i <= int(seat.SeatNum); i++ {
			seatNumber := fmt.Sprintf("%s-%d", seat.SeatType, i)
			newSeat := &biz.SeatsConf{
				TrainId:     id,
				SeatNumber:  seatNumber,
				SeatType:    seat.SeatType,
				IsAvailable: 1, // 默认可用
				CreateTime:  time.Now(),
				UpdateTime:  time.Now(),
			}
			// 插入座位信息到 seats 表
			_, err := t.seatsConfModel.AddSeat(ctx, newSeat)
			if err != nil {
				log.Println(ctx, tag, "AddTrain", "AddSeat err:%v", err)
				return nil, err
			}
		}
	}
	resp := &trainsRpc.AddTrainResp{Id: id}
	return resp, nil
}

// ToDo 应该可以考虑用协程吧
func (t *TrainsService) UpdateTrain(ctx context.Context, req *trainsRpc.UpdateTrainReq) (*trainsRpc.UpdateTrainResp, error) {
	tag := "TrainsService:UpdateTrain"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	// 先获取到要修改车次的信息
	getDetailReq := &trainsRpc.GetTrainDetailReq{Id: req.TrainId}
	trainDetail, err := t.GetTrainDetail(ctx, getDetailReq)
	if err != nil {
		log.Println(ctx, tag, "UpdateTrain", "GetTrainDetail err:%v", err)
		return nil, err
	}
	// 修改 trains 表信息
	parsedUpdateTime, err := time.Parse("2006-01-02 15:04:05", trainDetail.UpdateTime)
	var totalSeats int64 = 0
	for _, seat := range req.Seats {
		totalSeats += seat.SeatNum
	}
	trainMessage := &biz.TrainsConf{
		TrainId:          req.TrainId,
		TrainNumber:      req.TrainNumber,
		TotalSeat:        totalSeats,
		DepartureStation: req.DepartureStation,
		DepartureTime:    req.DepartureTime,
		ArrivalTime:      req.ArrivalTime,
		ArrivalStation:   req.ArrivalStation,
		CreateTime:       parsedUpdateTime,
		UpdateTime:       time.Now(),
	}
	err = t.trainsConfModel.UpdateTrain(ctx, trainMessage)
	if err != nil {
		log.Println(ctx, tag, "UpdateTrain", "UpdateTrain err:%v", err)
		return nil, err
	}
	// 如果车次的座位信息修改了 也要修改 seats 表
	if totalSeats != trainDetail.TotalSeat {
		// 先删除原来的座位信息
		err = t.seatsConfModel.DelSeat(ctx, req.TrainId)
		if err != nil {
			log.Println(ctx, tag, "UpdateTrain", "DelSeat err:%v", err)
			return nil, err
		}
		// 再重新插入座位信息
		for _, seat := range req.Seats {
			// 生成座位号，这里简单假设座位号从 1 开始递增
			for i := 1; i <= int(seat.SeatNum); i++ {
				seatNumber := fmt.Sprintf("%s-%d", seat.SeatType, i)
				newSeat := &biz.SeatsConf{
					TrainId:     req.TrainId,
					SeatNumber:  seatNumber,
					SeatType:    seat.SeatType,
					IsAvailable: 1, // 默认可用
					CreateTime:  time.Now(),
					UpdateTime:  time.Now(),
				}
				// 插入座位信息到 seats 表
				_, err := t.seatsConfModel.AddSeat(ctx, newSeat)
				if err != nil {
					log.Println(ctx, tag, "AddTrain", "AddSeat err:%v", err)
					return nil, err
				}
			}
		}
	}
	// ToDo 针对车次信息可能还需要删除/修改缓存 针对座位信息可能还需要修改 es
	//todo 之后可能还会涉及到要修改其他的表
	resp := &trainsRpc.UpdateTrainResp{Id: req.TrainId}
	return resp, nil
}

func (t *TrainsService) GetTrainDetail(ctx context.Context, req *trainsRpc.GetTrainDetailReq) (*trainsRpc.GetTrainDetailResp, error) {
	tag := "TrainsService:GetTrainDetail"
	if req.Id == 0 {
		return nil, errors.New("TrainId is empty")
	}
	trainDetail, err := t.trainsConfModel.GetTrainDetail(ctx, req.Id)
	if err != nil {
		log.Println(ctx, tag, "GetTrainDetail", "GetTrainDetail err:%v", err)
		return nil, err
	}
	if trainDetail == nil {
		return nil, errors.New("the train detail is empty")
	}

	resp := &trainsRpc.GetTrainDetailResp{
		TrainId:          trainDetail.TrainId,
		TrainNumber:      trainDetail.TrainNumber,
		DepartureTime:    trainDetail.DepartureTime,
		DepartureStation: trainDetail.DepartureStation,
		ArrivalStation:   trainDetail.ArrivalStation,
		ArrivalTime:      trainDetail.ArrivalTime,
		TotalSeat:        trainDetail.TotalSeat,
		CreateTime:       trainDetail.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:       trainDetail.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}

func (t *TrainsService) DelTrain(ctx context.Context, req *trainsRpc.DelTrainReq) (*trainsRpc.DelTrainResp, error) {
	tag := "TrainsService:DelTrain"
	if req.TrainId == 0 {
		return nil, errors.New("TrainId is empty")
	}
	// ToDo 这里应该要检查一下有没有人已经买了该车次的票吧 也有可能有人买也照样可以删除 我再考虑考虑
	// 检查车次是否存在 todo 这里后续应该要从缓存中读取
	trainMessage, err := t.trainsConfModel.GetTrainDetail(ctx, req.TrainId)
	if err != nil {
		log.Println(ctx, tag, "DelTrain", "GetTrainDetail err:%v", err)
		return nil, err
	}
	if trainMessage == nil {
		return nil, errors.New("the train message is empty")
	}
	// 删除 trains 表中的记录 ToDo 后续也要删除缓存
	err = t.trainsConfModel.DelTrain(ctx, req.TrainId)
	if err != nil {
		log.Println(ctx, tag, "DelTrain", "DelTrain err:%v", err)
		return nil, err
	}
	// 删除 seats 表中的记录 ToDo 这里后续应该需要删 Elasticsearch 中的记录 毕竟数据量太大 都存入缓存不现实
	err = t.seatsConfModel.DelSeat(ctx, req.TrainId)
	if err != nil {
		log.Println(ctx, tag, "DelTrain", "DelSeat err:%v", err)
		return nil, err
	}
	resp := &trainsRpc.DelTrainResp{Success: 1}
	return resp, nil
}
