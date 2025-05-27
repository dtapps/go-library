package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type GoodsOptGetGoodsOptList struct {
	OptId       int64  `json:"opt_id"`        // 商品标签ID
	OptName     string `json:"opt_name"`      // 商品标签名
	Level       int64  `json:"level"`         // 层级，1-一级，2-二级，3-三级，4-四级
	ParentOptId int64  `json:"parent_opt_id"` // id所属父ID，其中，parent_id=0时为顶级节点
}

type GoodsOptGet struct {
	GoodsOptGetResponse struct {
		GoodsOptList []GoodsOptGetGoodsOptList `json:"goods_opt_list"` // opt列表
	} `json:"goods_opt_get_response"`
}

// GoodsOptGet 查询商品标签列表
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.opt.get
func (c *Client) GoodsOptGet(ctx context.Context, parentOptId int, notMustParams ...*gorequest.Params) (response GoodsOptGet, err error) {

	// 参数
	params := NewParamsWithType("pdd.goods.opt.get", notMustParams...)
	params.Set("parent_opt_id", parentOptId) // 值=0时为顶点opt_id,通过树顶级节点获取opt树

	// 请求
	err = c.request(ctx, params, &response)
	return
}
