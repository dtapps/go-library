package kuaishou

import "regexp"

func (c *Client) ExtractVideoLink(content string) string {

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`"srcNoMark":"(.*?)"`)

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(content, -1)
	var videoLink string

	if len(result) > 0 {
		videoLink = result[0][1]
		return videoLink
	}

	return ""
}
