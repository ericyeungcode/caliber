package caliber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RestResp struct {
	StatusCode  int
	RespHeader  *http.Header
	RespContent []byte
}

func RestRespToString(rsp *RestResp) string {
	if rsp == nil {
		return "null"
	}
	return fmt.Sprintf("[status_code=%v, content=%v]", rsp.StatusCode, string(rsp.RespContent))
}

func NewHttpClientWithTimeout(dur time.Duration) *http.Client {
	return &http.Client{
		Timeout: dur,
	}
}

func DoHttp(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string) (*RestResp, error) {
	var body io.Reader = nil

	if jsonBodyStr != "" {
		body = bytes.NewBuffer([]byte(jsonBodyStr))
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &RestResp{
		StatusCode:  resp.StatusCode,
		RespHeader:  &resp.Header,
		RespContent: buf,
	}, nil
}

/*
DoHttpData send http request and parse whole response into value `data`
*/
func DoHttpData(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string, data interface{}) error {
	buRsp, err := DoHttp(client, method, url, headers, jsonBodyStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buRsp.RespContent, &data)
	if err != nil {
		return fmt.Errorf("DoHttpData fail to unmarshal data %v, err:%+v", string(buRsp.RespContent), err.Error())
	}
	return nil
}
