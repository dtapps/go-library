package wechatopen

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// ServeHttpAuthorizerAppid 授权跳转
func (c *Client) ServeHttpAuthorizerAppid(ctx context.Context, w http.ResponseWriter, r *http.Request, componentAccessToken string) (resp CgiBinComponentApiQueryAuthResponse, agentUserId string, err error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "ServeHttpAuthorizerAppid")
	defer span.End()

	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = query.Get("agent_user_id")

	span.SetAttributes(attribute.String("http.query.auth_code", authCode))
	span.SetAttributes(attribute.String("http.query.expires_in", expiresIn))
	span.SetAttributes(attribute.String("http.query.agent_user_id", agentUserId))

	if authCode == "" {
		err = errors.New("找不到授权码参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	if expiresIn == "" {
		err = errors.New("找不到过期时间参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	info, err := c.CgiBinComponentApiQueryAuth(ctx, componentAccessToken, authCode)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}
	if info.Result.AuthorizationInfo.AuthorizerAppid == "" {
		err = errors.New("获取失败")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	return info.Result, agentUserId, nil
}
