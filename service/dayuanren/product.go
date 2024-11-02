package dayuanren

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type ProductResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   []struct {
		Id       int64  `json:"id"`   // 分类ID
		Cate     string `json:"cate"` // 分类名称
		Sort     int64  `json:"sort"` // 排序
		Type     int64  `json:"type"` // 产品类型ID
		Products []struct {
			Id           int64   `json:"id"`   // 产品ID,下单报文中用此参数
			Name         string  `json:"name"` // 产品名称
			Yname        string  `json:"yname,omitempty"`
			Desc         string  `json:"desc"`     // 产品说明
			ApiOpen      int64   `json:"api_open"` // 自动充值
			Isp          string  `json:"isp"`      // 运营商集合（话费、流量有效），1移动,2电信,3联通,4虚拟
			YsTag        string  `json:"ys_tag"`   // 标签
			Price        string  `json:"price"`    // 价格，下单扣费金额
			ShowStyle    int64   `json:"show_style,omitempty"`
			CateId       int64   `json:"cate_id,omitempty"`
			DelayApi     string  `json:"delay_api,omitempty"`
			YPrice       float64 `json:"y_price"`   // 原价
			MaxPrice     string  `json:"max_price"` // 封顶价格
			Type         int64   `json:"type"`      // 产品类型ID
			AllowPro     string  `json:"allow_pro,omitempty"`
			AllowCity    string  `json:"allow_city,omitempty"`
			ForbidPro    string  `json:"forbid_pro,omitempty"`
			ForbidCity   string  `json:"forbid_city,omitempty"`
			JmapiId      int64   `json:"jmapi_id,omitempty"`
			JmapiParamId int64   `json:"jmapi_param_id,omitempty"`
			IsJiema      int64   `json:"is_jiema,omitempty"`
			CateName     string  `json:"cate_name"` // 产品分类名称
			TypeName     string  `json:"type_name"` // 产品类型名称
			TypecId      int64   `json:"typec_id,omitempty"`
		} `json:"products"` // 产品列表
	} `json:"data,omitempty"`
}

type ProductResult struct {
	Result ProductResponse    // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newProductResult(result ProductResponse, body []byte, http gorequest.Response) *ProductResult {
	return &ProductResult{Result: result, Body: body, Http: http}
}

// Product 获取产品
// type = 产品类型ID
// cate_id = 分类ID
// https://www.showdoc.com.cn/dyr/9227005691961526
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097253
func (c *Client) Product(ctx context.Context, notMustParams ...gorequest.Params) (*ProductResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID()) // 商户ID

	// 请求
	var response ProductResponse
	request, err := c.request(ctx, "index/product", params, &response)
	return newProductResult(response, request.ResponseBody, request), err
}
