package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ericyeungcode/caliber/common"
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

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewHttpClientWithTimeout(dur time.Duration) *http.Client {
	return &http.Client{
		Timeout: dur,
	}
}

func HttpRequestRaw(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string) (*Response, error) {
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

func HttpRequest[T any](client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string) (T, error) {
	rawRsp, err := HttpRequestRaw(client, method, url, headers, jsonBodyStr)
	if err != nil {
		var zero T
		return zero, err
	}

	return common.JsonToValue[T](string(rawRsp.Buffer))
}

func HttpGet[T any](client *http.Client, url string) (T, error) {
	return HttpRequest[T](client, http.MethodGet, url, nil, "")
}

func HttpPost[T any](client *http.Client, url string, headers map[string]string, jsonBodyStr string) (T, error) {
	return HttpRequest[T](client, http.MethodPost, url, headers, jsonBodyStr)
}
