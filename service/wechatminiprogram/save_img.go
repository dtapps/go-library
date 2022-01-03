package wechatminiprogram

import (
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
	"os"
)

func (app *App) SaveImg(resp gohttp.Response, dir, saveName string) string {
	// 返回是二进制图片，或者json错误
	if resp.Header.Get("Content-Type") == "image/jpeg" || resp.Header.Get("Content-Type") == "image/png" {
		// 保存在output目录
		outputFileName := saveName

		if resp.Header.Get("Content-Type") == "image/jpeg" {
			outputFileName = outputFileName + ".jpg"
		} else {
			outputFileName = outputFileName + ".png"
		}
	here:
		f, err := os.OpenFile(dir+outputFileName, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			os.Mkdir(dir, 0666)
			goto here
		}
		f.Write(resp.Body)
		f.Close()
		return dir + outputFileName
	} else {
		return ""
	}
}
