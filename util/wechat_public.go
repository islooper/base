package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type SimpleWeChatPublicClient struct {
	AccessToken string
}

func NewSimpleWeChatPublicClient(accessToken string) *SimpleWeChatPublicClient {
	return &SimpleWeChatPublicClient{AccessToken: accessToken}
}

///////////////// rps {
type Rsp struct {
	ErrCode   int32  `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	MsgId     string `json:"msg_id"`
	MsgDataId string `json:"msg_data_id"`
}

//  文本
func (s *SimpleWeChatPublicClient) CreateTextTask(touser []string, content string) (*Rsp, error) {
	if len(touser) < 2 || len(touser) > 10000 || content == "" {
		return nil, errors.New("touser or content cant be empty")
	}

	body := struct {
		touser  []string
		msgtype string
		text    struct {
			content string
		}
	}{
		touser:  touser,
		msgtype: "text",
		text:    struct{ content string }{content: content},
	}

	// 转json
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("struct format fail")
	}

	rsp, err := s.sendBody(string(b))
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (s *SimpleWeChatPublicClient) sendBody(data string) (*Rsp, error) {
	if s.AccessToken == "" {
		return nil, errors.New("access_token is empty")
	}

	url := "https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=" + s.AccessToken

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	var rsp Rsp
	err = json.Unmarshal(result, &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
