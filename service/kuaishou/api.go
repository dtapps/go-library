package kuaishou

import (
	"errors"
	"github.com/mvdan/xurls"
	"strings"
)

type AnalysisResponse struct {
	VideoLink     string   `json:"video_link"`
	ImageLinkList []string `json:"image_link_list"`
}

type AnalysisResult struct {
	Result AnalysisResponse // 结果
	Err    error            // 错误
}

func NewAnalysisResult(result AnalysisResponse, err error) *AnalysisResult {
	return &AnalysisResult{Result: result, Err: err}
}

// Analysis 快手解析
func (app *App) Analysis(content string) *AnalysisResult {

	// 定义
	var response AnalysisResponse

	// 提取url
	var url string
	if strings.Contains(content, "kuaishou.com") {
		url = xurls.Relaxed.FindString(content)
	} else if strings.Contains(content, "gifshow.com") {
		url = xurls.Relaxed.FindString(content)
	} else {
		return NewAnalysisResult(response, errors.New("url为空"))
	}

	// 获取重定向链接
	result := app.GetVideoLink(url)
	if result.Err != nil {
		return NewAnalysisResult(response, result.Err)
	}

	// 获取重定向内容
	html, err := app.GetVideoHtml(result.Link, result.Cookies)
	if err != nil {
		return NewAnalysisResult(response, result.Err)
	}

	// 判断
	imageLinks := app.ExtractImageLink(html)
	videoLink := app.ExtractVideoLink(html)

	// 0 是视频，1是图集
	if len(imageLinks) > 0 {
		response.ImageLinkList = imageLinks

	} else if len(videoLink) > 0 {
		response.VideoLink = videoLink
	}

	return NewAnalysisResult(response, err)
}
