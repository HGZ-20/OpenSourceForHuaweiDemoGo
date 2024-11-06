package logic

import (
	"context"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/model"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListUser) (resp *types.ListUserResponse, err error) {
	// todo: add your logic here and delete this line
	where := map[string]interface{}{}
	if req.Name != "" {
		where["name"] = req.Name
	}
	if req.Email != "" {
		where["email"] = req.Email
	}

	users := make([]model.User, 0)
	if err := l.svcCtx.DB.Limit(10).Where(where).Find(&users).Error; err != nil {
		return nil, err
	}

	listUsers := make([]*types.UpdateUser, 0)
	for _, user := range users {
		listUsers = append(listUsers, &types.UpdateUser{
			Id:   int(user.ID),
			Name: user.Name,
			// Email:    *user.Email,
			Age: int(user.Age),
			// Birthday: user.Birthday.String(),
		})
	}

	resp = new(types.ListUserResponse)
	resp.Message = "success"
	resp.Data = listUsers
	return
}
