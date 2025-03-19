package hertz_monitor_prometheus

import (
	prom "github.com/prometheus/client_golang/prometheus"
	"time"
)

// counterAdd 是 prom.Counter 的包装函数，用于增加计数器的值。
// counterVec 是计数器向量，value 是增加的值，labels 是与该计数器相关的标签。
func counterAdd(counterVec *prom.CounterVec, value int, labels prom.Labels) error {
	// 获取与标签匹配的计数器
	counter, err := counterVec.GetMetricWith(labels)
	if err != nil {
		return err // 如果获取计数器失败，返回错误
	}
	// 增加计数器的值
	counter.Add(float64(value))
	return nil
}

// histogramObserve 是 prom.Observer 的包装函数，用于观察时延或值。
// histogramVec 是直方图向量，value 是观察到的值（例如请求时延），labels 是与该直方图相关的标签。
func histogramObserve(histogramVec *prom.HistogramVec, value time.Duration, labels prom.Labels) error {
	// 获取与标签匹配的直方图
	histogram, err := histogramVec.GetMetricWith(labels)
	if err != nil {
		return err // 如果获取直方图失败，返回错误
	}
	// 观察时延值，转换为微秒
	histogram.Observe(float64(value.Microseconds()))
	return nil
}
