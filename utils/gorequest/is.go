package gorequest

import (
	"errors"
	"net/http"
)

// IsWechatMiniProgramRequest 判断是否是微信小程序
func IsWechatMiniProgramRequest(r *http.Request, appid string) error {
	referer := r.Referer()
	userAgent := r.UserAgent()
	if referer == "" {
		return errors.New("网络请求没有达到要求")
	}
	return isWechatMiniProgramRequestReferer(userAgent, referer, appid)
}

// GinIsWechatMiniProgramRequest Gin框架 -> 判断是否是微信小程序
// https://gin-gonic.com/zh-cn/docs/
//func GinIsWechatMiniProgramRequest(ctx *gin.Context, appid string) error {
//	referer := ctx.Request.Referer()
//	userAgent := ctx.Request.UserAgent()
//	if referer == "" {
//		return errors.New("网络请求没有达到要求")
//	}
//	return isWechatMiniProgramRequestReferer(userAgent, referer, appid)
//}

// IrisIsWechatMiniProgramRequest Iris框架 -> 判断是否是微信小程序
// https://www.iris-go.com/docs/
//func IrisIsWechatMiniProgramRequest(ctx iris.Context, appid string) error {
//	referer := ctx.Request().Referer()
//	userAgent := ctx.Request().UserAgent()
//	if referer == "" {
//		return errors.New("网络请求没有达到要求")
//	}
//	return isWechatMiniProgramRequestReferer(userAgent, referer, appid)
//}

// EchoIsWechatMiniProgramRequest Echo框架 -> 判断是否是微信小程序
// https://echo.labstack.com/docs/
//func EchoIsWechatMiniProgramRequest(ctx echo.Context, appid string) error {
//	referer := ctx.Request().Referer()
//	userAgent := ctx.Request().UserAgent()
//	if referer == "" {
//		return errors.New("网络请求没有达到要求")
//	}
//	return isWechatMiniProgramRequestReferer(userAgent, referer, appid)
//}
