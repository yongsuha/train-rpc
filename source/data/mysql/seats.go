package mysql

import (
	"context"
	//"../../../../../myTest/dbx"
	"github.com/yongsuha/train-rpc/source/biz"
)

const SeatsConfTableName = "seats"

type SeatsConfModel struct{}

func NewSeatsConfModel() biz.SeatsConfRepo {
	return &SeatsConfModel{}
}

func (s *SeatsConfModel) AddSeat(ctx context.Context, train *biz.SeatsConf) (int64, error) {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(SeatsConfTableName).Create(train).Error
	return id, err
}

func (s *SeatsConfModel) DelSeat(ctx context.Context, trainId int64) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(SeatsConfTableName).Where("train_id = ?", id).Delete(&biz.SeatsConf{}).Error
}

func (s *SeatsConfModel) UpdateSeat(ctx context.Context, train *biz.SeatsConf) error {
	err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(SeatsConfTableName).Where("train_id = ?", train.TrainId).Updates(train).Error
	return err
}

func (t *SeatsConfModel) GetSeatDetail(ctx context.Context, getDetailParam *biz.GetSeatDetailParam) (*biz.SeatsConf, error) {
	resp := &biz.SeatsConf{}
	err := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(SeatsConfTableName).Where("train_id = ? and seat_number = ?", getDetailParam.TrainId, getDetailParam.SeatNumber).First(&resp).Error
	return resp, err
}
