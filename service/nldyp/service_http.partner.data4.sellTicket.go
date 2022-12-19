package nldyp

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestServeHttpPartnerData4SellTicket 返回参数
type RequestServeHttpPartnerData4SellTicket struct {
	OrderId     string `json:"orderId"`
	Account     string `json:"account"`
	CinemaId    string `json:"cinemaId"`
	RequestTime string `json:"requestTime"`
	Reason      string `json:"reason"`
	VerifyInfo  string `json:"verifyInfo"`
}

// ServeHttpPartnerData4SellTicket 退票通知
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=6f36e1f2-bdbc-476e-b9ca-a8806c0087ab
func (c *Client) ServeHttpPartnerData4SellTicket(ctx context.Context, gCtx *gin.Context) (validateJson RequestServeHttpPartnerData4SellTicket, err error) {
	// 声明接收的变量
	err = gCtx.ShouldBind(&validateJson)
	return
}

// Success 数据正常
func (r *RequestServeHttpPartnerData4SellTicket) Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

// Error 数据错误
func (r *RequestServeHttpPartnerData4SellTicket) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
