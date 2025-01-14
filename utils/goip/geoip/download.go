package geoip

import (
	"io"
	"log"
	"net/http"
	"os"
)

func OnlineDownload(downloadUrl string, downloadName string) {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	err = os.WriteFile("./"+downloadName, body, 0644)
	if err != nil {
		panic(err)
	}
	log.Printf("已下载最新 geoip 数据库 %s ", "./"+downloadName)
}
