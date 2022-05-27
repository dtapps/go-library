package goip

import (
	"encoding/json"
	"go.dtapp.net/gorequest"
	"net"
)

// GetInsideIp 内网ip
func GetInsideIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// Ips 获取全部网卡的全部IP
func Ips() (map[string]string, error) {

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

var respGetOutsideIp struct {
	Data struct {
		Ip string `json:"ip"`
	} `json:"data"`
}

// GetOutsideIp 外网ip
func GetOutsideIp() (ip string) {
	ip = "0.0.0.0"
	get := gorequest.NewHttp()
	get.SetUri("https://api.dtapp.net/ip")
	response, err := get.Get()
	if err != nil {
		return
	}
	err = json.Unmarshal(response.ResponseBody, &respGetOutsideIp)
	if err != nil {
		return
	}
	if respGetOutsideIp.Data.Ip == "" {
		return
	}
	ip = respGetOutsideIp.Data.Ip
	return respGetOutsideIp.Data.Ip
}

// GetMacAddr 获取Mac地址
func GetMacAddr() (arrays []string) {
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
