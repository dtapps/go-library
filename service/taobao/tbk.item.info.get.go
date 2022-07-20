package taobao

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkItemInfoGetResponse struct {
	TbkItemInfoGetResponse struct {
		Results struct {
			NTbkItem []struct {
				CatLeafName          string `json:"cat_leaf_name"`
				CatName              string `json:"cat_name"`
				FreeShipment         bool   `json:"free_shipment"`
				HotFlag              string `json:"hot_flag"`
				ItemUrl              string `json:"item_url"`
				JuOnlineEnd          string `json:"ju_online_end"`
				JuOnlineStartTime    string `json:"ju_online_start_time"`
				JuPreShowEndTime     string `json:"ju_pre_show_end_time"`
				JuPreShowStartTime   string `json:"ju_pre_show_start_time"`
				MaterialLibType      string `json:"material_lib_type"`
				Nick                 string `json:"nick"`
				NumIid               int64  `json:"num_iid"`
				PictUrl              string `json:"pict_url"`
				PresaleDeposit       string `json:"presale_deposit"`
				PresaleEndTime       int    `json:"presale_end_time"`
				PresaleStartTime     int    `json:"presale_start_time"`
				PresaleTailEndTime   int    `json:"presale_tail_end_time"`
				PresaleTailStartTime int    `json:"presale_tail_start_time"`
				Provcity             string `json:"provcity"`
				ReservePrice         string `json:"reserve_price"`
				SellerId             int64  `json:"seller_id"`
				SmallImages          struct {
					String []string `json:"string"`
				} `json:"small_images"`
				SuperiorBrand                 string `json:"superior_brand"`
				Title                         string `json:"title"`
				TmallPllPlayActivityStartTime int    `json:"tmall_pll_play_activity_start_time"`
				UserType                      int    `json:"user_type"`
				Volume                        int64  `json:"volume"`
				ZkFinalPrice                  string `json:"zk_final_price"`
			} `json:"n_tbk_item"`
		} `json:"results"`
		RequestId string `json:"request_id"`
	} `json:"tbk_item_info_get_response"`
}

type TbkItemInfoGetResult struct {
	Result TbkItemInfoGetResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
	Err    error                  // 错误
}

func newTbkItemInfoGetResult(result TbkItemInfoGetResponse, body []byte, http gorequest.Response, err error) *TbkItemInfoGetResult {
	return &TbkItemInfoGetResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkItemInfoGet 淘宝客-公用-淘宝客商品详情查询(简版)
// https://open.taobao.com/api.htm?docId=24518&docType=2&source=search
func (c *Client) TbkItemInfoGet(notMustParams ...Params) *TbkItemInfoGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.item.info.get", notMustParams...)
	// 请求
	request, err := c.request(params)
	// 定义
	var response TbkItemInfoGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTbkItemInfoGetResult(response, request.ResponseBody, request, err)
}
