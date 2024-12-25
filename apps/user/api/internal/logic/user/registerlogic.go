package user

import (
	"context"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"
	"google.golang.org/grpc/status"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		Name:     req.Name,
		Gender:   req.Gender,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	var res types.RegisterResp
	copier.Copy(&res, registerResp)
	logx.Infof("RegisterResp: %+v", registerResp)

	return &res, nil
}
