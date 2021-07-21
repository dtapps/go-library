package kashangwl

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	md52 "gopkg.in/dtapps/go-library.v2/md5"
	params2 "gopkg.in/dtapps/go-library.v2/params"
	string2 "gopkg.in/dtapps/go-library.v2/string"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const api = "http://www.kashangwl.com/api/"

// KaShangWl 每次请求需传入以下参数
type KaShangWl struct {
	CustomerId  int    // 商家编号
	CustomerKey string // 商家密钥
}

// Send 发送 http://cha.kashangwl.com/api-doc/
func (w *KaShangWl) Send(msg interface{}, url string) (*simplejson.Json, error) {
	// 	当前时间戳（单位：秒）
	timestamp := time.Now().UnixNano() / 1e6
	// 处理数据
	marshal, _ := json.Marshal(msg)
	newJson, _ := simplejson.NewJson(marshal)
	newJson.Set("customer_id", w.CustomerId)
	newJson.Set("timestamp", timestamp)
	signStr := sign(newJson, w.CustomerKey)
	newJson.Set("sign", signStr)
	j, e := newJson.MarshalJSON()
	if e != nil {
		return nil, e
	}
	resp, e := http.Post(api+url, "application/json", strings.NewReader(string(j)))
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	respJson, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// md5(key + 参数1名称 + 参数1值 + 参数2名称 + 参数2值...) 加密源串应为{key}customer_id1192442order_id827669582783timestamp1626845767
func sign(params *simplejson.Json, key string) string {
	var dataParams string
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params.MustMap() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	dataParams = fmt.Sprintf("%s%s", dataParams, key)
	for _, k := range keys {
		dataParams = fmt.Sprintf("%s%s%s", dataParams, k, params2.GetParamsString(params.Get(k)))
	}
	// MD5加密
	md5Str := md52.GetMD5Encode(dataParams)
	str := string2.ToLower(md5Str)
	return str
}
