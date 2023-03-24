package goip

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"log"
	"net"
)

// GetInsideIp 内网ip
func GetInsideIp(ctx context.Context) string {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// Ips 获取全部网卡的全部IP
func Ips(ctx context.Context) (map[string]string, error) {

	ips := make(map[string]string)

	//返回 interface 结构体对象的列表，包含了全部网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	//遍历全部网卡
	for _, i := range interfaces {

		// Addrs() 方法返回一个网卡上全部的IP列表
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		//遍历一个网卡上全部的IP列表，组合为一个字符串，放入对应网卡名称的map中
		for _, v := range address {
			ips[i.Name] += v.String() + " "
		}
	}
	return ips, nil
}

// GetOutsideIp 外网ip
func GetOutsideIp(ctx context.Context) string {

	// 返回结果
	type respGetOutsideIp struct {
		Data struct {
			Ip string `json:"ip,omitempty"`
		} `json:"data"`
	}

	// 请求
	getHttp := gorequest.NewHttp()
	getHttp.SetUri("https://api.dtapp.net/ip")
	response, err := getHttp.Get(ctx)
	if err != nil {
		log.Printf("[GetOutsideIp]getHttp.Get：%s\n", err)
		return "0.0.0.0"
	}
	// 解析
	var responseJson respGetOutsideIp
	err = gojson.Unmarshal(response.ResponseBody, &responseJson)
	if err != nil {
		log.Printf("[GetOutsideIp]json.Unmarshal：%s\n", err)
		return "0.0.0.0"
	}
	if responseJson.Data.Ip == "" {
		responseJson.Data.Ip = "0.0.0.0"
	}
	return responseJson.Data.Ip
}

// GetMacAddr 获取Mac地址
func GetMacAddr(ctx context.Context) (arrays []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return arrays
	}
	for _, netInterface := range netInterfaces {
		addr := netInterface.HardwareAddr.String()
		if len(addr) == 0 {
			continue
		}

		arrays = append(arrays, addr)
	}
	return arrays
}
