package kuaidi100

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PollQueryResponse struct {
	Message string `json:"message"` // 消息体，请忽略
	Nu      string `json:"nu"`      // 单号
	Ischeck string `json:"ischeck"` // 是否签收标记，0未签收，1已签收，请忽略，明细状态请参考state字段
	Com     string `json:"com"`     // 快递公司编码,一律用小写字母
	Status  string `json:"status"`  // 通讯状态，请忽略
	Data    []struct {
		Time       string `json:"time"`       // 时间，原始格式
		Context    string `json:"context"`    // 内容
		Ftime      string `json:"ftime"`      // 格式化后时间
		AreaCode   string `json:"areaCode"`   // 本数据元对应的行政区域的编码，实时查询接口中提交resultv2=1或者resultv2=4标记后才会出现
		AreaName   string `json:"areaName"`   // 本数据元对应的行政区域的名称，实时查询接口中提交resultv2=1或者resultv2=4标记后才会出现
		Status     string `json:"status"`     // 本数据元对应的物流状态名称或者高级状态名称，实时查询接口中提交resultv2=1或者resultv2=4标记后才会出现
		Location   string `json:"location"`   // 本数据元对应的快件当前地点，实时查询接口中提交resultv2=4标记后才会出现
		AreaCenter string `json:"areaCenter"` // 本数据元对应的行政区域经纬度，实时查询接口中提交resultv2=4标记后才会出现
		AreaPinYin string `json:"areaPinYin"` // 本数据元对应的行政区域拼音，实时查询接口中提交resultv2=4标记后才会出现
		StatusCode string `json:"statusCode"` // 本数据元对应的高级物流状态值，实时查询接口中提交resultv2=4标记后才会出现
	} `json:"data"` // 最新查询结果，数组，包含多项，全量，倒序（即时间最新的在最前），每项都是对象，对象包含字段请展开
	State     string `json:"state"`     // 快递单当前状态，默认为0在途，1揽收，2疑难，3签收，4退签，5派件，8清关，14拒签等10个基础物流状态，如需要返回高级物流状态，请参考 resultv2 传值
	Condition string `json:"condition"` // 快递单明细状态标记，暂未实现，请忽略
	RouteInfo struct {
		From struct {
			Number string `json:"number"`
			Name   string `json:"name"`
		} `json:"from"`
		Cur struct {
			Number string `json:"number"`
			Name   string `json:"name"`
		} `json:"cur"`
		To interface{} `json:"to"`
	} `json:"routeInfo"`
	IsLoop bool `json:"isLoop"`
}

type PollQueryResult struct {
	Result PollQueryResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newPollQueryResult(result PollQueryResponse, body []byte, http gorequest.Response) *PollQueryResult {
	return &PollQueryResult{Result: result, Body: body, Http: http}
}

// PollQuery 实时快递查询接口
// https://api.kuaidi100.com/document/5f0ffb5ebc8da837cbd8aefc
func (c *Client) PollQuery(ctx context.Context, notMustParams ...gorequest.Params) (*PollQueryResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/poll/query.do", params, http.MethodPost)
	if err != nil {
		return newPollQueryResult(PollQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PollQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPollQueryResult(response, request.ResponseBody, request), err
}
