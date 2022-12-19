package nldyp

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestServeHttpPartnerData4GetOrderData 返回参数
type RequestServeHttpPartnerData4GetOrderData struct {
	Vendor  string `json:"vendor"`
	Ts      int    `json:"ts"`
	Sign    string `json:"sign"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status      int     `json:"status"`
		OrderNumber string  `json:"orderNumber"`
		Cpje        float64 `json:"cpje"`
		TicketList  []struct {
			TicketCode string `json:"ticketCode"`
			Url        string `json:"url"`
			YuanUrl    string `json:"yuan_url"`
		} `json:"ticketList"`
		Beizhu         string `json:"beizhu"`
		IsChangeSeat   string `json:"is_changeSeat"`
		ChangeSeatName string `json:"changeSeatName"`
	} `json:"data"`
}

// ServeHttpPartnerData4GetOrderData 回调
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=2e99efc0-6c77-457f-80d5-adaf19fdf313
func (c *Client) ServeHttpPartnerData4GetOrderData(ctx context.Context, gCtx *gin.Context) (validateJson RequestServeHttpPartnerData4GetOrderData, err error) {
	// 声明接收的变量
	err = gCtx.ShouldBind(&validateJson)
	return
}

// Success 数据正常
func (r *RequestServeHttpPartnerData4GetOrderData) Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

// Error 数据错误
func (r *RequestServeHttpPartnerData4GetOrderData) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
