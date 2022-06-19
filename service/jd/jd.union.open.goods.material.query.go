package jd

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type UnionOpenGoodsMaterialQueryResultResponse struct {
	JdUnionOpenGoodsMaterialQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_material_query_responce"`
}

type UnionOpenGoodsMaterialQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		BrandCode    string `json:"brandCode"`
		BrandName    string `json:"brandName"`
		CategoryInfo struct {
			Cid1     int    `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     int    `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     int    `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		Comments       int `json:"comments"`
		CommissionInfo struct {
			Commission          float64 `json:"commission"`
			CommissionShare     float64 `json:"commissionShare"`
			CouponCommission    float64 `json:"couponCommission"`
			PlusCommissionShare float64 `json:"plusCommissionShare"`
		} `json:"commissionInfo"`
		CouponInfo struct {
			CouponList []struct {
				BindType     int     `json:"bindType"`
				Discount     float64 `json:"discount"`
				GetEndTime   int64   `json:"getEndTime"`
				GetStartTime int64   `json:"getStartTime"`
				IsBest       int     `json:"isBest"`
				Link         string  `json:"link"`
				PlatformType int     `json:"platformType"`
				Quota        float64 `json:"quota"`
				UseEndTime   int64   `json:"useEndTime"`
				UseStartTime int64   `json:"useStartTime"`
			} `json:"couponList"`
		} `json:"couponInfo"`
		DeliveryType      int     `json:"deliveryType"`
		ForbidTypes       []int   `json:"forbidTypes"`
		GoodCommentsShare float64 `json:"goodCommentsShare"`
		ImageInfo         struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			WhiteImage string `json:"whiteImage,omitempty"`
		} `json:"imageInfo"`
		InOrderCount30Days    int    `json:"inOrderCount30Days"`
		InOrderCount30DaysSku int    `json:"inOrderCount30DaysSku"`
		IsHot                 int    `json:"isHot"`
		JxFlags               []int  `json:"jxFlags,omitempty"`
		MaterialUrl           string `json:"materialUrl"`
		Owner                 string `json:"owner"`
		PinGouInfo            struct {
			PingouEndTime   int64   `json:"pingouEndTime,omitempty"`
			PingouPrice     float64 `json:"pingouPrice,omitempty"`
			PingouStartTime int64   `json:"pingouStartTime,omitempty"`
			PingouTmCount   int     `json:"pingouTmCount,omitempty"`
			PingouUrl       string  `json:"pingouUrl,omitempty"`
		} `json:"pinGouInfo"`
		PriceInfo struct {
			LowestCouponPrice float64 `json:"lowestCouponPrice"`
			LowestPrice       float64 `json:"lowestPrice"`
			LowestPriceType   int     `json:"lowestPriceType"`
			Price             float64 `json:"price"`
		} `json:"priceInfo"`
		PromotionInfo struct {
			ClickURL string `json:"clickURL"`
		} `json:"promotionInfo"`
		ResourceInfo struct {
			EliteId   int    `json:"eliteId"`
			EliteName string `json:"eliteName"`
		} `json:"resourceInfo"`
		ShopInfo struct {
			ShopId                        int     `json:"shopId"`
			ShopLabel                     string  `json:"shopLabel"`
			ShopLevel                     float64 `json:"shopLevel"`
			ShopName                      string  `json:"shopName"`
			AfsFactorScoreRankGrade       string  `json:"afsFactorScoreRankGrade,omitempty"`
			AfterServiceScore             string  `json:"afterServiceScore,omitempty"`
			CommentFactorScoreRankGrade   string  `json:"commentFactorScoreRankGrade,omitempty"`
			LogisticsFactorScoreRankGrade string  `json:"logisticsFactorScoreRankGrade,omitempty"`
			LogisticsLvyueScore           string  `json:"logisticsLvyueScore,omitempty"`
			ScoreRankRate                 string  `json:"scoreRankRate,omitempty"`
			UserEvaluateScore             string  `json:"userEvaluateScore,omitempty"`
		} `json:"shopInfo"`
		SkuId     int64  `json:"skuId"`
		SkuName   string `json:"skuName"`
		Spuid     int64  `json:"spuid"`
		VideoInfo struct {
		} `json:"videoInfo"`
	} `json:"data"`
	Message    string `json:"message"`
	RequestId  string `json:"requestId"`
	TotalCount int    `json:"totalCount"`
}

type UnionOpenGoodsMaterialQueryResult struct {
	Responce UnionOpenGoodsMaterialQueryResultResponse // 结果
	Result   UnionOpenGoodsMaterialQueryQueryResult    // 结果
	Body     []byte                                    // 内容
	Http     gorequest.Response                        // 请求
	Err      error                                     // 错误
}

func newUnionOpenGoodsMaterialQueryResult(responce UnionOpenGoodsMaterialQueryResultResponse, result UnionOpenGoodsMaterialQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenGoodsMaterialQueryResult {
	return &UnionOpenGoodsMaterialQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenGoodsMaterialQuery 猜你喜欢商品推荐
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.material.query
func (c *Client) UnionOpenGoodsMaterialQuery(notMustParams ...Params) *UnionOpenGoodsMaterialQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.goods.material.query", notMustParams...)
	// 请求
	request, err := c.request(params)
	// 定义
	var responce UnionOpenGoodsMaterialQueryResultResponse
	var result UnionOpenGoodsMaterialQueryQueryResult
	err = json.Unmarshal(request.ResponseBody, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenGoodsMaterialQueryResponce.QueryResult), &result)
	return newUnionOpenGoodsMaterialQueryResult(responce, result, request.ResponseBody, request, err)
}
