package golog

import (
	"context"
	"errors"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

// ApiGorm 接口日志
type ApiGorm struct {
	gormClient *gorm.DB // 数据库驱动
	gormConfig struct {
		stats     bool   // 状态
		tableName string // 表名
	}
}

// ApiGormFun 接口日志驱动
type ApiGormFun func() *ApiGorm

// NewApiGorm 创建接口实例化
func NewApiGorm(ctx context.Context, gormClient *gorm.DB, gormTableName string) (*ApiGorm, error) {
	gl := &ApiGorm{}

	if gormClient == nil {
		gl.gormConfig.stats = false
	} else {

		gl.gormClient = gormClient

		if gormTableName == "" {
			return nil, errors.New("没有设置表名")
		} else {
			gl.gormConfig.tableName = gormTableName
		}

		gl.gormConfig.stats = true

	}

	return gl, nil
}

// Middleware 中间件
func (ag *ApiGorm) Middleware(ctx context.Context, request gorequest.Response) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "api")
	defer span.End()

	if ag.gormConfig.stats {
		ag.gormMiddleware(ctx, span, request)
	}

}

// MiddlewareXml 中间件
func (ag *ApiGorm) MiddlewareXml(ctx context.Context, request gorequest.Response) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "api.xml")
	defer span.End()

	if ag.gormConfig.stats {
		ag.gormMiddlewareXml(ctx, span, request)
	}

}

// MiddlewareCustom 中间件
func (ag *ApiGorm) MiddlewareCustom(ctx context.Context, api string, request gorequest.Response) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "api.custom")
	defer span.End()

	if ag.gormConfig.stats {
		ag.gormMiddlewareCustom(ctx, span, api, request)
	}

}
