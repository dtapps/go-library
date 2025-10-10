package hertz_ent_log

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.dtapp.net/library/contrib/hertz_requestid"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

// HertzLogFunc Hertz框架日志函数
type HertzLogFunc func(ctx context.Context, response *HertzLogData)

// HertzLog 框架日志
type HertzLog struct {
	hertzLogFunc      HertzLogFunc // Hertz框架日志函数
	debug             bool         // 是否开启调试模式
	debugPathPrefixes []string     // 仅当 path 匹配这些前缀时才 debug
}

// HertzLogFun *HertzLog 框架日志驱动
type HertzLogFun func() *HertzLog

// NewHertzLog 创建框架实例化
func NewHertzLog(ctx context.Context) (*HertzLog, error) {
	hg := &HertzLog{}

	return hg, nil
}

// SetDebug 设置 debug
func (hg *HertzLog) SetDebug(debug bool) {
	hg.debug = debug
}

// SetDebugPathPrefixes 设置 debug 路径前缀（例如：[]string{"/notify", "/api/debug"}）
func (hg *HertzLog) SetDebugPathPrefixes(prefixes []string) {
	hg.debugPathPrefixes = prefixes
}

// shouldDebug 根据路径判断是否应打印 debug 日志
func (hg *HertzLog) shouldDebug(path string) bool {
	if !hg.debug {
		return false
	}
	if len(hg.debugPathPrefixes) == 0 {
		return true // 无前缀限制，全部打印
	}
	for _, prefix := range hg.debugPathPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
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

		// 请求内容
		lowerCT := strings.ToLower(string(h.ContentType()))
		switch {
		case strings.Contains(lowerCT, MIMEApplicationHTMLForm):
			log.RequestBody = gorequest.ParseQueryString(string(h.Request.Body()))
		case strings.Contains(lowerCT, MIMEMultipartPOSTForm):
			log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
		case strings.Contains(lowerCT, MIMEApplicationJSON):
			log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
		case strings.Contains(lowerCT, MIMEApplicationXML) || strings.Contains(lowerCT, MIMETextXML):
			log.RequestBody = gorequest.XmlDecodeNoError(h.Request.Body())
		default:
			if gorequest.IsValidJSON(string(h.Request.Body())) {
				log.RequestBody = gorequest.JsonDecodeNoError(string(h.Request.Body()))
			}
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

		// 打印
		currentPath := string(h.Request.URI().Path())
		if hg.shouldDebug(currentPath) {
			// 拷贝 body 防止后续 handler 读不到
			bodyCopy := append([]byte(nil), h.Request.Body()...)
			h.Request.SetBody(bodyCopy)

			// 路由参数
			pathParams := make(map[string]string)
			for _, p := range h.Params {
				pathParams[p.Key] = p.Value
			}

			// 获取 Query 参数
			queryParams := make(map[string]string)
			h.Request.URI().QueryArgs().VisitAll(func(key, value []byte) {
				queryParams[string(key)] = string(value)
			})

			// 获取 Header
			headers := make(map[string]string)
			h.Request.Header.VisitAll(func(key, value []byte) {
				headers[string(key)] = string(value)
			})

			// 获取 Form 数据
			formData := make(map[string]string)
			h.Request.PostArgs().VisitAll(func(key, value []byte) {
				formData[string(key)] = string(value)
			})

			// Body 内容（截断保护）
			bodyStr := string(bodyCopy)
			const MaxBodyLogSize = 2048
			if len(bodyStr) > MaxBodyLogSize {
				bodyStr = bodyStr[:MaxBodyLogSize] + "..."
			}

			// 输出结构化日志
			slog.Debug("hertz request debug",
				slog.String("client_ip", h.ClientIP()),
				slog.String("method", string(h.Request.Header.Method())),
				slog.String("path", string(h.Request.URI().Path())),
				slog.Any("path_params", pathParams),
				slog.Any("query_params", queryParams),
				slog.Any("headers", headers),
				slog.Any("form_data", formData),
				slog.String("body", bodyStr),
			)
		}

	}
}

// SetLogFunc 设置日志记录方法
func (hg *HertzLog) SetLogFunc(hertzLogFunc HertzLogFunc) {
	hg.hertzLogFunc = hertzLogFunc
}
