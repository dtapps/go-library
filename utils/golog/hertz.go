package golog

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertz_requestid "go.dtapp.net/library/contrib/hertz-requestid"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"strings"
	"time"
)

// HertzLogFunc Hertz框架日志函数
type HertzLogFunc func(ctx context.Context, response *HertzLogData)

// HertzLog 框架日志
type HertzLog struct {
	hertzLogFunc HertzLogFunc // Hertz框架日志函数
}

// HertzLogFun *HertzLog 框架日志驱动
type HertzLogFun func() *HertzLog

// NewHertzLog 创建框架实例化
func NewHertzLog(ctx context.Context) (*HertzLog, error) {
	hg := &HertzLog{}

	return hg, nil
}

// Middleware 中间件
func (hg *HertzLog) Middleware() app.HandlerFunc {
	return func(c context.Context, h *app.RequestContext) {

		// 开始时间
		start := time.Now().UTC()

		// 模型
		var log = HertzLogData{}

		// 请求时间
		log.RequestTime = gotime.Current().Time

		// 处理请求
		h.Next(c)

		// 结束时间
		end := time.Now().UTC()

		// 请求消耗时长
		log.RequestCostTime = end.Sub(start).Milliseconds()

		// 响应时间
		log.ResponseTime = gotime.Current().Time

		// 输出路由日志
		hlog.CtxTracef(c, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
			h.Response.StatusCode(),
			log.RequestCostTime,
			h.Request.Header.Method(),
			h.Request.URI().PathOriginal(),
			h.ClientIP(),
			h.Request.Host(),
		)

		// 请求编号
		log.RequestID = hertz_requestid.Get(h)
		if log.RequestID == "" {
			log.RequestID = hertz_requestid.GetX(h)
			if log.RequestID == "" {
				log.RequestID = gorequest.GetRequestIDContext(c)
			}
		}

		// 请求主机
		log.RequestHost = string(h.Request.Host())

		// 请求地址
		log.RequestPath = string(h.Request.URI().Path())

		// 请求参数
		log.RequestQuery = gorequest.ParseQueryString(string(h.Request.QueryString()))

		// 请求方式
		log.RequestMethod = string(h.Request.Header.Method())

		if strings.Contains(string(h.ContentType()), consts.MIMEApplicationHTMLForm) {
			log.RequestBody = gorequest.ParseQueryString(string(h.Request.Body()))
		} else if strings.Contains(string(h.ContentType()), consts.MIMEMultipartPOSTForm) {
			//log.RequestBody = h.Request.Body()
			log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
		} else if strings.Contains(string(h.ContentType()), consts.MIMEApplicationJSON) {
			log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
		} else {
			//log.RequestBody = string(h.Request.Body())
			log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
		}

		// 请求IP
		log.RequestIP = h.ClientIP()

		// 请求头
		requestHeader := make(map[string][]string)
		h.Request.Header.VisitAll(func(k, v []byte) {
			requestHeader[string(k)] = append(requestHeader[string(k)], string(v))
		})
		log.RequestHeader = requestHeader

		// 响应头
		responseHeader := make(map[string][]string)
		h.Response.Header.VisitAll(func(k, v []byte) {
			responseHeader[string(k)] = append(responseHeader[string(k)], string(v))
		})
		log.ResponseHeader = responseHeader

		// 响应状态
		log.ResponseCode = h.Response.StatusCode()

		// 响应内容
		if gorequest.IsValidJSON(string(h.Response.Body())) {
			_ = json.Unmarshal(h.Response.Body(), &log.ResponseBody)
			//log.ResponseBody = gojson.JsonDecodeNoError(string(h.Response.Body()))
		} else {
			//log.ResponseBody = string(h.Response.Body())
		}

		// 调用Hertz框架日志函数
		if hg.hertzLogFunc != nil {
			hg.hertzLogFunc(c, &log)
		}

	}
}

// SetLogFunc 设置日志记录方法
func (hg *HertzLog) SetLogFunc(hertzLogFunc HertzLogFunc) {
	hg.hertzLogFunc = hertzLogFunc
}
