package ip2region

import (
	"io/ioutil"
	"log"
	"net/http"
)

func OnlineDownload() {
	resp, err := http.Get("https://ghproxy.com/?q=https://github.com/lionsoul2014/ip2region/blob/master/v1.0/data/ip2region.db?raw=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile("./ip2region.db", body, 0644)
	if err != nil {
		panic(err)
	}
	log.Printf("已下载最新 ip2region 数据库 %s ", "./ip2region.db")
}
