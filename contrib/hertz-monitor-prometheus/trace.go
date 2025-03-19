package hertz_monitor_prometheus

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/tracer"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

const (
	labelMethod     = "method"     // 请求方法（GET, POST 等）
	labelStatusCode = "statusCode" // HTTP 响应状态码
	labelPath       = "path"       // 请求路径
	labelIP         = "ip"         // 客户端 IP 地址

	unknownLabelValue = "unknown" // 默认的标签值，表示未知
)

// genLabels make labels values.
func genLabels(ctx *app.RequestContext) prom.Labels {
	labels := make(prom.Labels)
	// 默认值处理：如果标签为空，则使用默认值
	labels[labelMethod] = defaultValIfEmpty(string(ctx.Request.Method()), unknownLabelValue)
	labels[labelStatusCode] = defaultValIfEmpty(strconv.Itoa(ctx.Response.Header.StatusCode()), unknownLabelValue)
	labels[labelPath] = defaultValIfEmpty(ctx.FullPath(), unknownLabelValue)
	labels[labelIP] = defaultValIfEmpty(ctx.ClientIP(), unknownLabelValue)

	return labels
}

// serverTracer 用于监控服务器请求数据，记录请求数量、时延和 IP 请求数量
type serverTracer struct {
	serverHandledCounter   *prom.CounterVec   // 请求量计数器
	serverHandledHistogram *prom.HistogramVec // 请求时延直方图
	serverIPRequestCounter *prom.CounterVec   // 按 IP 请求计数器
}

// Start record the beginning of server handling request from client.
func (s *serverTracer) Start(ctx context.Context, c *app.RequestContext) context.Context {
	return ctx
}

// Finish record the ending of server handling request from client.
func (s *serverTracer) Finish(ctx context.Context, c *app.RequestContext) {
	if c.GetTraceInfo().Stats().Level() == stats.LevelDisabled {
		return
	}

	httpStart := c.GetTraceInfo().Stats().GetEvent(stats.HTTPStart)
	httpFinish := c.GetTraceInfo().Stats().GetEvent(stats.HTTPFinish)
	if httpFinish == nil || httpStart == nil {
		return
	}
	// 计算请求处理耗时
	cost := httpFinish.Time().Sub(httpStart.Time())
	// 生成标签
	labels := genLabels(c)
	_ = counterAdd(s.serverHandledCounter, 1, labels)            // 增加请求量计数
	_ = histogramObserve(s.serverHandledHistogram, cost, labels) // 记录时延数据
	_ = counterAdd(s.serverIPRequestCounter, 1, labels)          // 增加 IP 请求计数
}

// NewServerTracer 创建一个新的 ServerTracer，提供 Prometheus 监控功能
func NewServerTracer(addr, path string, opts ...Option) tracer.Tracer {
	cfg := defaultConfig()

	// 处理用户传入的配置选项
	for _, opts := range opts {
		opts.apply(cfg)
	}

	// 启动 Prometheus HTTP 监控接口（如果未禁用）
	if !cfg.disableServer {
		httpServer := http.DefaultServeMux
		if !cfg.useDefaultServerMux {
			httpServer = http.NewServeMux()
		}
		httpServer.Handle(path, promhttp.HandlerFor(cfg.registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}))
		go func() {
			if err := http.ListenAndServe(addr, httpServer); err != nil {
				log.Fatal("HERTZ: 无法启动 promhttp 服务器，错误：" + err.Error())
			}
		}()
	}

	// 创建并注册请求量计数器
	serverHandledCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "hertz_server_throughput",
			Help: "服务器完成的 HTTP 请求总数，不论成功或失败。",
		},
		[]string{labelMethod, labelStatusCode, labelPath}, // 标签：方法、状态码、路径
	)
	cfg.registry.MustRegister(serverHandledCounter)

	// 创建并注册请求时延监控
	serverHandledHistogram := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "hertz_server_latency_us",
			Help:    "服务器处理请求的时延（微秒）。",
			Buckets: cfg.buckets, // 时延桶配置
		},
		[]string{labelMethod, labelStatusCode, labelPath}, // 标签：方法、状态码、路径
	)
	cfg.registry.MustRegister(serverHandledHistogram)

	// 创建并注册按 IP 请求计数器
	serverIPRequestCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "hertz_server_ip_requests",
			Help: "来自每个 IP 地址的 HTTP 请求总数。",
		},
		[]string{labelIP, labelMethod, labelStatusCode, labelPath}, // 标签：IP、方法、状态码、路径
	)
	cfg.registry.MustRegister(serverIPRequestCounter)

	// 注册 Go runtime 指标（可选）
	if cfg.enableGoCollector {
		cfg.registry.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(cfg.runtimeMetricRules...)))
	}

	// 返回 ServerTracer 实例，开始监控
	return &serverTracer{
		serverHandledCounter:   serverHandledCounter,
		serverHandledHistogram: serverHandledHistogram,
		serverIPRequestCounter: serverIPRequestCounter,
	}
}

// defaultValIfEmpty 返回默认值（如果给定值为空）
func defaultValIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
