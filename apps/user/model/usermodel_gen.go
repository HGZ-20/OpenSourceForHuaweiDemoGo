// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{}, true)
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	userRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(userFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		FindOneByName(ctx context.Context, name sql.NullString) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id       int64          `db:"id"`
		Name     sql.NullString `db:"name"`
		Password sql.NullString `db:"password"`
		Mobile   string         `db:"mobile"`
		Gender   string         `db:"gender"`
		CreateAt time.Time      `db:"create_at"`
		UpdateAt time.Time      `db:"update_at"`
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: `"public"."user"`,
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where mobile = $1 limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByName(ctx context.Context, name sql.NullString) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where name = $1 limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, userRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Password, data.Mobile, data.Gender)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, userRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Id, newData.Name, newData.Password, newData.Mobile, newData.Gender)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
