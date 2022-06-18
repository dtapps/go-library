package pinduoduo

import (
	"encoding/json"
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
	Err    error               // 错误
}

func NewGoodsOptGetResult(result GoodsOptGetResponse, body []byte, http gorequest.Response, err error) *GoodsOptGetResult {
	return &GoodsOptGetResult{Result: result, Body: body, Http: http, Err: err}
}

// GoodsOptGet 查询商品标签列表
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.opt.get
func (app *App) GoodsOptGet(parentOptId int) *GoodsOptGetResult {
	// 参数
	param := NewParams()
	param.Set("parent_opt_id", parentOptId)
	params := NewParamsWithType("pdd.goods.opt.get", param)
	// 请求
	request, err := app.request(params)
	// 定义
	var response GoodsOptGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGoodsOptGetResult(response, request.ResponseBody, request, err)
}
