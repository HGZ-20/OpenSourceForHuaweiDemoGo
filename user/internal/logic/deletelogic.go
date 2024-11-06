package logic

import (
	"context"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/model"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DelUser) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.DB.Delete(&model.User{}, req.Id).Error; err != nil {
		return nil, err
	}

	resp = new(types.Response)
	resp.Message = "success"
	return
}
