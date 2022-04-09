package topsdk

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
	"topsdk/util"
)

type TopClient struct {
	AppKey         string
	AppSecret      string
	ServerUrl      string
	Format         string
	SignMethod     string
	ConnectTimeout int64
	ReadTimeout    int64
	Version        string
	Simplify       bool
}

func NewDefaultTopClient(AppKey string, AppSecret string, ServerUrl string, connectTimeount int64, readTimeout int64) TopClient {
	return TopClient{
		AppKey:         AppKey,
		AppSecret:      AppSecret,
		ServerUrl:      ServerUrl,
		Format:         "json",
		SignMethod:     "hmac-sha256",
		ConnectTimeout: connectTimeount,
		ReadTimeout:    readTimeout,
		Version:        "2.0",
		Simplify:       true,
	}
}

func (client *TopClient) ExecuteWithSession(method string, data map[string]interface{}, fileData map[string]interface{}, session string) (string, error) {
	var publicParam = make(map[string]interface{})
	publicParam["method"] = method
	publicParam["app_key"] = client.AppKey
	publicParam["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	publicParam["v"] = client.Version
	publicParam["sign_method"] = client.SignMethod
	publicParam["format"] = client.Format
	publicParam["simplify"] = client.Simplify
	publicParam["partner_id"] = "new_go_sdk"
	if session != "" {
		publicParam["session"] = session
	}
	sign := util.GetSign(publicParam, data, client.AppSecret)
	// 构建url
	serverUrl, _ := url.Parse(client.ServerUrl)
	urlValues := url.Values{}
	urlValues.Add("sign", sign)
	for k, v := range publicParam {
		urlValues.Add(k, fmt.Sprint(v))
	}
	serverUrl.RawQuery = urlValues.Encode()
	urlPath := serverUrl.String()
	// 构建body
	if fileData != nil && len(fileData) > 0 {
		return doPostWithFile(urlPath, data, fileData, client.ConnectTimeout)
	} else {
		return doPost(urlPath, data, client.ConnectTimeout)
	}

}

func doPost(urlPath string, data map[string]interface{}, timeout int64) (string, error) {
	bodyParam := url.Values{}
	for k, v := range data {
		bodyParam.Add(k, fmt.Sprint(v))
	}
	httpClient := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	resp, err := httpClient.Post(urlPath, "application/x-www-form-urlencoded", strings.NewReader(bodyParam.Encode()))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal("http.PostForm error", err)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll", err)
		return "", err
	}
	return string(body), nil
}

func doPostWithFile(urlPath string, data map[string]interface{}, fileData map[string]interface{}, timeout int64) (string, error) {
	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)
	for k, v := range data {
		err := writer.WriteField(k, fmt.Sprint(v))
		if err != nil {
			return "", err
		}
	}
	for k, v := range fileData {
		value, ok := v.([]byte)
		if ok {
			fileWriter, err := writer.CreateFormFile(k, "file")
			if err != nil {
				return "", err
			}
			_, err = io.Copy(fileWriter, bytes.NewReader(value))
			if err != nil {
				return "", err
			}
		}
	}

	err := writer.Close()
	if err != nil {
		return "", err
	}

	httpClient := http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	resp, err := httpClient.Post(urlPath, writer.FormDataContentType(), bodyBuf)
	if err != nil {
		log.Fatal("http.PostForm error", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll", err)
		return "", err
	}
	return string(body), nil
}

func (client *TopClient) Execute(method string, data map[string]interface{}, fileData map[string]interface{}) (string, error) {
	return client.ExecuteWithSession(method, data, fileData, "")
}
