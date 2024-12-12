package rpc

import (
	"context"
	trainRpc "github.com/yongsuha/train-proto/train"
	"github.com/yongsuha/train-rpc/source/service"
)

type TrainsServer struct {
	trainsCen *service.TrainsService
}

func NewTrainsService() *TrainsServer {
	return &TrainsServer{
		trainsCen: service.NewTrainsService(),
	}
}

func (t *TrainsServer) AddTrain(ctx context.Context, req *trainRpc.AddTrainReq) (*trainRpc.AddTrainResp, error) {
	return t.trainsCen.AddTrain(ctx, req)
}

func (t *TrainsServer) UpdateTrain(ctx context.Context, req *trainRpc.UpdateTrainReq) (*trainRpc.UpdateTrainResp, error) {
	return t.trainsCen.UpdateTrain(ctx, req)
}

func (t *TrainsServer) GetTrainDetail(ctx context.Context, req *trainRpc.GetTrainDetailReq) (*trainRpc.GetTrainDetailResp, error) {
	return t.trainsCen.GetTrainDetail(ctx, req)
}

func (t *TrainsServer) DelTrain(ctx context.Context, req *trainRpc.DelTrainReq) (*trainRpc.DelTrainResp, error) {
	return t.trainsCen.DelTrain(ctx, req)
}
