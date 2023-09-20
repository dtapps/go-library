package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkDgMaterialOptionalResponse struct {
	TbkDgMaterialOptionalResponse struct {
		ResultList struct {
			MapData []struct {
				CategoryId             int    `json:"category_id"`
				CategoryName           string `json:"category_name"`
				CommissionRate         string `json:"commission_rate"`
				CommissionType         string `json:"commission_type"`
				CouponId               string `json:"coupon_id"`
				CouponInfo             string `json:"coupon_info"`
				CouponRemainCount      int    `json:"coupon_remain_count"`
				CouponCount            int    `json:"coupon__count"`
				CpaRewardType          string `json:"cpa_reward_type"`
				IncludeDxjh            string `json:"include_dxjh"`
				IncludeMkt             string `json:"include_mkt"`
				InfoDxjh               string `json:"info_dxjh"`
				ItemDescription        string `json:"item_description"`
				ItemId                 string `json:"item_id"`
				ItemUrl                string `json:"item_url"`
				LevelOneCategoryId     int64  `json:"level_one_category_id"`
				LevelOneCategoryName   string `json:"level_one_category_name"`
				Nick                   string `json:"nick"`
				NumIid                 int64  `json:"num_iid"`
				PictUrl                string `json:"pict_url"`
				Presale                int    `json:"presale"`
				PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
				PresaleEndTime         int64  `json:"presale_end_time"`
				PresaleStartTime       int64  `json:"presale_start_time"`
				PresaleTailEndTime     int64  `json:"presale_tail_end_time"`
				PresaleTailStartTime   int64  `json:"presale_tail_start_time"`
				Provcity               string `json:"provcity"`
				RealPostFe             int    `json:"real_post_fe"`
				ReservePrice           string `json:"reserve_price"`
				SellerId               int64  `json:"seller_id"`
				ShopDsr                int    `json:"shop_dsr"`
				ShopTitle              string `json:"shop_title"`
				ShortTitle             string `json:"short_title"`
				SmallImages            struct {
					String []string `json:"string"`
				} `json:"small_images"`
				SuperiorBrand  string `json:"superior_brand"`
				Title          string `json:"title"`
				TkTotalCommi   string `json:"tk_total_commi"`
				TkTotalSales   string `json:"tk_total_sales"`
				Url            string `json:"url"`
				UserType       int    `json:"user_type"`
				Volume         int64  `json:"volume"`
				WhiteImage     string `json:"white_image"`
				XId            string `json:"x_id"`
				ZkFinalPrice   string `json:"zk_final_price"`
				CouponShareUrl string `json:"coupon_share_url"`
				CouponAmount   string `json:"coupon_amount"`
			} `json:"map_data"`
		} `json:"result_list"`
		TotalResults  int64  `json:"total_results"`
		RequestId     string `json:"request_id"`
		PageResultKey string `json:"page_result_key,omitempty"`
	} `json:"tbk_dg_material_optional_response"`
}

type TbkDgMaterialOptionalResult struct {
	Result TbkDgMaterialOptionalResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func newTbkDgMaterialOptionalResult(result TbkDgMaterialOptionalResponse, body []byte, http gorequest.Response, err error) *TbkDgMaterialOptionalResult {
	return &TbkDgMaterialOptionalResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkDgMaterialOptional 淘宝客-推广者-物料搜索
// https://open.taobao.com/api.htm?docId=35896&docType=2&source=search
func (c *Client) TbkDgMaterialOptional(ctx context.Context, notMustParams ...*gorequest.Params) *TbkDgMaterialOptionalResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.dg.material.optional", notMustParams...)
	params.Set("adzone_id", c.GetAdzoneId())
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var response TbkDgMaterialOptionalResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkDgMaterialOptionalResult(response, request.ResponseBody, request, err)
}
