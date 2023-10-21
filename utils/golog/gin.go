package golog

import (
	"bytes"
	"context"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// GinClientFun *GinClient 驱动
type GinClientFun func() *GinClient

// GinClient 框架
type GinClient struct {
	ipService *goip.Client // IP服务
	config    struct {
		systemHostname      string  // 主机名
		systemOs            string  // 系统类型
		systemVersion       string  // 系统版本
		systemKernel        string  // 系统内核
		systemKernelVersion string  // 系统内核版本
		systemBootTime      uint64  // 系统开机时间
		cpuCores            int     // CPU核数
		cpuModelName        string  // CPU型号名称
		cpuMhz              float64 // CPU兆赫
		systemInsideIp      string  // 内网ip
		systemOutsideIp     string  // 外网ip
		goVersion           string  // go版本
		sdkVersion          string  // sdk版本
	}
	slog struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

// GinClientConfig 框架实例配置
type GinClientConfig struct {
	IpService *goip.Client // IP服务
	CurrentIp string       // 当前IP
}

// NewGinClient 创建框架实例化
func NewGinClient(ctx context.Context, config *GinClientConfig) (*GinClient, error) {

	c := &GinClient{}

	if config.CurrentIp != "" && config.CurrentIp != "0.0.0.0" {
		c.config.systemOutsideIp = config.CurrentIp
	}
	c.config.systemOutsideIp = goip.IsIp(c.config.systemOutsideIp)
	if c.config.systemOutsideIp == "" {
		return nil, currentIpNoConfig
	}

	c.ipService = config.IpService

	// 配置信息
	c.setConfig(ctx)

	return c, nil
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (c *GinClient) jsonUnmarshal(data string) (result interface{}) {
	_ = gojson.Unmarshal([]byte(data), &result)
	return
}

// Middleware 中间件
func (c *GinClient) Middleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {

		// 开始时间
		startTime := gotime.Current().TimestampWithMillisecond()
		requestTime := gotime.Current().Time

		// 获取全部内容
		paramsBody := gorequest.NewParams()
		queryParams := ginCtx.Request.URL.Query() // 请求URL参数
		for key, values := range queryParams {
			for _, value := range values {
				paramsBody.Set(key, value)
			}
		}
		var dataMap map[string]interface{}
		rawData, _ := ginCtx.GetRawData() // 请求内容参数
		if gojson.IsValidJSON(string(rawData)) {
			dataMap = gojson.JsonDecodeNoError(string(rawData))
		} else {
			dataMap = gojson.ParseQueryString(string(rawData))
		}
		for key, value := range dataMap {
			paramsBody.Set(key, value)
		}

		// 重新赋值
		ginCtx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

		// 处理请求
		ginCtx.Next()

		//结束时间
		endTime := gotime.Current().TimestampWithMillisecond()

		go func() {

			clientIp := gorequest.ClientIp(ginCtx.Request)
			var info = goip.AnalyseResult{}

			if c.ipService != nil {
				info = c.ipService.Analyse(clientIp)
			}

			var traceId = gotrace_id.GetGinTraceId(ginCtx)

			// 记录
			c.recordJson(ginCtx, traceId, requestTime, paramsBody, startTime, endTime, info)

		}()
	}
}
