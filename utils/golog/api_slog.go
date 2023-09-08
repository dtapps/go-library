package golog

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"os"
	"runtime"
)

type ApiSLogFun func() *ApiSLog

type ApiSLogConfig struct {
	LogPath     string // 日志文件路径
	LogName     string // 日志文件名
	MaxSize     int    // 单位为MB,默认为512MB
	MaxBackups  int    // 保留旧文件的最大个数
	MaxAge      int    // 文件最多保存多少天 0=不删除
	LocalTime   bool   // 采用本地时间
	Compress    bool   // 是否压缩日志
	ShowLine    bool   // 显示代码行
	LogSaveFile bool   // 是否保存到文件
}

type ApiSLog struct {
	config       *ApiSLogConfig
	jsonHandler  *slog.JSONHandler
	logger       *slog.Logger
	systemConfig struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
}

func NewApiSlog(ctx context.Context, config *ApiSLogConfig) *ApiSLog {

	sl := &ApiSLog{config: config}

	opts := slog.HandlerOptions{
		AddSource: sl.config.ShowLine,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format(gotime.DateTimeFormat))
				//return slog.Attr{}
			}
			return a
		},
	}

	// 是否保存到文件
	if sl.config.LogSaveFile {
		lumberjackLogger := lumberjack.Logger{
			Filename:   sl.config.LogPath + sl.config.LogName, // ⽇志⽂件路径
			MaxSize:    sl.config.MaxSize,                     // 单位为MB,默认为512MB
			MaxAge:     sl.config.MaxAge,                      // 文件最多保存多少天
			MaxBackups: sl.config.MaxBackups,                  // 保留旧文件的最大个数
			LocalTime:  sl.config.LocalTime,                   // 采用本地时间
			Compress:   sl.config.Compress,                    // 是否压缩日志
		}
		sl.jsonHandler = slog.NewJSONHandler(&lumberjackLogger, &opts)
		sl.logger = slog.New(sl.jsonHandler)
	} else {
		sl.jsonHandler = slog.NewJSONHandler(os.Stdout, &opts)
		sl.logger = slog.New(sl.jsonHandler)
	}

	sl.setConfig(ctx)

	return sl
}

// Middleware 中间件
func (sl *ApiSLog) Middleware(ctx context.Context, request gorequest.Response) {
	jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
		slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		slog.String("request_time", gotime.SetCurrent(request.RequestTime).Format()),
		slog.String("request_uri", request.RequestUri),
		slog.String("request_url", gourl.UriParse(request.RequestUri).Url),
		slog.String("request_api", gourl.UriParse(request.RequestUri).Path),
		slog.String("request_method", request.RequestMethod),
		slog.String("request_params", gostring.ToString(request.RequestParams)),
		slog.String("request_header", gostring.ToString(request.RequestHeader)),
		slog.String("request_ip", sl.systemConfig.systemOutsideIp),
		slog.String("response_header", gostring.ToString(request.ResponseHeader)),
		slog.Int("response_status_code", request.ResponseStatusCode),
		slog.String("response_body", gostring.ToString(dorm.JsonDecodeNoError(request.ResponseBody))),
		slog.Int64("response_content_length", request.ResponseContentLength),
		slog.String("response_time", gotime.SetCurrent(request.ResponseTime).Format()),
		slog.String("system_host_name", sl.systemConfig.systemHostname),
		slog.String("system_inside_ip", sl.systemConfig.systemInsideIp),
		slog.String("system_os", sl.systemConfig.systemOs),
		slog.String("system_arch", sl.systemConfig.systemKernel),
		slog.String("go_version", sl.systemConfig.goVersion),
		slog.String("sdk_version", sl.systemConfig.sdkVersion),
	})
	logger := slog.New(jsonHandler)
	logger.Info("Middleware")
}

// MiddlewareXml 中间件
func (sl *ApiSLog) MiddlewareXml(ctx context.Context, request gorequest.Response) {
	jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
		slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		slog.String("request_time", gotime.SetCurrent(request.RequestTime).Format()),
		slog.String("request_uri", request.RequestUri),
		slog.String("request_url", gourl.UriParse(request.RequestUri).Url),
		slog.String("request_api", gourl.UriParse(request.RequestUri).Path),
		slog.String("request_method", request.RequestMethod),
		slog.String("request_params", gostring.ToString(request.RequestParams)),
		slog.String("request_header", gostring.ToString(request.RequestHeader)),
		slog.String("request_ip", sl.systemConfig.systemOutsideIp),
		slog.String("response_header", gostring.ToString(request.ResponseHeader)),
		slog.Int("response_status_code", request.ResponseStatusCode),
		slog.String("response_body", gostring.ToString(dorm.XmlDecodeNoError(request.ResponseBody))),
		slog.Int64("response_content_length", request.ResponseContentLength),
		slog.String("response_time", gotime.SetCurrent(request.ResponseTime).Format()),
		slog.String("system_host_name", sl.systemConfig.systemHostname),
		slog.String("system_inside_ip", sl.systemConfig.systemInsideIp),
		slog.String("system_os", sl.systemConfig.systemOs),
		slog.String("system_arch", sl.systemConfig.systemKernel),
		slog.String("go_version", sl.systemConfig.goVersion),
		slog.String("sdk_version", sl.systemConfig.sdkVersion),
	})
	logger := slog.New(jsonHandler)
	logger.Info("MiddlewareXml")
}

// MiddlewareCustom 中间件
func (sl *ApiSLog) MiddlewareCustom(ctx context.Context, api string, request gorequest.Response) {
	jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
		slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		slog.String("request_time", gotime.SetCurrent(request.RequestTime).Format()),
		slog.String("request_uri", request.RequestUri),
		slog.String("request_url", gourl.UriParse(request.RequestUri).Url),
		slog.String("request_api", api),
		slog.String("request_method", request.RequestMethod),
		slog.String("request_params", gostring.ToString(request.RequestParams)),
		slog.String("request_header", gostring.ToString(request.RequestHeader)),
		slog.String("request_ip", sl.systemConfig.systemOutsideIp),
		slog.String("response_header", gostring.ToString(request.ResponseHeader)),
		slog.Int("response_status_code", request.ResponseStatusCode),
		slog.String("response_body", gostring.ToString(dorm.JsonDecodeNoError(request.ResponseBody))),
		slog.Int64("response_content_length", request.ResponseContentLength),
		slog.String("response_time", gotime.SetCurrent(request.ResponseTime).Format()),
		slog.String("system_host_name", sl.systemConfig.systemHostname),
		slog.String("system_inside_ip", sl.systemConfig.systemInsideIp),
		slog.String("system_os", sl.systemConfig.systemOs),
		slog.String("system_arch", sl.systemConfig.systemKernel),
		slog.String("go_version", sl.systemConfig.goVersion),
		slog.String("sdk_version", sl.systemConfig.sdkVersion),
	})
	logger := slog.New(jsonHandler)
	logger.Info("MiddlewareCustom")
}

func (sl *ApiSLog) setConfig(ctx context.Context) {

	info := getSystem()

	sl.systemConfig.systemHostname = info.SystemHostname
	sl.systemConfig.systemOs = info.SystemOs
	sl.systemConfig.systemKernel = info.SystemKernel

	sl.systemConfig.systemInsideIp = goip.GetInsideIp(ctx)

	sl.systemConfig.sdkVersion = go_library.Version()
	sl.systemConfig.goVersion = runtime.Version()

}
