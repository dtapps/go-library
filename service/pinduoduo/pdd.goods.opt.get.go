package pinduoduo

// GoodsOptGetResult 返回参数
type GoodsOptGetResult struct {
	GoodsOptGetResponse struct {
		GoodsOptList []struct {
			Level       int    `json:"level"`         // 层级，1-一级，2-二级，3-三级，4-四级
			OptId       int    `json:"opt_id"`        // 商品标签ID
			OptName     string `json:"opt_name"`      // 商品标签名
			ParentOptId int    `json:"parent_opt_id"` // id所属父ID，其中，parent_id=0时为顶级节点
		} `json:"goods_opt_list"`
	} `json:"goods_opt_get_response"`
}

// GoodsOptGet 查询商品标签列表 https://open.pinduoduo.com/application/document/api?id=pdd.goods.opt.get
func (app *App) GoodsOptGet(parentOptId int) (body []byte, err error) {
	// 参数
	param := NewParams()
	param.Set("parent_opt_id", parentOptId)
	params := NewParamsWithType("pdd.goods.opt.get", param)
	// 请求
	body, err = app.request(params)
	return
}
