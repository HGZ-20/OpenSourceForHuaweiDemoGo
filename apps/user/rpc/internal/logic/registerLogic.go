package logic

import (
	"context"
	"database/sql"
	"errors"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/model"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/pkg/ctxdata"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/pkg/encrypt"
	"time"

	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/svc"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegister = errors.New("手机号已被注册")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 1.验证用户是否注册，根据手机号
	userEntity, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	if userEntity != nil {
		return nil, ErrPhoneIsRegister
	}

	// 定义用户数据
	userEntity = &model.User{
		Name:   sql.NullString{String: in.Name, Valid: true},
		Mobile: in.Mobile,
		Gender: in.Gender,
	}

	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPassWordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		userEntity.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, string(userEntity.Id))
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		Count:  count,
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
