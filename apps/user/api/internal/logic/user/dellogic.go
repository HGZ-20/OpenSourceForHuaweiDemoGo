package user

import (
	"context"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户信息
func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	delResp, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteUserResp{
		Status: delResp.Status,
	}

	return resp, nil
}
