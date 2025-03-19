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
	labelMethod     = "method"
	labelStatusCode = "statusCode"
	labelPath       = "path"

	unknownLabelValue = "unknown"
)

// genLabels make labels values.
func genLabels(c *gin.Context) prom.Labels {
	labels := make(prom.Labels)
	labels[labelMethod] = defaultValIfEmpty(c.Request.Method, unknownLabelValue)
	labels[labelStatusCode] = defaultValIfEmpty(strconv.Itoa(c.Writer.Status()), unknownLabelValue)
	labels[labelPath] = defaultValIfEmpty(c.FullPath(), unknownLabelValue)

	return labels
}

type ServerTracer struct {
	serverHandledCounter   *prom.CounterVec
	serverHandledHistogram *prom.HistogramVec
}

func (s *ServerTracer) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 耗时
		cost := time.Since(start)

		labels := genLabels(c)
		_ = counterAdd(s.serverHandledCounter, 1, labels)
		_ = histogramObserve(s.serverHandledHistogram, cost, labels)
	}
}

// NewServerTracer provides tracer for server access, addr and path is the scrape_configs for prometheus server.
func NewServerTracer(addr, path string, opts ...Option) *ServerTracer {
	cfg := defaultConfig()

	for _, opts := range opts {
		opts.apply(cfg)
	}

	if !cfg.disableServer {
		httpServer := http.DefaultServeMux
		if !cfg.useDefaultServerMux {
			httpServer = http.NewServeMux()
		}
		httpServer.Handle(path, promhttp.HandlerFor(cfg.registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}))
		go func() {
			if err := http.ListenAndServe(addr, httpServer); err != nil {
				log.Fatal("GIN: Unable to start a promhttp server, err: " + err.Error())
			}
		}()
	}

	serverHandledCounter := prom.NewCounterVec(
		prom.CounterOpts{
			Name: "gin_server_throughput",
			Help: "Total number of HTTPs completed by the server, regardless of success or failure.",
		},
		[]string{labelMethod, labelStatusCode, labelPath},
	)
	cfg.registry.MustRegister(serverHandledCounter)

	serverHandledHistogram := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:    "gin_server_latency_us",
			Help:    "Latency (microseconds) of HTTP that had been application-level handled by the server.",
			Buckets: cfg.buckets,
		},
		[]string{labelMethod, labelStatusCode, labelPath},
	)
	cfg.registry.MustRegister(serverHandledHistogram)

	if cfg.enableGoCollector {
		cfg.registry.MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(cfg.runtimeMetricRules...)))
	}

	return &ServerTracer{
		serverHandledCounter:   serverHandledCounter,
		serverHandledHistogram: serverHandledHistogram,
	}
}

func defaultValIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}
