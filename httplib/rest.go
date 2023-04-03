package httplib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	FetchGetter interface {
		GetFetcher()
	}

	Fetcher interface {
		Get(param map[string]interface{}) Fetcher
		Result(data interface{}) (int, error)
		Post(data interface{}) Fetcher
		Error() error
		Response() *http.Response
	}
)

// http客户端
type RestClient struct {
	Address string // 服务器域名
	Port    int    // 服务器端口
	Prefix  string // 路径前缀如，可为空: /api, /live

	urlPrefix string
}

// http请求器
type RestFetcher struct {
	URL      string
	Request  *http.Request
	response *http.Response
	Err      error
}

// 生成RestClient
// @param host: http://www.example.com  http://www.example.com:8080
// @param prefix: 统一前缀 可为空
func NewRestClient(host string, prefix string) (*RestClient, error) {
	var port int
	var address string
	var err error

	// 校验host是否合法
	server := strings.Split(host, ":")
	switch len(server) {
	case 2: // host = http://www.example.com
		address = host
		port = 80
	case 3: // host = http://www.example.com:8080
		address = fmt.Sprintf("%s:%s", server[0], server[1])
		port, err = strconv.Atoi(server[2])
		if err != nil {
			return nil, fmt.Errorf("port error: %s", server[1])
		}
	default:
		return nil, fmt.Errorf("host invalid: %s", host)
	}

	// 创建client
	if _, err := url.Parse(address); err != nil {
		return nil, err
	}
	client := &RestClient{
		Address: address,
		Port:    port,
		Prefix:  prefix,
	}
	client.SetURLPrefix()
	return client, nil
}

// 拼接地址、端口、业务前缀
// 如果c.Port是80则不进行拼接
// 如果c.Prefix为空则不进行拼接
func (c *RestClient) SetURLPrefix() {
	c.urlPrefix = c.Address

	if c.Port != 80 {
		c.urlPrefix = fmt.Sprintf("%s:%d", c.urlPrefix, c.Port)
	}
	if strings.TrimSpace(c.Prefix) != "" {
		c.urlPrefix = fmt.Sprintf("%s/%s", c.urlPrefix, c.Prefix)
	}
}

// 设置request为json请求
// @param path: 路径
// @param request: 可为nil
func (c *RestClient) GetFetcher(path string, request *http.Request) *RestFetcher {
	if request == nil {
		request = &http.Request{
			Header: http.Header{},
		}
	}
	request.Header.Set(ContentType, ApplicationJson)

	return &RestFetcher{
		URL:     c.completedURL(path),
		Request: request,
	}
}

func (c *RestClient) completedURL(path string) string {
	return fmt.Sprintf("%s/%s", c.urlPrefix, path)
}

var defClient = &http.Client{}

// http get请求
// @param param: url参数，以?key=val形式拼接到url
func (fetcher *RestFetcher) Get(param map[string]interface{}) Fetcher {
	var params []string
	var err error

	fetcher.Request.Method = http.MethodGet

	// 处理url参数
	for paramName, paramVal := range param {
		params = append(params, fmt.Sprintf("%s=%v", paramName, paramVal))
	}
	if len(params) > 0 {
		fetcher.URL = fmt.Sprintf("%s?%s", fetcher.URL, strings.Join(params, "&"))
	}
	if fetcher.Request.URL, err = url.ParseRequestURI(fetcher.URL); err != nil {
		fetcher.Err = err
	}
	fetcher.response, fetcher.Err = defClient.Do(fetcher.Request)
	return fetcher
}

// http post请求
// @param data:
func (fetcher *RestFetcher) Post(data interface{}) Fetcher {
	var err error

	fetcher.Request.Method = http.MethodPost
	// 处理请求body
	body, err := json.Marshal(data)
	if err != nil {
		fetcher.Err = err
		return fetcher
	}
	bodyBuffer := bytes.NewReader(body)
	fetcher.Request.Body = ioutil.NopCloser(bodyBuffer)
	fetcher.Request.ContentLength = int64(bodyBuffer.Len())
	fetcher.Request.GetBody = func() (io.ReadCloser, error) {
		r := *bodyBuffer
		return ioutil.NopCloser(&r), nil
	}

	//发起请求
	if fetcher.Request.URL, err = url.ParseRequestURI(fetcher.URL); err != nil {
		fetcher.Err = err
	}
	fetcher.response, fetcher.Err = defClient.Do(fetcher.Request)
	return fetcher
}

// 对请求信息处理
// @param data: response.body unmarshal to data
// @return int: http状态码
// @return error:
func (fetcher RestFetcher) Result(data interface{}) (int, error) {
	body, err := ioutil.ReadAll(fetcher.response.Body)
	if err != nil {
		return fetcher.response.StatusCode, err
	}

	if err = json.Unmarshal(body, data); err != nil {
		return fetcher.response.StatusCode, err
	}
	return fetcher.response.StatusCode, nil
}

func (fetcher RestFetcher) Error() error {
	return fetcher.Err
}

func (fetcher RestFetcher) Response() *http.Response {
	return fetcher.response
}
