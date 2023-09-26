package leshuazf

import (
	"context"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 环境
	if c.GetEnvironment() == "test" {
		url = apiTestUrl + url
	} else {
		url = apiUrl + url
	}

	// 参数
	param.Set("agentId", c.GetAgentId())                                                       // 服务商编号，由乐刷分配的接入方唯一标识，明文传输。
	param.Set("version", "2.0")                                                                // 目前固定值2.0
	param.Set("reqSerialNo", gotime.Current().SetFormat("20060102150405")+gorandom.Numeric(5)) // 请求流水号(yyyyMMddHHmmssSSSXXXXX，其中 XXXXX为5位顺序号,禁止使用UUID等无意义数据)
	param.Set("sign", c.getSign(param))

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
