package httplib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

var defJsonClient = &http.Client{}

type (
	HttpRequestor interface {
		SetHeader(key, val string)
		SetRequest(url string, body map[string]interface{}) error
		Send() (*http.Response, error)
		SetAuth(user, passwd string)
	}
)

// jsonRequest json请求
type jsonRequest struct {
	clint   *http.Client
	request *http.Request
}

// JsonPostRequest json post请求
type JsonPostRequest struct {
	jsonRequest
}

/* jsonRequest post & get 通用方法 */

func (httpRequestor *jsonRequest) SetHeader(key, val string) {
	httpRequestor.request.Header.Set(key, val)
}

func (httpRequestor *jsonRequest) SetAuth(user, passwd string) {
	httpRequestor.request.SetBasicAuth(user, passwd)
}

func (httpRequestor *jsonRequest) Send() (*http.Response, error) {

	fmt.Println(httpRequestor.request)

	return defJsonClient.Do(httpRequestor.request)
}

// setJsonContentType 设置http json context-type
func (httpRequestor *jsonRequest) setJsonContentType() {
	httpRequestor.request.Header.Set(ContentType, ApplicationJson)
}

/* JsonPostRequest */

func (httpRequestor *JsonPostRequest) SetRequest(url string, data map[string]interface{}) error {
	var body io.Reader
	var err error

	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%s, body: %v", err, data)
	}
	body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if httpRequestor.request, err = http.NewRequest("POST", url, body); err != nil {
		return fmt.Errorf("%s, url: %s, body: %v", err, url, body)
	}
	fmt.Printf("url: %s, body: %v\n", url, body)
	httpRequestor.setJsonContentType()
	return nil
}
