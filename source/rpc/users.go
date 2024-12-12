package rpc

import (
	"context"
	usersRpc "github.com/yongsuha/train-proto/user"
	"github.com/yongsuha/train-rpc/source/service"
)

type UsersServer struct {
	usersCen *service.UsersService
}

func NewUsersService() *UsersServer {
	return &UsersServer{
		usersCen: service.NewUsersService(),
	}
}

func (u *UsersServer) AddUser(ctx context.Context, req *usersRpc.AddUserReq) (*usersRpc.AddUserResp, error) {
	return u.usersCen.AddUser(ctx, req)
}

func (u *UsersServer) DelUser(ctx context.Context, req *usersRpc.DelUserReq) (*usersRpc.DelUserResp, error) {
	return u.usersCen.DelUser(ctx, req)
}

func (u *UsersServer) GetUserDetail(ctx context.Context, req *usersRpc.GetUserDetailReq) (*usersRpc.GetUserDetailResp, error) {
	return u.usersCen.GetUserDetail(ctx, req)
}

func (u *UsersServer) UpdateUser(ctx context.Context, req *usersRpc.UpdateUserReq) (*usersRpc.UpdateUserResp, error) {
	return u.usersCen.UpdateUser(ctx, req)
}
