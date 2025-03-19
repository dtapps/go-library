package hertz_monitor_prometheus

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	labelMethod     = "method"      // 请求方法（GET, POST 等）
	labelStatusCode = "status_code" // HTTP 响应状态码
	labelPath       = "path"        // 请求路径
	labelIP         = "ip"          // 客户端 IP 地址

	unknownLabelValue = "unknown" // 默认的标签值，表示未知
)

// genLabels make labels values.
func genLabels(ctx *app.RequestContext) prometheus.Labels {
	labels := make(prometheus.Labels)
	labels[labelMethod] = string(ctx.Request.Method())
	labels[labelStatusCode] = defaultValIfEmpty(strconv.Itoa(ctx.Response.Header.StatusCode()), unknownLabelValue)
	labels[labelPath] = string(ctx.Request.URI().Path())
	labels[labelIP] = defaultValIfEmpty(ctx.ClientIP(), unknownLabelValue)
	return labels
}

// ServerTracer 用于监控服务器请求数据，记录请求数量、时延和 IP 请求数量
type ServerTracer struct {
	serverHandledTotal           *prometheus.CounterVec   // 请求总数（按方法+路径）
	serverHandledIpTotal         *prometheus.CounterVec   // IP 请求（按 IP+方法+路径）
	serverHandledStatusCodeTotal *prometheus.CounterVec   // 状态码统计（仅按状态码）
	serverHandledDurationSeconds *prometheus.HistogramVec // 耗时（按方法+路径）
}

// Middleware 中间件用于跟踪请求信息
func (s *ServerTracer) Middleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 记录请求开始时间
		start := time.Now()
		// 处理请求
		c.Next(ctx)
		// 计算请求处理耗时
		cost := time.Since(start)
		// 生成标签
		labels := genLabels(c)
		if err := counterAdd(s.serverHandledTotal, 1, labels); err != nil {
			log.Printf("增加请求总数指标时发生错误: %v", err)
		}
		if err := counterAdd(s.serverHandledIpTotal, 1, labels); err != nil {
			log.Printf("增加请求IP指标时发生错误: %v", err)
		}
		if err := counterAdd(s.serverHandledStatusCodeTotal, 1, labels); err != nil {
			log.Printf("增加请求状态码指标时发生错误: %v", err)
		}
		if err := histogramObserve(s.serverHandledDurationSeconds, cost, labels); err != nil {
			log.Printf("记录请求耗时指标时发生错误: %v", err)
		}
	}
}

// NewServerTracer 创建一个新的 ServerTracer，提供 Prometheus 监控功能
func NewServerTracer(addr string, path string, prefix string, opts ...Option) *ServerTracer {
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

	serverHandledTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "http_request_total",
			Help: "HTTP 请求总数（按方法和路径统计）",
		},
		[]string{labelMethod, labelPath},
	)

	serverHandledIpTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "http_request_ip_total",
			Help: "客户端 IP 请求统计（按 IP、方法和路径）",
		},
		[]string{labelIP, labelMethod, labelPath},
	)

	serverHandledStatusCodeTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prefix + "http_request_status_code_total",
			Help: "HTTP 状态码统计（仅按状态码）",
		},
		[]string{labelStatusCode},
	)

	serverHandledDurationSeconds := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    prefix + "http_request_duration_seconds",
			Help:    "HTTP 请求耗时（秒）",
			Buckets: prometheus.DefBuckets,
		},
		[]string{labelMethod, labelPath},
	)

	// 注册指标
	cfg.registry.MustRegister(serverHandledTotal)
	cfg.registry.MustRegister(serverHandledIpTotal)
	cfg.registry.MustRegister(serverHandledStatusCodeTotal)
	cfg.registry.MustRegister(serverHandledDurationSeconds)

	// 注册 Go runtime 指标（可选）
	if cfg.enableGoCollector {
		cfg.registry.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(cfg.runtimeMetricRules...)))
	}

	// 返回 ServerTracer 实例，开始监控
	return &ServerTracer{
		serverHandledTotal:           serverHandledTotal,
		serverHandledIpTotal:         serverHandledIpTotal,
		serverHandledStatusCodeTotal: serverHandledStatusCodeTotal,
		serverHandledDurationSeconds: serverHandledDurationSeconds,
	}
}

// defaultValIfEmpty 返回默认值（如果给定值为空）
func defaultValIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
