package mysql

import (
	//"../../../../../myTest/dbx"
	"context"
	"github.com/yongsuha/train-rpc/source/biz"
)

const TicketsConfTableName = "tickets"

type TicketsConfModel struct{}

func NewTicketsConfModel() biz.TicketsConfRepo {
	return &TicketsConfModel{}
}

func (t *TicketsConfModel) AddTicket(ctx context.Context, ticket *biz.TicketsConf) (int64, error) {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TicketsConfTableName).Create(ticket).Error
	return id, err
}

func (t *TicketsConfModel) DelTicket(ctx context.Context, id int64) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TicketsConfTableName).Where("ticket_id = ?", id).Delete(&biz.TicketsConf{}).Error
}

func (t *TicketsConfModel) UpdateTicket(ctx context.Context, ticket *biz.TicketsConf) error {
	err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TicketsConfTableName).Where("ticket_id = ?", ticket.TicketId).Updates(ticket).Error
	return err
}

func (t *TicketsConfModel) GetTicketDetail(ctx context.Context, req *biz.GetTicketDetailParam) ([]*biz.TicketsConf, error) {
	db := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(TicketsConfTableName)

	if req.TicketId > 0 {
		db = db.Where("ticket_id = ?", req.TicketId)
	}
	if req.UserId != 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	var resp []*biz.TicketsConf
	err := db.Find(&resp).Error
	return resp, err
}
