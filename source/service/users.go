package service

import (
	"context"
	"errors"
	usersRpc "github.com/yongsuha/train-proto/user"
	"github.com/yongsuha/train-rpc/source/biz"
	"github.com/yongsuha/train-rpc/source/data/mysql"
	"log"
	"time"
)

type UsersService struct {
	usersConfModel biz.UsersConfRepo
}

func NewUsersService() *UsersService {
	return &UsersService{
		usersConfModel: mysql.NewUsersConfModel(),
	}
}

func (u *UsersService) AddUser(ctx context.Context, req *usersRpc.AddUserReq) (*usersRpc.AddUserResp, error) {
	tag := "UsersService:AddUser"
	if req.UserName == "" {
		return nil, errors.New("UserName is empty")
	}
	if req.PassWord == "" {
		return nil, errors.New("PassWord is empty")
	}
	userInfo := &biz.UsersConf{
		UserName:    req.UserName,
		PassWord:    req.PassWord,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	id, err := u.usersConfModel.AddUser(ctx, userInfo)
	if err != nil {
		log.Println(ctx, tag, "AddUser", "AddUser err:%v", err)
		return nil, err
	}
	resp := &usersRpc.AddUserResp{UserId: id}
	return resp, nil
}

func (u *UsersService) DelUser(ctx context.Context, req *usersRpc.DelUserReq) (*usersRpc.DelUserResp, error) {
	tag := "UsersService:DelUser"
	if req.UserId == 0 {
		return nil, errors.New("UserId is empty")
	}
	err := u.usersConfModel.DelUser(ctx, req.UserId)
	if err != nil {
		log.Println(ctx, tag, "DelUser", "DelUser err:%v", err)
		return nil, err
	}
	resp := &usersRpc.DelUserResp{Success: 1}
	return resp, nil
}

func (u *UsersService) GetUserDetail(ctx context.Context, req *usersRpc.GetUserDetailReq) (*usersRpc.GetUserDetailResp, error) {
	tag := "UsersService:GetUserDetail"
	userInfo := false
	if req.UserId != 0 {
		userInfo = true
	}
	if req.UserName != "" {
		userInfo = true
	}
	if req.Email != "" {
		userInfo = true
	}
	if req.PhoneNumber != "" {
		userInfo = true
	}
	if userInfo == false {
		return nil, errors.New("userInfo is empty")
	}
	// 填充参数
	param := &biz.GetUserDetailParam{
		UserId:      req.UserId,
		UserName:    req.UserName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	userMessage, err := u.usersConfModel.GetUserDetail(ctx, param)
	if err != nil {
		log.Println(ctx, tag, "GetUserDetail", "GetUserDetail err:%v", err)
		return nil, err
	}
	resp := &usersRpc.GetUserDetailResp{
		UserId:      userMessage.UserId,
		UserName:    userMessage.UserName,
		Email:       userMessage.Email,
		PhoneNumber: userMessage.PhoneNumber,
		PassWord:    userMessage.PassWord,
		CreateTime:  userMessage.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  userMessage.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}

func (u *UsersService) UpdateUser(ctx context.Context, req *usersRpc.UpdateUserReq) (*usersRpc.UpdateUserResp, error) {
	tag := "UsersService:UpdateUser"
	if req.UserId == 0 {
		return nil, errors.New("UserId is empty")
	}
	// 先获取用户信息
	getUserParam := &biz.GetUserDetailParam{
		UserId:      req.UserId,
		UserName:    req.UserName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	userDetail, err := u.usersConfModel.GetUserDetail(ctx, getUserParam)
	if err != nil {
		log.Println(ctx, tag, "UpdateUser", "GetUserDetail err:%v", err)
		return nil, err
	}
	if userDetail == nil {
		return nil, errors.New("userDetail is empty")
	}
	// 不为空再更新
	userInfo := &biz.UsersConf{
		UserId:      req.UserId,
		UserName:    req.UserName,
		PassWord:    req.PassWord,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreateTime:  userDetail.CreateTime,
		UpdateTime:  time.Now(),
	}
	err = u.usersConfModel.UpdateUser(ctx, userInfo)
	if err != nil {
		log.Println(ctx, tag, "UpdateUser", "UpdateUser err:%v", err)
		return nil, err
	}
	resp := &usersRpc.UpdateUserResp{
		UserId:      userInfo.UserId,
		UserName:    userInfo.UserName,
		PassWord:    userInfo.PassWord,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		CreateTime:  userInfo.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  userInfo.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}
