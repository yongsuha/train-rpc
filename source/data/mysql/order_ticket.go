package mysql

import (
	"context"
	//"../../../../../myTest/dbx"
	"github.com/yongsuha/train-rpc/source/biz"
)

const OrderTicketConfTableName = "order_ticket"

type OrderTicketConfModel struct{}

func NewOrderTicketConfModel() biz.OrderTicketConfRepo {
	return &OrderTicketConfModel{}
}

func (o *OrderTicketConfModel) AddOrderTicket(ctx context.Context, orderTicket *biz.OrderTicketConf) error {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(OrderTicketConfTableName).Create(orderTicket).Error
	return id, err
}

func (o *OrderTicketConfModel) DelOrderTicket(ctx context.Context, param *biz.OTParam) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(OrderTicketConfTableName).Where("order_id = ? and ticket_id = ?", param.OrderId, param.TicketId).Delete(&biz.OrderTicketConf{}).Error
}

func (o *OrderTicketConfModel) GetOrderTicketDetail(ctx context.Context, orderId int64) ([]*biz.OrderTicketConf, error) {
	var resp []*biz.OrderTicketConf
	err := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(OrderTicketConfTableName).Where("order_id = ?", orderId).Find(&resp).Error
	return resp, err
}
