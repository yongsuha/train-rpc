package mysql

import (
	//"../../../../../myTest/dbx"
	"context"
	"github.com/yongsuha/train-rpc/source/biz"
)

const OrdersConfTableName = "orders"

type OrdersConfModel struct{}

func NewOrdersConfModel() biz.OrdersConfRepo {
	return &OrdersConfModel{}
}

func (o *OrdersConfModel) AddOrder(ctx context.Context, order *biz.OrdersConf) (int64, error) {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(OrdersConfTableName).Create(order).Error
	return id, err
}

func (o *OrdersConfModel) DelOrder(ctx context.Context, id int64) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(OrdersConfTableName).Where("order_id = ?", id).Delete(&biz.OrdersConf{}).Error
}

func (o *OrdersConfModel) UpdateOrder(ctx context.Context, order *biz.OrdersConf) error {
	err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(OrdersConfTableName).Where("order_id = ?", order.OrderId).Updates(order).Error
	return err
}

func (o *OrdersConfModel) GetOrderDetail(ctx context.Context, req *biz.GetOrderDetailParam) ([]*biz.OrdersConf, error) {
	db := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(OrdersConfTableName)

	if req.OrderId > 0 {
		db = db.Where("order_id = ?", req.OrderId)
	}
	if req.UserId != 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	var resp []*biz.OrdersConf
	err := db.Find(&resp).Error
	return resp, err
}
