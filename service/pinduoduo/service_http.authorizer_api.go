package pinduoduo

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// ServeHttpAuthorizerApi 授权跳转
func (c *Client) ServeHttpAuthorizerApi(ctx context.Context, w http.ResponseWriter, r *http.Request) (PddDdkPopAuthTokenCreateResponse, string, string, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "ServeHttpAuthorizerApi")
	defer span.End()

	var (
		query = r.URL.Query()

		code  = query.Get("code")
		state = query.Get("state")
	)

	span.SetAttributes(attribute.String("http.query.code", code))
	span.SetAttributes(attribute.String("http.query.state", state))

	if code == "" {
		err := errors.New("找不到授权码参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return PddDdkPopAuthTokenCreateResponse{}, code, state, err
	}

	response, err := c.PopAuthTokenCreate(ctx, code)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return response.Result, code, state, err
	}

	return response.Result, code, state, nil
}
