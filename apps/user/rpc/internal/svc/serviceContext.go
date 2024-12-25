package svc

import (
	"database/sql"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/model"
	"gitcode.com/HuaweiCloudDeveloper/OpenSourceForHuaweiDemoGo/apps/user/rpc/internal/config"
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	GaussDBConn *sql.DB
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("opengauss", c.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
