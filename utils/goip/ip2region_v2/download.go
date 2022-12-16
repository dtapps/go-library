package ip2region_v2

import (
	"io/ioutil"
	"log"
	"net/http"
)

func OnlineDownload() {
	resp, err := http.Get("https://ghproxy.com/?q=https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb?raw=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile("./ip2region.xdb", body, 0644)
	if err != nil {
		panic(err)
	}
	log.Printf("已下载最新 ip2region.xdb 数据库 %s ", "./ip2region.xdb")
}
