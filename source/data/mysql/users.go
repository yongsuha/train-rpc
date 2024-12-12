package mysql

import (
	"context"
	//"../../../../../myTest/dbx"
	"github.com/yongsuha/train-rpc/source/biz"
)

const UsersConfTableName = "users"

type UsersConfModel struct{}

func NewUsersConfModel() biz.UsersConfRepo {
	return &UsersConfModel{}
}

func (u *UsersConfModel) AddUser(ctx context.Context, user *biz.UsersConf) (int64, error) {
	id, err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(UsersConfTableName).Create(user).Error
	return id, err
}

func (u *UsersConfModel) DelUser(ctx context.Context, userId int64) error {
	return dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(UsersConfTableName).Where("user_id = ?", userId).Delete(&biz.SeatsConf{}).Error
}

func (u *UsersConfModel) UpdateUser(ctx context.Context, user *biz.UsersConf) error {
	err := dbx.GetDBWithContext(ctx, dbx.MasterDB).WithContext(ctx).Table(UsersConfTableName).Where("user_id = ?", user.UserId).Updates(user).Error
	return err
}

func (u *UsersConfModel) GetUserDetail(ctx context.Context, req *biz.GetUserDetailParam) (*biz.UsersConf, error) {
	db := dbx.GetDBWithContext(ctx, dbx.SlaveDB).WithContext(ctx).Table(UsersConfTableName)

	if req.UserId > 0 {
		db = db.Where("user_id = ?", req.UserId)
	}
	if req.UserName != "" {
		db = db.Where("user_name = ?", req.UserName)
	}
	if req.Email != "" {
		db = db.Where("email in ?", req.Email)
	}
	if req.PhoneNumber != "" {
		db = db.Where("phone_number > ?", req.PhoneNumber)
	}
	resp := &biz.UsersConf{}
	err := dbx.First(&resp).Error
	return resp, err
}
