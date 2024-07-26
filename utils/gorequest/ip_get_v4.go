package gorequest

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
)

func getIPV4_DtappNet(ctx context.Context) string {

	// 返回结果
	type respGetOutsideIp struct {
		Data struct {
			Ip string `json:"ip,omitempty"`
		} `json:"data"`
	}

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://api.dtapp.net/ip")
	getHttp.SetUserAgent(GetRandomUserAgentSystem())
	response, err := getHttp.Get(ctx)
	if err != nil {
		return "0.0.0.0"
	}
	// 解析
	var responseJson respGetOutsideIp
	err = json.Unmarshal(response.ResponseBody, &responseJson)
	if err != nil {
		return "0.0.0.0"
	}
	if responseJson.Data.Ip == "" {
		responseJson.Data.Ip = "0.0.0.0"
	}
	return responseJson.Data.Ip
}

func getIPV4_MyipIpipNet(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://myip.ipip.net/")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	// 定义IPv4地址的正则表达式模式
	ipPattern := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

	// 使用正则表达式查找IPv4地址
	ipv4 := ipPattern.FindString(string(response.ResponseBody))

	return strings.Replace(strings.Replace(ipv4, " ", "", -1), "\n", "", -1)
}

func getIPV4_DdnsOrayCom(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://ddns.oray.com/checkip")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	// 定义IPv4地址的正则表达式模式
	ipPattern := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	// 使用正则表达式查找IPv4地址
	ipv4 := ipPattern.FindString(string(response.ResponseBody))

	return strings.Replace(strings.Replace(ipv4, " ", "", -1), "\n", "", -1)
}

func getIPV4_Ip3322Net(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://ip.3322.net/")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	return strings.Replace(strings.Replace(string(response.ResponseBody), " ", "", -1), "\n", "", -1)
}

func getIPV4_4IpwCn(ctx context.Context) string {

	// 请求
	getHttp := NewHttp()
	getHttp.SetUri("https://4.ipw.cn/")
	response, err := getHttp.Get(ctx)
	if err != nil {
		return ""
	}

	return strings.Replace(strings.Replace(string(response.ResponseBody), " ", "", -1), "\n", "", -1)
}
