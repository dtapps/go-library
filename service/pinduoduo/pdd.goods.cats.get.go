package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newGoodsCatsGetResult(result GoodsCatsGetResponse, body []byte, http gorequest.Response) *GoodsCatsGetResult {
	return &GoodsCatsGetResult{Result: result, Body: body, Http: http}
}

// GoodsCatsGet 商品标准类目接口
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.cats.get
func (c *Client) GoodsCatsGet(ctx context.Context, parentOptId int, notMustParams ...gorequest.Params) (*GoodsCatsGetResult, error) {
	// 参数
	params := NewParamsWithType("pdd.goods.cats.get", notMustParams...)
	params.Set("parent_opt_id", parentOptId)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newGoodsCatsGetResult(GoodsCatsGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GoodsCatsGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGoodsCatsGetResult(response, request.ResponseBody, request), err
}
