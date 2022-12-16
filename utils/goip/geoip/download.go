package geoip

import (
	"io/ioutil"
	"log"
	"net/http"
)

func OnlineDownload(downloadUrl string, downloadName string) {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = ioutil.WriteFile("./"+downloadName, body, 0644)
	if err != nil {
		panic(err)
	}
	log.Printf("已下载最新 ip2region.xdb 数据库 %s ", "./"+downloadName)
}
