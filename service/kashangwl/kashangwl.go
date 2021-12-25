package kashangwl

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/dtapps/go-library/utils/goparams"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const api = "http://www.kashangwl.com/api/"

// Parameter 参数
type Parameter map[string]interface{}

// ParameterEncode 参数
type ParameterEncode []string

// KaShangWl 每次请求需传入以下参数
type KaShangWl struct {
	CustomerId  int    // 商家编号
	CustomerKey string // 商家密钥
}

// Send 发送 http://cha.kashangwl.com/api-doc/
func (w *KaShangWl) Send(url string, param Parameter) (*simplejson.Json, error) {
	// 处理数据
	param.setRequestData(w)
	// 请求
	resp, e := http.Post(api+url, "application/json", strings.NewReader(param.getRequestData()))
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	// 返回结果
	body, _ := ioutil.ReadAll(resp.Body)
	respJson, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	return respJson, nil
}

// md5(key + 参数1名称 + 参数1值 + 参数2名称 + 参数2值...) 加密源串应为{key}customer_id1192442order_id827669582783timestamp1626845767
func sign(params Parameter, customerKey string) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	query := bytes.NewBufferString(customerKey)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(goparams.GetParamsString(params[k]))
	}
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

// 设置请求数据
func (p Parameter) setRequestData(w *KaShangWl) {
	// 	当前时间戳（单位：秒）
	timestamp := time.Now().UnixNano() / 1e6
	p["timestamp"] = timestamp
	p["customer_id"] = w.CustomerId
	// 设置签名
	p["sign"] = sign(p, w.CustomerKey)
}

// 获取请求数据
func (p Parameter) getRequestData() string {
	j, _ := json.Marshal(p)
	return string(j)
}
