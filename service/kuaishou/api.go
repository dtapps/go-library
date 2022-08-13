package kuaishou

import (
	"context"
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

func newAnalysisResult(result AnalysisResponse, err error) *AnalysisResult {
	return &AnalysisResult{Result: result, Err: err}
}

// Analysis 快手解析
func (c *Client) Analysis(ctx context.Context, content string) *AnalysisResult {

	// 定义
	var response AnalysisResponse

	// 提取url
	var url string
	if strings.Contains(content, "kuaishou.com") {
		url = xurls.Relaxed.FindString(content)
	} else if strings.Contains(content, "gifshow.com") {
		url = xurls.Relaxed.FindString(content)
	} else {
		return newAnalysisResult(response, errors.New("url为空"))
	}

	// 获取重定向链接
	result := c.GetVideoLink(url)
	if result.Err != nil {
		return newAnalysisResult(response, result.Err)
	}

	// 获取重定向内容
	html, err := c.GetVideoHtml(result.Link, result.Cookies)
	if err != nil {
		return newAnalysisResult(response, result.Err)
	}

	// 判断
	imageLinks := c.ExtractImageLink(html)
	videoLink := c.ExtractVideoLink(html)

	// 0 是视频，1是图集
	if len(imageLinks) > 0 {
		response.ImageLinkList = imageLinks

	} else if len(videoLink) > 0 {
		response.VideoLink = videoLink
	}

	return newAnalysisResult(response, err)
}
