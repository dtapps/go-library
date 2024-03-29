package leshuazf

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataMccResponse struct {
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	ReqSerialNo string `json:"reqSerialNo"`
	Data        []struct {
		GrandpaCode string `json:"grandpaCode"`
		GrandpaText string `json:"grandpaText"`
	} `json:"data"`
}

type DataMccResult struct {
	Result DataMccResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newDataMccResult(result DataMccResponse, body []byte, http gorequest.Response) *DataMccResult {
	return &DataMccResult{Result: result, Body: body, Http: http}
}

// DataMcc 代理商通过MccCode来查商户类别明细
// https://www.yuque.com/leshuazf/doc/dbmxyi#jRTHN
func (c *Client) DataMcc(ctx context.Context, notMustParams ...gorequest.Params) (*DataMccResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/data/mcc", params, http.MethodPost)
	if err != nil {
		return newDataMccResult(DataMccResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DataMccResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataMccResult(response, request.ResponseBody, request), err
}
