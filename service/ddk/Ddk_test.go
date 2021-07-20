package ddk

import (
	"fmt"
	"testing"
)

func init() {
	ClientId = "c0372aa7ffa149cbbce852e4d397a577"
	ClientSecret = "7d527f81d80bc41527dd8d680a462ff06fbfb14b"
}

func TestName(t *testing.T) {
	fmt.Println("Hello World")
	res, err := Execute("pdd.ddk.goods.recommend.get", Parameter{
		"limit":        10,
		"channel_type": 3,
		"offset":       0,
		"pid":          "1923953_141325051",
		"goods_sign_list": ParameterJsonEncode{
			"Y9v2lh2s6e1GWdnxwfbZF9sHlepFWs13_JmF4wnW72",
		},
	})

	if err != nil {
		fmt.Printf("错误：%#v\n", err)
	}
	fmt.Printf("结果：%#v\n", res)
}
