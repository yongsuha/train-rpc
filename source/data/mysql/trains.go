package mysql

import (
	//"../../../../../myTest/dbx"
	"context"
	"github.com/yongsuha/train-rpc/source/biz"
)

const TrainsConfTableName = "trains"

type TrainsConfModel struct{}

func NewTrainsConfModel() biz.TrainsConfRepo {
	return &TrainsConfModel{}
}

func (t *TrainsConfModel) AddTrain(ctx context.Context, train *biz.TrainsConf) (int64, error) {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TrainsConfTableName).Create(train).Error
	return id, err
}

func (t *TrainsConfModel) UpdateTrain(ctx context.Context, train *biz.TrainsConf) error {
	err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TrainsConfTableName).Where("train_id = ?", train.Id).Updates(train).Error
	return err
}

func (t *TrainsConfModel) GetTrainDetail(ctx context.Context, id int64) (*biz.TrainsConf, error) {
	resp := &biz.TrainsConf{}
	err := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(TrainsConfTableName).Where("train_id = ?", id).First(&resp).Error
	return resp, err
}

func (t *TrainsConfModel) DelTrain(ctx context.Context, id int64) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(TrainsConfTableName).Where("train_id = ?", id).Delete(&biz.TrainsConf{}).Error
}
