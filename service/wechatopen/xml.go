package wechatopen

import (
	"encoding/xml"
	"io"
	"strings"
)

func XmlDecode(data string) map[string]string {
	decoder := xml.NewDecoder(strings.NewReader(data))
	result := make(map[string]string)
	key := ""
	for {
		token, err := decoder.Token() //读取一个标签或者文本内容
		if err == io.EOF {
			return result
		}
		if err != nil {
			return result
		}
		switch tp := token.(type) { //读取的TOKEN可以是以下三种类型：StartElement起始标签，EndElement结束标签，CharData文本内容
		case xml.StartElement:
			se := xml.StartElement(tp) //强制类型转换
			if se.Name.Local != "xml" {
				key = se.Name.Local
			}
			if len(se.Attr) != 0 {
				//读取标签属性
			}
		case xml.EndElement:
			ee := xml.EndElement(tp)
			if ee.Name.Local == "xml" {
				return result
			}
		case xml.CharData: //文本数据，注意一个结束标签和另一个起始标签之间可能有空格
			cd := xml.CharData(tp)
			data := strings.TrimSpace(string(cd))
			if len(data) != 0 {
				result[key] = data
			}
		}
	}
}
