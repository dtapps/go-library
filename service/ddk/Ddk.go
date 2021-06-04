package ddk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/dtapps/go-library/service"
	"github.com/nilorg/sdk/convert"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	// ClientId 应用Key
	ClientId string
	// ClientSecret 秘密
	ClientSecret string
	// Router 环境请求地址
	Router = "https://gw-api.pinduoduo.com/api/router"
	// Timeout ...
	Timeout time.Duration
)

// Parameter 参数
type Parameter map[string]interface{}

// ParameterJsonEncode 参数
type ParameterJsonEncode []string

// copyParameter 复制参数
func copyParameter(srcParams Parameter) Parameter {
	newParams := make(Parameter)
	for key, value := range srcParams {
		newParams[key] = value
	}
	return newParams
}

// execute 执行API接口
func execute(param Parameter) (bytes []byte, err error) {
	err = checkConfig()
	if err != nil {
		return
	}

	var req *http.Request
	req, err = http.NewRequest("POST", Router, strings.NewReader(param.getRequestData()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := &http.Client{}
	httpClient.Timeout = Timeout
	var response *http.Response
	response, err = httpClient.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}

// Execute 执行API接口
func Execute(method string, param Parameter) (res *simplejson.Json, err error) {
	param["type"] = method
	param.setRequestData()

	var bodyBytes []byte
	bodyBytes, err = execute(param)
	if err != nil {
		return
	}

	return bytesToResult(bodyBytes)
}

func bytesToResult(bytes []byte) (res *simplejson.Json, err error) {
	res, err = simplejson.NewJson(bytes)
	if err != nil {
		return
	}

	if responseError, ok := res.CheckGet("error_response"); ok {
		if subMsg, subOk := responseError.CheckGet("sub_msg"); subOk {
			err = errors.New(subMsg.MustString())
		} else {
			err = errors.New(responseError.Get("msg").MustString())
		}
		res = nil
	}
	return
}

// 检查配置
func checkConfig() error {
	if ClientId == "" {
		return errors.New("ClientId 不能为空")
	}
	if ClientSecret == "" {
		return errors.New("ClientSecret 不能为空")
	}
	if Router == "" {
		return errors.New("Router 不能为空")
	}
	return nil
}

func (p Parameter) setRequestData() {
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["client_id"] = ClientId
	p["data_type"] = "JSON"
	p["version"] = "v1"
	// 设置签名
	p["sign"] = getSign(p)
}

// 获取请求数据
func (p Parameter) getRequestData() string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range p {
		args.Set(key, interfaceToString(val))
	}
	return args.Encode()
}

// 获取签名
func getSign(params Parameter) string {
	// 获取Key
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := bytes.NewBufferString(ClientSecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(interfaceToString(params[k]))
	}
	query.WriteString(ClientSecret)
	// 使用MD5加密
	h := md5.New()
	io.Copy(h, query)
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func interfaceToString(src interface{}) string {
	if src == nil {
		panic(service.ErrTypeIsNil)
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return convert.ToString(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}
