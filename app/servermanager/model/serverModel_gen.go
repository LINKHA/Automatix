// Code generated by goctl. DO NOT EDIT.

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	serverFieldNames          = builder.RawFieldNames(&Server{})
	serverRows                = strings.Join(serverFieldNames, ",")
	serverRowsExpectAutoSet   = strings.Join(stringx.Remove(serverFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	serverRowsWithPlaceHolder = strings.Join(stringx.Remove(serverFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAmxServermanagerServerIdPrefix = "cache:amxServermanager:server:id:"
)

type (
	serverModel interface {
		Insert(ctx context.Context, data *Server) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Server, error)
		Update(ctx context.Context, data *Server) error
		Delete(ctx context.Context, id int64) error
	}

	defaultServerModel struct {
		sqlc.CachedConn
		table string
	}

	Server struct {
		Id            int64     `db:"id"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
		DeleteTime    time.Time `db:"delete_time"`
		DelState      int64     `db:"del_state"`
		Name          string    `db:"name"`           // 服务器名称
		ServerType    int64     `db:"server_type"`    // 服务类型(0:运营 1:测试)
		Switch        int64     `db:"switch"`         // 服务器是否开启(0:关闭 1:开启)
		StartTime     int64     `db:"start_time"`     // 开服时间
		Area          string    `db:"area"`           // 服务器地区
		Tags          string    `db:"tags"`           // 标签列表(使用逗号分隔)
		MaxOnline     int64     `db:"max_online"`     // 最大在线人数
		MaxQueue      int64     `db:"max_queue"`      // 最大排队人数
		MaxSign       int64     `db:"max_sign"`       // 最大注册人数
		TemplateValue string    `db:"template_value"` // 自定义参数(json格式)
	}
)

func newServerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultServerModel {
	return &defaultServerModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`server`",
	}
}

func (m *defaultServerModel) Delete(ctx context.Context, id int64) error {
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, amxServermanagerServerIdKey)
	return err
}

func (m *defaultServerModel) FindOne(ctx context.Context, id int64) (*Server, error) {
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, id)
	var resp Server
	err := m.QueryRowCtx(ctx, &resp, amxServermanagerServerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", serverRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultServerModel) Insert(ctx context.Context, data *Server) (sql.Result, error) {
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, serverRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.DeleteTime, data.DelState, data.Name, data.ServerType, data.Switch, data.StartTime, data.Area, data.Tags, data.MaxOnline, data.MaxQueue, data.MaxSign, data.TemplateValue)
	}, amxServermanagerServerIdKey)
	return ret, err
}

func (m *defaultServerModel) Update(ctx context.Context, data *Server) error {
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, serverRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Name, data.ServerType, data.Switch, data.StartTime, data.Area, data.Tags, data.MaxOnline, data.MaxQueue, data.MaxSign, data.TemplateValue, data.Id)
	}, amxServermanagerServerIdKey)
	return err
}

func (m *defaultServerModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, primary)
}

func (m *defaultServerModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", serverRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultServerModel) tableName() string {
	return m.table
}
