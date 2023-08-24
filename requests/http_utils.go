package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int
	Header     *http.Header
	Buffer     []byte
}

func ResponseToString(rsp *Response) string {
	if rsp == nil {
		return "null"
	}
	return fmt.Sprintf("[status_code=%v, buffer=%v]", rsp.StatusCode, string(rsp.Buffer))
}

func NewHttpClientWithTimeout(dur time.Duration) *http.Client {
	return &http.Client{
		Timeout: dur,
	}
}

func DoHttp(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string) (*Response, error) {
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

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Header:     &resp.Header,
		Buffer:     buf,
	}, nil
}

/*
DoHttpData send http request and parse whole response into value `data`
*/
func DoHttpData(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string, data any) error {
	buRsp, err := DoHttp(client, method, url, headers, jsonBodyStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buRsp.Buffer, &data)
	if err != nil {
		return fmt.Errorf("DoHttpData fail to unmarshal data %v, err:%+v", string(buRsp.Buffer), err.Error())
	}
	return nil
}

/*
DoHttpEx send http request and parse payload (extract `RawResponse.Data`)
Useful to parse business object inside response
*/
func DoHttpPayload(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string, output interface{}) error {
	buRsp, err := DoHttp(client, method, url, headers, jsonBodyStr)
	if err != nil {
		return err
	}
	return ParseCommonPayload(buRsp.Buffer, &output)
}
