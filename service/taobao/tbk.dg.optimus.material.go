package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkDgOptimusMaterialResponse struct {
	TbkDgOptimusMaterialResponse struct {
		IsDefault  string `json:"is_default"`
		ResultList struct {
			MapData []struct {
				CategoryId           int    `json:"category_id"`
				ClickUrl             string `json:"click_url"`
				CommissionRate       string `json:"commission_rate"`
				CouponAmount         int64  `json:"coupon_amount"`
				CouponClickUrl       string `json:"coupon_click_url"`
				CouponEndTime        string `json:"coupon_end_time"`
				CouponRemainCount    int    `json:"coupon_remain_count"`
				CouponShareUrl       string `json:"coupon_share_url"`
				CouponStartFee       string `json:"coupon_start_fee"`
				CouponStartTime      string `json:"coupon_start_time"`
				CouponTotalCount     int    `json:"coupon_total_count"`
				CpaRewardType        string `json:"cpa_reward_type"`
				ItemDescription      string `json:"item_description"`
				ItemId               string `json:"item_id"`
				JhsPriceUspList      string `json:"jhs_price_usp_list"`
				LevelOneCategoryId   int64  `json:"level_one_category_id"`
				LevelOneCategoryName string `json:"level_one_category_name"`
				Nick                 string `json:"nick"`
				PictUrl              string `json:"pict_url"`
				ReservePrice         string `json:"reserve_price"`
				SellerId             int64  `json:"seller_id"`
				ShopTitle            string `json:"shop_title"`
				ShortTitle           string `json:"short_title"`
				SmallImages          struct {
					String []string `json:"string"`
				} `json:"small_images"`
				SubTitle     string `json:"sub_title"`
				Title        string `json:"title"`
				UserType     int    `json:"user_type"`
				Volume       int64  `json:"volume"`
				WhiteImage   string `json:"white_image"`
				ZkFinalPrice string `json:"zk_final_price"`
			} `json:"map_data"`
		} `json:"result_list"`
		RequestId  string `json:"request_id"`
		TotalCount int64  `json:"total_count,omitempty"`
	} `json:"tbk_dg_optimus_material_response"`
}

type TbkDgOptimusMaterialResult struct {
	Result TbkDgOptimusMaterialResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newTbkDgOptimusMaterialResult(result TbkDgOptimusMaterialResponse, body []byte, http gorequest.Response) *TbkDgOptimusMaterialResult {
	return &TbkDgOptimusMaterialResult{Result: result, Body: body, Http: http}
}

// TbkDgOptimusMaterial 淘宝客-推广者-物料精选
// https://open.taobao.com/api.htm?docId=33947&docType=2&source=search
func (c *Client) TbkDgOptimusMaterial(ctx context.Context, notMustParams ...gorequest.Params) (*TbkDgOptimusMaterialResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.dg.optimus.material", notMustParams...)
	params.Set("adzone_id", c.GetAdzoneId())
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkDgOptimusMaterialResult(TbkDgOptimusMaterialResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkDgOptimusMaterialResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkDgOptimusMaterialResult(response, request.ResponseBody, request), err
}
