package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PromoterProductCategoryResponse struct {
	Errcode     int    `json:"errcode"` // 错误码
	Errmsg      string `json:"errmsg"`  // 错误信息
	ProductCats []struct {
		CatId string `json:"catId"` // 类目ID
		Name  string `json:"name"`  // 类目名称
	} `json:"productCats"` // 类目数据
}

type PromoterProductCategoryResult struct {
	Result PromoterProductCategoryResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
	Err    error                           // 错误
}

func NewPromoterProductCategoryResult(result PromoterProductCategoryResponse, body []byte, http gorequest.Response, err error) *PromoterProductCategoryResult {
	return &PromoterProductCategoryResult{Result: result, Body: body, Http: http, Err: err}
}

// PromoterProductCategory 获取联盟商品类目列表及类目ID
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/product/category.html#_1-%E8%8E%B7%E5%8F%96%E8%81%94%E7%9B%9F%E5%95%86%E5%93%81%E7%B1%BB%E7%9B%AE%E5%88%97%E8%A1%A8%E5%8F%8A%E7%B1%BB%E7%9B%AEID
func (c *Client) PromoterProductCategory() *PromoterProductCategoryResult {
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/promoter/product/category?access_token=%s", c.getAccessToken()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response PromoterProductCategoryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPromoterProductCategoryResult(response, request.ResponseBody, request, err)
}
