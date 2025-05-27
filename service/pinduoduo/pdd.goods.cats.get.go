package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type GoodsCatsGetGoodsCatsList struct {
	CatId       int64  `json:"cat_id"`        // 商品类目ID
	CatName     string `json:"cat_name"`      // 商品类目名称
	Level       int64  `json:"level"`         // 类目层级，1-一级类目，2-二级类目，3-三级类目，4-四级类目
	ParentCatID int64  `json:"parent_cat_id"` // id所属父类目ID，其中，parent_id=0时为顶级节点
}

type GoodsCatsGet struct {
	GoodsCatsGetResponse struct {
		GoodsCatsList []GoodsCatsGetGoodsCatsList `json:"goods_cats_list"` // 类目树对象
	} `json:"goods_cats_get_response"`
}

// GoodsCatsGet 商品标准类目接口
// https://open.pinduoduo.com/application/document/api?id=pdd.goods.cats.get
func (c *Client) GoodsCatsGet(ctx context.Context, parentCatId int64, notMustParams ...*gorequest.Params) (response GoodsCatsGet, err error) {

	// 参数
	params := NewParamsWithType("pdd.goods.cats.get", notMustParams...)
	params.Set("parent_cat_id", parentCatId) // 值=0时为顶点cat_id,通过树顶级节点获取cat树

	// 请求
	err = c.request(ctx, params, &response)
	return
}
