package wechatunion

import (
	"fmt"
	"net/http"
)

type PromoterProductCategoryResult struct {
	Errcode     int    `json:"errcode"` // 错误码
	Errmsg      string `json:"errmsg"`  // 错误信息
	ProductCats []struct {
		CatId string `json:"catId"` // 类目ID
		Name  string `json:"name"`  // 类目名称
	} `json:"productCats"` // 类目数据
}

// PromoterProductCategory
// 获取联盟商品类目列表及类目ID
// 通过该接口获取联盟商品的一级类目列表以及类目ID，可用于筛选联盟商品
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/product/category.html
func (app *App) PromoterProductCategory() (body []byte, err error) {
	// 请求
	body, err = app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/product/category?access_token=%s", app.AccessToken), map[string]interface{}{}, http.MethodGet)
	return body, err
}
