package gin_monitor_prometheus

import (
	"github.com/gin-gonic/gin"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	labelMethod     = "method"     // 请求方法（GET, POST 等）
	labelStatusCode = "statusCode" // HTTP 响应状态码
	labelPath       = "path"       // 请求路径
	labelIP         = "ip"         // 客户端 IP 地址

	unknownLabelValue = "unknown" // 默认的标签值，表示未知
)

// genLabels make labels values.
func genLabels(c *gin.Context) prom.Labels {
	labels := make(prom.Labels)
	// 默认值处理：如果标签为空，则使用默认值
	labels[labelMethod] = defaultValIfEmpty(c.Request.Method, unknownLabelValue)
	labels[labelStatusCode] = defaultValIfEmpty(strconv.Itoa(c.Writer.Status()), unknownLabelValue)
	labels[labelPath] = defaultValIfEmpty(c.FullPath(), unknownLabelValue)
	labels[labelIP] = defaultValIfEmpty(c.ClientIP(), unknownLabelValue)

	return labels
}

// ServerTracer 用于监控服务器请求数据，记录请求数量、时延和 IP 请求数量
type ServerTracer struct {
	serverHandledCounter   *prom.CounterVec   // 请求量计数器
	serverHandledHistogram *prom.HistogramVec // 请求时延直方图
	serverIPRequestCounter *prom.CounterVec   // 按 IP 请求计数器
}

// Middleware 中间件用于跟踪请求信息
func (s *ServerTracer) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 计算请求处理耗时
		cost := time.Since(start)
		// 生成标签
		labels := genLabels(c)
		if err := counterAdd(s.serverHandledCounter, 1, labels); err != nil {
			log.Printf("增加请求量计数时发生错误: %v", err)
		}
		if err := histogramObserve(s.serverHandledHistogram, cost, labels); err != nil {
			log.Printf("记录时延数据时发生错误: %v", err)
		}
		if err := counterAdd(s.serverIPRequestCounter, 1, labels); err != nil {
			log.Printf("增加IP请求计数时发生错误: %v", err)
		}
	}
}

// NewServerTracer 创建一个新的 ServerTracer，提供 Prometheus 监控功能
func NewServerTracer(addr, path string, opts ...Option) *ServerTracer {
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
				log.Fatal("GIN: 无法启动 promhttp 服务器，错误：" + err.Error())
			}
		}()
	}

	// 创建并注册请求量计数器
	serverHandledCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "gin_server_throughput",
			Help: "服务器完成的 HTTP 请求总数，不论成功或失败。",
		},
		[]string{labelIP, labelMethod, labelStatusCode, labelPath}, // 标签：IP、方法、状态码、路径
	)
	cfg.registry.MustRegister(serverHandledCounter)

	// 创建并注册请求时延监控
	serverHandledHistogram := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "gin_server_latency_us",
			Help:    "服务器处理请求的时延（微秒）。",
			Buckets: cfg.buckets, // 时延桶配置
		},
		[]string{labelIP, labelMethod, labelStatusCode, labelPath}, // 标签：IP、方法、状态码、路径
	)
	cfg.registry.MustRegister(serverHandledHistogram)

	// 创建并注册按 IP 请求计数器
	serverIPRequestCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "gin_server_ip_requests",
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
	return &ServerTracer{
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
