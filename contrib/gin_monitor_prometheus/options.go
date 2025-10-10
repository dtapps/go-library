package gin_monitor_prometheus

import (
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// Option opts for monitor prometheus
type Option interface {
	apply(cfg *config)
}

type option func(cfg *config)

func (fn option) apply(cfg *config) {
	fn(cfg)
}

type config struct {
	enableGoCollector   bool
	registry            *prom.Registry
	runtimeMetricRules  []collectors.GoRuntimeMetricsRule
	disableServer       bool
	useDefaultServerMux bool
}

func defaultConfig() *config {
	return &config{
		enableGoCollector:   false,
		registry:            prom.NewRegistry(),
		disableServer:       false,
		useDefaultServerMux: true,
	}
}

// WithEnableGoCollector 启用go收集器
func WithEnableGoCollector(enable bool) Option {
	return option(func(cfg *config) {
		cfg.enableGoCollector = enable
	})
}

// WithGoCollectorRule 定义自定义go收集器规则
func WithGoCollectorRule(rules ...collectors.GoRuntimeMetricsRule) Option {
	return option(func(cfg *config) {
		cfg.runtimeMetricRules = rules
	})
}

// WithDisableServer 禁用 Prometheus 服务
func WithDisableServer(disable bool) Option {
	return option(func(cfg *config) {
		cfg.disableServer = disable
	})
}

// WithRegistry define your custom registry
func WithRegistry(registry *prom.Registry) Option {
	return option(func(cfg *config) {
		if registry != nil {
			cfg.registry = registry
		}
	})
}

// WithDefaultServerMux use http.DefaultServeMux
func WithDefaultServerMux(useDefault bool) Option {
	return option(func(cfg *config) {
		cfg.useDefaultServerMux = useDefault
	})
}
