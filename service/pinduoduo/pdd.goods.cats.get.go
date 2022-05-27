package pinduoduo

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsCatsGetResponse struct {
	GoodsCatsGetResponse struct {
		GoodsCatsList []struct {
			CatId       int    `json:"cat_id"`        // 商品类目ID
			CatName     string `json:"cat_name"`      // 商品类目名称
			Level       int    `json:"level"`         // 类目层级，1-一级类目，2-二级类目，3-三级类目，4-四级类目
			ParentCatID int    `json:"parent_cat_id"` // id所属父类目ID，其中，parent_id=0时为顶级节点
		} `json:"goods_cats_list"`
	} `json:"goods_cats_get_response"`
}

type GoodsCatsGetResult struct {
	Result GoodsCatsGetResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
	Err    error                // 错误
}

func NewGoodsCatsGetResult(result GoodsCatsGetResponse, body []byte, http gorequest.Response, err error) *GoodsCatsGetResult {
	return &GoodsCatsGetResult{Result: result, Body: body, Http: http, Err: err}
}

// GoodsCatsGet 商品标准类目接口
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.cats.get
func (app *App) GoodsCatsGet(parentOptId int) *GoodsCatsGetResult {
	// 参数
	param := NewParams()
	param.Set("parent_cat_id", parentOptId)
	params := NewParamsWithType("pdd.goods.cats.get", param)
	// 请求
	request, err := app.request(params)
	// 定义
	var response GoodsCatsGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGoodsCatsGetResult(response, request.ResponseBody, request, err)
}
