package pinduoduo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	params3 "gitee.com/dtapps/go-library/utils/params"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const api = "https://gw-api.pinduoduo.com/api/router"

// Parameter 参数
type Parameter map[string]interface{}

// ParameterEncode 参数
type ParameterEncode []string

// PinDuoDuo 公共请求参数
type PinDuoDuo struct {
	ClientId     string //必填	POP分配给应用的client_id
	ClientSecret string //必填	POP分配给应用的client_secret
}

func (d *PinDuoDuo) Send(method string, param Parameter) (*simplejson.Json, error) {
	// 处理数据
	param["type"] = method
	param.setRequestData(d)
	// 请求
	resp, err := http.Post(api, "application/x-www-form-urlencoded;charset=utf-8", strings.NewReader(param.getRequestData()))
	if err != nil {
		return nil, err
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

// 签名数据
func sign(params Parameter, clientSecret string) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接参数
	query := bytes.NewBufferString(clientSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(params3.GetParamsString(params[k]))
	}
	query.WriteString(clientSecret)
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// 设置请求数据
func (p Parameter) setRequestData(d *PinDuoDuo) {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["client_id"] = d.ClientId
	p["data_type"] = "JSON"
	p["version"] = "v1"
	// 设置签名
	p["sign"] = sign(p, d.ClientSecret)
}

// 获取请求数据
func (p Parameter) getRequestData() string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range p {
		args.Set(key, params3.GetParamsString(val))
	}
	return args.Encode()
}
