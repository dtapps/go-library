package gojobs

type TaskHelperOption interface {
	apply(cfg *taskHelperConfig)
}

type taskHelperOption func(cfg *taskHelperConfig)

func (fn taskHelperOption) apply(cfg *taskHelperConfig) {
	fn(cfg)
}

type taskHelperConfig struct {
	logIsDebug            bool   // [日志]日志是否启动
	traceIsFilter         bool   // [过滤]链路追踪是否过滤
	traceIsFilterKeyName  string // [过滤]Key名称
	traceIsFilterKeyValue string // [过滤]Key值
}

// defaultTaskHelperConfig 默认配置
func defaultTaskHelperConfig() *taskHelperConfig {
	return &taskHelperConfig{}
}

// newTaskHelperConfig 初始配置
func newTaskHelperConfig(opts []TaskHelperOption) *taskHelperConfig {
	cfg := defaultTaskHelperConfig()
	for _, opt := range opts {
		opt.apply(cfg)
	}
	return cfg
}

// TaskHelperWithDebug 设置日志是否打印
func TaskHelperWithDebug(is bool) TaskHelperOption {
	return taskHelperOption(func(cfg *taskHelperConfig) {
		cfg.logIsDebug = is
	})
}

// TaskHelperWithFilter 设置链路追踪是否过滤
func TaskHelperWithFilter(is bool, keyName string, keyValue string) TaskHelperOption {
	return taskHelperOption(func(cfg *taskHelperConfig) {
		cfg.traceIsFilter = is
		cfg.traceIsFilterKeyName = keyName
		cfg.traceIsFilterKeyValue = keyValue
	})
}
