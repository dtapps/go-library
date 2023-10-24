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

// NewGinClient 创建框架实例化
func NewGinClient(ctx context.Context, config *GinClientConfig) (*GinClient, error) {
	c := &GinClient{}
	c.ipService = config.IpService
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

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ginCtx.Writer}
		ginCtx.Writer = blw

		// 处理请求
		ginCtx.Next()

		// 响应
		responseCode := ginCtx.Writer.Status()
		responseBody := blw.body.String()

		// 结束时间
		endTime := gotime.Current().TimestampWithMillisecond()

		go func() {

			clientIp := gorequest.ClientIp(ginCtx.Request)
			var info = goip.AnalyseResult{}

			if c.ipService != nil {
				info = c.ipService.Analyse(clientIp)
			}

			var traceId = gotrace_id.GetGinTraceId(ginCtx)

			// 记录
			c.recordJson(ginCtx, traceId, requestTime, paramsBody, responseCode, responseBody, startTime, endTime, info)

		}()
	}
}
