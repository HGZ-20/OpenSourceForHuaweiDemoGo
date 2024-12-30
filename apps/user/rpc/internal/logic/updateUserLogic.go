package logic

import (
	"context"
	"database/sql"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/pkg/encrypt"
	"google.golang.org/grpc/status"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(500, "用户不存在")
	}
	_, err = l.svcCtx.UserModel.FindOneByName(l.ctx, sql.NullString{String: in.Name, Valid: true})
	if err == nil {
		return nil, status.Error(500, "用户名已存在")
	}
	if in.Name != "" {
		userInfo.Name = sql.NullString{String: in.Name, Valid: true}
	}
	if in.Mobile != "" {
		userInfo.Mobile = in.Mobile
	}
	if in.Gender != "" {
		userInfo.Gender = in.Gender
	}
	if in.Password != "" {
		genPassword, err := encrypt.GenPassWordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		userInfo.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &user.UpdateUserResp{
		Status: true,
	}, nil
}
