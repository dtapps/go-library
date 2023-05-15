package golog

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

// ApiZapLogFun *ApiClient 驱动
type ApiZapLogFun func() *ApiZapLog

type ApiZapLogConfig struct {
	LogPath    string // 日志文件路径
	LogName    string // 日志文件名
	MaxSize    int    // 单位为MB,默认为512MB
	MaxBackups int    // 保留旧文件的最大个数
	MaxAge     int    // 文件最多保存多少天 0=不删除
	LocalTime  bool   // 采用本地时间
	Compress   bool   // 是否压缩日志
	ShowLine   bool   // 显示代码行
}

type ApiZapLog struct {
	config       *ApiZapLogConfig
	logger       *zap.Logger
	zapCore      zapcore.Core
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

func NewApiZapLog(ctx context.Context, config *ApiZapLogConfig) *ApiZapLog {

	zl := &ApiZapLog{config: config}

	var syncer zapcore.WriteSyncer

	// 定义日志切割配置
	hook := lumberjack.Logger{
		Filename:   zl.config.LogPath + zl.config.LogName, // ⽇志⽂件路径
		MaxSize:    zl.config.MaxSize,                     // 单位为MB,默认为512MB
		MaxBackups: zl.config.MaxBackups,                  // 保留旧文件的最大个数
		LocalTime:  zl.config.LocalTime,                   // 采用本地时间
		Compress:   zl.config.Compress,                    // 是否压缩日志
	}
	if zl.config.MaxAge > 0 {
		// 文件最多保存多少天
		hook.MaxAge = zl.config.MaxAge
	}
	// 在文件输出日志
	syncer = zapcore.AddSync(&hook)

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(gotime.DateTimeFormat))
	}

	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level.CapitalString())
	}

	// 定义zap配置信息
	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_type",
		MessageKey:     "msg",
		TimeKey:        "zap_time",
		NameKey:        "logger",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,          // 自定义时间格式
		EncodeLevel:    customLevelEncoder,         // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// json格式输出
	zl.zapCore = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
		syncer, zap.NewAtomicLevelAt(zapcore.InfoLevel))

	zl.logger = zl.withShowLine(zap.New(zl.zapCore))

	zl.setConfig(ctx)

	return zl
}

// 判断是否显示代码行号
func (zl *ApiZapLog) withShowLine(logger *zap.Logger) *zap.Logger {
	if zl.config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// Middleware 中间件
func (zl *ApiZapLog) Middleware(ctx context.Context, request gorequest.Response) {
	zl.logger.With(
		zapcore.Field{
			Key:    "trace_id",
			Type:   zapcore.StringType,
			String: gotrace_id.GetTraceIdContext(ctx),
		}, zapcore.Field{
			Key:    "request_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.RequestTime).Format(),
		}, zapcore.Field{
			Key:    "request_uri",
			Type:   zapcore.StringType,
			String: request.RequestUri,
		}, zapcore.Field{
			Key:    "request_url",
			Type:   zapcore.StringType,
			String: gourl.UriParse(request.RequestUri).Url,
		}, zapcore.Field{
			Key:    "request_api",
			Type:   zapcore.StringType,
			String: gourl.UriParse(request.RequestUri).Path,
		}, zapcore.Field{
			Key:    "request_method",
			Type:   zapcore.StringType,
			String: request.RequestMethod,
		}, zapcore.Field{
			Key:       "request_params",
			Type:      zapcore.StringType,
			Interface: request.RequestParams, //
		}, zapcore.Field{
			Key:       "request_header",
			Type:      zapcore.StringType,
			Interface: request.RequestHeader, //
		}, zapcore.Field{
			Key:    "request_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOutsideIp,
		}, zapcore.Field{
			Key:       "response_header",
			Type:      zapcore.StringType,
			Interface: request.ResponseHeader, //
		}, zapcore.Field{
			Key:       "response_status_code",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseStatusCode,
		}, zapcore.Field{
			Key:       "response_body",
			Type:      zapcore.StringType,
			Interface: dorm.JsonDecodeNoError(request.ResponseBody), //
		}, zapcore.Field{
			Key:       "response_content_length",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseContentLength,
		}, zapcore.Field{
			Key:    "response_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.ResponseTime).Format(),
		}, zapcore.Field{
			Key:    "system_host_name",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemHostname,
		}, zapcore.Field{
			Key:    "system_inside_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemInsideIp,
		}, zapcore.Field{
			Key:    "system_os",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOs,
		}, zapcore.Field{
			Key:    "system_arch",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemKernel,
		}, zapcore.Field{
			Key:    "go_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.goVersion,
		}, zapcore.Field{
			Key:    "sdk_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.sdkVersion,
		}).
		Info("Middleware")
}

// MiddlewareXml 中间件
func (zl *ApiZapLog) MiddlewareXml(ctx context.Context, request gorequest.Response) {
	zl.logger.With(
		zapcore.Field{
			Key:    "trace_id",
			Type:   zapcore.StringType,
			String: gotrace_id.GetTraceIdContext(ctx),
		}, zapcore.Field{
			Key:    "request_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.RequestTime).Format(),
		}, zapcore.Field{
			Key:    "request_uri",
			Type:   zapcore.StringType,
			String: request.RequestUri,
		}, zapcore.Field{
			Key:    "request_url",
			Type:   zapcore.StringType,
			String: gourl.UriParse(request.RequestUri).Url,
		}, zapcore.Field{
			Key:    "request_api",
			Type:   zapcore.StringType,
			String: gourl.UriParse(request.RequestUri).Path,
		}, zapcore.Field{
			Key:    "request_method",
			Type:   zapcore.StringType,
			String: request.RequestMethod,
		}, zapcore.Field{
			Key:       "request_params",
			Type:      zapcore.StringType,
			Interface: request.RequestParams, //
		}, zapcore.Field{
			Key:       "request_header",
			Type:      zapcore.StringType,
			Interface: request.RequestHeader, //
		}, zapcore.Field{
			Key:    "request_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOutsideIp,
		}, zapcore.Field{
			Key:       "response_header",
			Type:      zapcore.StringType,
			Interface: request.ResponseHeader, //
		}, zapcore.Field{
			Key:       "response_status_code",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseStatusCode,
		}, zapcore.Field{
			Key:       "response_body",
			Type:      zapcore.StringType,
			Interface: dorm.XmlDecodeNoError(request.ResponseBody), //
		}, zapcore.Field{
			Key:       "response_content_length",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseContentLength,
		}, zapcore.Field{
			Key:    "response_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.ResponseTime).Format(),
		}, zapcore.Field{
			Key:    "system_host_name",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemHostname,
		}, zapcore.Field{
			Key:    "system_inside_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemInsideIp,
		}, zapcore.Field{
			Key:    "system_os",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOs,
		}, zapcore.Field{
			Key:    "system_arch",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemKernel,
		}, zapcore.Field{
			Key:    "go_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.goVersion,
		}, zapcore.Field{
			Key:    "sdk_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.sdkVersion,
		}).
		Info("MiddlewareXml")
}

// MiddlewareCustom 中间件
func (zl *ApiZapLog) MiddlewareCustom(ctx context.Context, api string, request gorequest.Response) {
	zl.logger.With(
		zapcore.Field{
			Key:    "trace_id",
			Type:   zapcore.StringType,
			String: gotrace_id.GetTraceIdContext(ctx),
		}, zapcore.Field{
			Key:    "request_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.RequestTime).Format(),
		}, zapcore.Field{
			Key:    "request_uri",
			Type:   zapcore.StringType,
			String: request.RequestUri,
		}, zapcore.Field{
			Key:    "request_url",
			Type:   zapcore.StringType,
			String: gourl.UriParse(request.RequestUri).Url,
		}, zapcore.Field{
			Key:    "request_api",
			Type:   zapcore.StringType,
			String: api,
		}, zapcore.Field{
			Key:    "request_method",
			Type:   zapcore.StringType,
			String: request.RequestMethod,
		}, zapcore.Field{
			Key:       "request_params",
			Type:      zapcore.StringType,
			Interface: request.RequestParams, //
		}, zapcore.Field{
			Key:       "request_header",
			Type:      zapcore.StringType,
			Interface: request.RequestHeader, //
		}, zapcore.Field{
			Key:    "request_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOutsideIp,
		}, zapcore.Field{
			Key:       "response_header",
			Type:      zapcore.StringType,
			Interface: request.ResponseHeader, //
		}, zapcore.Field{
			Key:       "response_status_code",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseStatusCode,
		}, zapcore.Field{
			Key:       "response_body",
			Type:      zapcore.StringType,
			Interface: dorm.JsonDecodeNoError(request.ResponseBody), //
		}, zapcore.Field{
			Key:       "response_content_length",
			Type:      zapcore.Int64Type,
			Interface: request.ResponseContentLength,
		}, zapcore.Field{
			Key:    "response_time",
			Type:   zapcore.StringType,
			String: gotime.SetCurrent(request.ResponseTime).Format(),
		}, zapcore.Field{
			Key:    "system_host_name",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemHostname,
		}, zapcore.Field{
			Key:    "system_inside_ip",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemInsideIp,
		}, zapcore.Field{
			Key:    "system_os",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemOs,
		}, zapcore.Field{
			Key:    "system_arch",
			Type:   zapcore.StringType,
			String: zl.systemConfig.systemKernel,
		}, zapcore.Field{
			Key:    "go_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.goVersion,
		}, zapcore.Field{
			Key:    "sdk_version",
			Type:   zapcore.StringType,
			String: zl.systemConfig.sdkVersion,
		}).
		Info("MiddlewareCustom")
}

func (zl *ApiZapLog) setConfig(ctx context.Context) {

	info := getSystem()

	zl.systemConfig.systemHostname = info.SystemHostname
	zl.systemConfig.systemOs = info.SystemOs
	zl.systemConfig.systemKernel = info.SystemKernel

	zl.systemConfig.systemInsideIp = goip.GetInsideIp(ctx)

	zl.systemConfig.sdkVersion = go_library.Version()
	zl.systemConfig.goVersion = runtime.Version()
}
