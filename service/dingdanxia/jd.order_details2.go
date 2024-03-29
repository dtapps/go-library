package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

// JdOrderDetails2 【官方不维护】 京东联盟订单行查询
// https://www.dingdanxia.com/doc/180/94
func (c *Client) JdOrderDetails2(ctx context.Context, notMustParams ...gorequest.Params) (*JdJyOrderDetailsResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/jd/order_details2", params, http.MethodPost)
	if err != nil {
		return newJdJyOrderDetailsResult(JdJyOrderDetailsResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response JdJyOrderDetailsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newJdJyOrderDetailsResult(response, request.ResponseBody, request), err
}
