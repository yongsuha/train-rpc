package service

import (
	"context"
	"errors"
	ordersRpc "github.com/yongsuha/train-proto/order"
	"github.com/yongsuha/train-rpc/source/biz"
	"github.com/yongsuha/train-rpc/source/data/mysql"
	"log"
	"time"
)

type OrdersService struct {
	ordersConfModel biz.OrdersConfRepo
}

func NewOrdersService() *OrdersService {
	return &OrdersService{
		ordersConfModel: mysql.NewOrdersConfModel(),
	}
}

func (o *OrdersService) AddOrder(ctx context.Context, req *ordersRpc.AddOrderReq) (*ordersRpc.AddOrderResp, error) {
	tag := "OrdersService:AddOrder"
	if req.UserId == 0 {
		return nil, errors.New("UserId is empty")
	}
	if req.OrderStatus == "" {
		return nil, errors.New("OrderStatus is empty")
	}
	if req.TotalPrice == 0 {
		return nil, errors.New("TotalPrice is empty")
	}
	orderInfo := &biz.OrdersConf{
		UserId:      req.UserId,
		TotalPrice:  req.TotalPrice,
		OrderStatus: req.OrderStatus,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	id, err := o.ordersConfModel.AddOrder(ctx, orderInfo)
	if err != nil {
		log.Println(ctx, tag, "AddOrder", "AddOrder err:%v", err)
		return nil, err
	}
	resp := &ordersRpc.AddOrderResp{OrderId: id}
	return resp, nil
}

func (o *OrdersService) GetOrderDetail(ctx context.Context, req *ordersRpc.GetOrderDetailReq) (*ordersRpc.GetOrderDetailResp, error) {
	tag := "OrdersService:GetOrderDetail"
	orderParam := false
	if req.OrderId != 0 {
		orderParam = true
	}
	if req.UserId != 0 {
		orderParam = true
	}
	if orderParam == false {
		return nil, errors.New("orderParam is empty")
	}
	getOrderParam := &biz.GetOrderDetailParam{
		OrderId: req.OrderId,
		UserId:  req.UserId,
	}
	orders, err := o.ordersConfModel.GetOrderDetail(ctx, getOrderParam)
	if err != nil {
		log.Println(ctx, tag, "GetOrderDetail", "GetOrderDetail err:%v", err)
		return nil, err
	}
	if orders == nil {
		return nil, errors.New("orders is empty")
	}
	orderList := []*ordersRpc.OrderInfo{}
	for _, order := range orders {
		orderInfo := &ordersRpc.OrderInfo{
			OrderId:     order.OrderId,
			UserId:      order.UserId,
			TotalPrice:  order.TotalPrice,
			OrderStatus: order.OrderStatus,
			CreateTime:  order.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  order.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		orderList = append(orderList, orderInfo)
	}
	resp := &ordersRpc.GetOrderDetailResp{OrderList: orderList}
	return resp, nil
}

func (o *OrdersService) DelOrder(ctx context.Context, req *ordersRpc.DelOrderReq) (*ordersRpc.DelOrderResp, error) {
	tag := "OrdersService:DelOrder"
	if req.OrderId == 0 {
		return nil, errors.New("OrderId is empty")
	}
	getOrderParam := &biz.GetOrderDetailParam{
		OrderId: req.OrderId,
	}
	orderDetail, err := o.ordersConfModel.GetOrderDetail(ctx, getOrderParam)
	if err != nil {
		log.Println(ctx, tag, "DelOrder", "GetOrderDetail err:%v", err)
		return nil, err
	}
	if orderDetail == nil {
		return nil, errors.New("orderDetail is empty")
	}
	err = o.ordersConfModel.DelOrder(ctx, req.OrderId)
	if err != nil {
		log.Println(ctx, tag, "DelOrder", "DelOrder err:%v", err)
		return nil, err
	}
	resp := &ordersRpc.DelOrderResp{Success: 1}
	return resp, nil
}

func (o *OrdersService) UpdateOrder(ctx context.Context, req *ordersRpc.UpdateOrderReq) (*ordersRpc.UpdateOrderResp, error) {
	tag := "OrdersService:UpdateOrder"
	if req.OrderId == 0 {
		return nil, errors.New("OrderId is empty")
	}
	if req.UserId == 0 {
		return nil, errors.New("UserId is empty")
	}
	if req.OrderStatus == "" {
		return nil, errors.New("OrderStatus is empty")
	}
	if req.TotalPrice == 0 {
		return nil, errors.New("TotalPrice is empty")
	}
	getOrderParam := &biz.GetOrderDetailParam{
		OrderId: req.OrderId,
		UserId:  req.UserId,
	}
	orderDetail, err := o.ordersConfModel.GetOrderDetail(ctx, getOrderParam)
	if err != nil {
		log.Println(ctx, tag, "UpdateOrder", "GetOrderDetail err:%v", err)
		return nil, err
	}
	if orderDetail == nil {
		return nil, errors.New("orderDetail is empty")
	}
	order := orderDetail[0]
	orderInfo := &biz.OrdersConf{
		OrderId:     req.OrderId,
		UserId:      req.UserId,
		TotalPrice:  req.TotalPrice,
		OrderStatus: req.OrderStatus,
		CreateTime:  order.CreateTime,
		UpdateTime:  time.Now(),
	}
	err = o.ordersConfModel.UpdateOrder(ctx, orderInfo)
	if err != nil {
		log.Println(ctx, tag, "UpdateOrder", "UpdateOrder err:%v", err)
		return nil, err
	}
	resp := &ordersRpc.UpdateOrderResp{
		OrderStatus: req.OrderStatus,
		UserId:      req.UserId,
		OrderId:     req.OrderId,
		TotalPrice:  req.TotalPrice,
		CreateTime:  orderInfo.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  orderInfo.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}
