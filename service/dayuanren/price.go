package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type Price struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   struct {
		Name     string `json:"name"`      // 产品名称
		Desc     string `json:"desc"`      // 产品说明
		ApiOpen  string `json:"api_open"`  // 自动充值
		Isp      string `json:"isp"`       // 运营商集合（话费、流量有效），1移动,2电信,3联通,4虚拟
		YsTag    string `json:"ys_tag"`    // 标签
		Price    string `json:"price"`     // 价格，下单扣费金额
		YPrice   string `json:"y_price"`   // 原价
		MaxPrice string `json:"max_price"` // 封顶价格
		Type     string `json:"type"`      // 产品类型ID
		CateName string `json:"cate_name"` // 产品分类名称
		TypeName string `json:"type_name"` // 产品类型名称
	} `json:"data,omitempty"`
}

// Price 产品ID查询【新增】
// id = 产品ID
// https://www.showdoc.com.cn/dyr/9757701226597233
func (c *Client) Price(ctx context.Context, id int64, notMustParams ...*gorequest.Params) (response Price, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID()) // 商户ID
	params.Set("id", id)                // 产品ID

	// 请求
	err = c.request(ctx, "index/price", params, &response)
	return
}
