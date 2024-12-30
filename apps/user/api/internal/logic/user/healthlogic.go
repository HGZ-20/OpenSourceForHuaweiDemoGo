package user

import (
	"context"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"
	"github.com/jinzhu/copier"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health(req *types.HealthReq) (resp *types.HealthResp, err error) {
	healResp, err := l.svcCtx.UserRpc.Ping(l.ctx, &user.Request{
		Ping: req.Ping,
	})
	if err != nil {
		return nil, err
	}
	var res types.HealthResp
	copier.Copy(&res, healResp)

	return &res, nil
}
