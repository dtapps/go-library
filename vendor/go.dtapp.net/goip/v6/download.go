package v6

import (
	"github.com/saracen/go7z"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getOnline() (data []byte, err error) {
	resp, err := http.Get("https://ip.zxinc.org/ip.7z")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	file7z, err := ioutil.TempFile("", "*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file7z.Name())

	if err := ioutil.WriteFile(file7z.Name(), body, 0644); err == nil {
		return Un7z(file7z.Name())
	}
	return
}

func Un7z(filePath string) (data []byte, err error) {
	sz, err := go7z.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	defer sz.Close()

	fileNoNeed, err := ioutil.TempFile("", "*")
	if err != nil {
		return nil, err
	}
	fileNeed, err := ioutil.TempFile("", "*")
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	for {
		hdr, err := sz.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}

		if hdr.Name == "ipv6wry.db" {
			if _, err := io.Copy(fileNeed, sz); err != nil {
				log.Fatalln("ZX ipv6数据库解压出错：", err.Error())
			}
		} else {
			if _, err := io.Copy(fileNoNeed, sz); err != nil {
				log.Fatalln("ZX ipv6数据库解压出错：", err.Error())
			}
		}
	}
	err = fileNoNeed.Close()
	if err != nil {
		return nil, err
	}
	defer os.Remove(fileNoNeed.Name())
	defer os.Remove(fileNeed.Name())
	return ioutil.ReadFile(fileNeed.Name())
}
