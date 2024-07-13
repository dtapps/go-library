package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsOptGetResponse struct {
	GoodsOptGetResponse struct {
		GoodsOptList []struct {
			Level       int    `json:"level"`         // 层级，1-一级，2-二级，3-三级，4-四级
			OptId       int    `json:"opt_id"`        // 商品标签ID
			OptName     string `json:"opt_name"`      // 商品标签名
			ParentOptId int    `json:"parent_opt_id"` // id所属父ID，其中，parent_id=0时为顶级节点
		} `json:"goods_opt_list"`
	} `json:"goods_opt_get_response"`
}

type GoodsOptGetResult struct {
	Result GoodsOptGetResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGoodsOptGetResult(result GoodsOptGetResponse, body []byte, http gorequest.Response) *GoodsOptGetResult {
	return &GoodsOptGetResult{Result: result, Body: body, Http: http}
}

// GoodsOptGet 查询商品标签列表
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.opt.get
func (c *Client) GoodsOptGet(ctx context.Context, parentOptId int, notMustParams ...gorequest.Params) (*GoodsOptGetResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "pdd.goods.opt.get")
	defer c.TraceEndSpan()

	// 参数
	params := NewParamsWithType("pdd.goods.opt.get", notMustParams...)
	params.Set("parent_opt_id", parentOptId)

	// 请求
	var response GoodsOptGetResponse
	request, err := c.request(ctx, params, &response)
	return newGoodsOptGetResult(response, request.ResponseBody, request), err
}
