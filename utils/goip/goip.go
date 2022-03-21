package goip

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gohttp"
	"log"
	"net"
)

// GetInsideIp 内网ip
func GetInsideIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

var respGetOutsideIp struct {
	Ip string `json:"ip"`
}

// GetOutsideIp 外网ip
func GetOutsideIp() string {
	get, _ := gohttp.Get("https://api.dtapp.net/ip", map[string]interface{}{})
	_ = json.Unmarshal(get.Body, &respGetOutsideIp)
	if respGetOutsideIp.Ip != "" {
		return respGetOutsideIp.Ip
	}
	return "0.0.0.0"
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
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
