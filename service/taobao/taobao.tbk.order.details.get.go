package taobao

import "encoding/json"

type TbkOrderDetailsGetResponse struct {
	TbkOrderDetailsGetResponse struct {
		Data struct {
			HasNext       bool   `json:"has_next"`
			HasPre        bool   `json:"has_pre"`
			PageNo        int    `json:"page_no"`
			PageSize      int    `json:"page_size"`
			PositionIndex string `json:"position_index"`
			Results       struct {
				PublisherOrderDto []struct {
					AdzoneId                           int64  `json:"adzone_id"`
					AdzoneName                         string `json:"adzone_name"`
					AlimamaRate                        string `json:"alimama_rate"`
					AlimamaShareFee                    string `json:"alimama_share_fee"`
					AlipayTotalPrice                   string `json:"alipay_total_price"`
					ClickTime                          string `json:"click_time"`
					DepositPrice                       string `json:"deposit_price"`
					FlowSource                         string `json:"flow_source"`
					IncomeRate                         string `json:"income_rate"`
					IsLx                               string `json:"is_lx"`
					ItemCategoryName                   string `json:"item_category_name"`
					ItemImg                            string `json:"item_img"`
					ItemNum                            int    `json:"item_num"`
					ItemTitle                          string `json:"item_title"`
					MarketingType                      string `json:"marketing_type"`
					ModifiedTime                       string `json:"modified_time"`
					OrderType                          string `json:"order_type"`
					PayPrice                           string `json:"pay_price"`
					PubId                              int    `json:"pub_id"`
					PubShareFee                        string `json:"pub_share_fee"`
					PubSharePreFee                     string `json:"pub_share_pre_fee"`
					PubShareRate                       string `json:"pub_share_rate"`
					RefundTag                          int    `json:"refund_tag"`
					SellerNick                         string `json:"seller_nick"`
					SellerShopTitle                    string `json:"seller_shop_title"`
					SiteId                             int    `json:"site_id"`
					SiteName                           string `json:"site_name"`
					SubsidyFee                         string `json:"subsidy_fee"`
					SubsidyRate                        string `json:"subsidy_rate"`
					SubsidyType                        string `json:"subsidy_type"`
					TbDepositTime                      string `json:"tb_deposit_time"`
					TbPaidTime                         string `json:"tb_paid_time"`
					TerminalType                       string `json:"terminal_type"`
					TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
					TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
					TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
					TkCreateTime                       string `json:"tk_create_time"`
					TkDepositTime                      string `json:"tk_deposit_time"`
					TkEarningTime                      string `json:"tk_earning_time"`
					TkOrderRole                        int    `json:"tk_order_role"`
					TkPaidTime                         string `json:"tk_paid_time"`
					TkStatus                           int    `json:"tk_status"`
					TkTotalRate                        string `json:"tk_total_rate"`
					TotalCommissionFee                 string `json:"total_commission_fee"`
					TotalCommissionRate                string `json:"total_commission_rate"`
					TradeId                            string `json:"trade_id"`
					TradeParentId                      string `json:"trade_parent_id"`
				} `json:"publisher_order_dto"`
			} `json:"results"`
		} `json:"data"`
		RequestId string `json:"request_id"`
	} `json:"tbk_order_details_get_response"`
}

type TbkOrderDetailsGetResult struct {
	Result TbkOrderDetailsGetResponse // 结果
	Body   []byte                     // 内容
	Err    error                      // 错误
}

func NewTbkOrderDetailsGetResult(result TbkOrderDetailsGetResponse, body []byte, err error) *TbkOrderDetailsGetResult {
	return &TbkOrderDetailsGetResult{Result: result, Body: body, Err: err}
}

// TbkOrderDetailsGet 淘宝客-推广者-所有订单查询
// https://open.taobao.com/api.htm?docId=43328&docType=2&scopeId=16175
func (app *App) TbkOrderDetailsGet(notMustParams ...Params) *TbkOrderDetailsGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.order.details.get", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response TbkOrderDetailsGetResponse
	err = json.Unmarshal(body, &response)
	return NewTbkOrderDetailsGetResult(response, body, err)
}
