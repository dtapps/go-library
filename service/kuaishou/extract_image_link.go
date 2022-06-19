package kuaishou

import "regexp"

func (c *Client) ExtractImageLink(content string) []string {

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`\{"path":"(.*?)","width":\d+,"height":\d*}`)

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(content, -1)
	var imageLinks []string

	if len(result) > 0 {
		for _, link := range result {
			imageLinks = append(imageLinks, "https://tx2.a.yximgs.com"+link[1])
		}

	}

	return imageLinks
}
