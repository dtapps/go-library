package qqwry

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"
)

// 解密key
// https://zhangzifan.com/update-qqwry-dat.html
func getKey() (uint32, error) {
	resp, err := http.Get("https://update.cz88.net/ip/copywrite.rar")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return 0, err
	} else {
		// @see https://stackoverflow.com/questions/34078427/how-to-read-packed-binary-data-in-go
		return binary.LittleEndian.Uint32(body[5*4:]), nil
	}
}

func OnlineDownload() {
	resp, err := http.Get("https://update.cz88.net/ip/qqwry.rar")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	key, err := getKey()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 0x200; i++ {
		key = key * 0x805
		key++
		key = key & 0xff

		body[i] = byte(uint32(body[i]) ^ key)
	}

	reader, err := zlib.NewReader(bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	tmpData, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./qqwry.dat", tmpData, 0644)
	if err != nil {
		panic(err)
	}

	log.Printf("已下载最新 纯真 IPv4数据库 %s ", "./qqwry.dat")
}
