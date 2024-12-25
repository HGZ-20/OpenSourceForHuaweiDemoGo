package logic

import (
	"context"
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
	ErrPasswordWrong    = errors.New("密码错误")
	ErrPhoneNotRegister = errors.New("手机号未注册")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 验证用户是否注册，根据手机号验证
	userEntity, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			return nil, ErrPhoneNotRegister
		}
		return nil, err
	}

	// 密码验证
	if !encrypt.ValidatePassword(in.Password, userEntity.Password.String) {
		return nil, ErrPasswordWrong
	}

	// 生成Token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, string(userEntity.Id))
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
		Id:     userEntity.Id,
		Token:  token,
	}, nil
}
