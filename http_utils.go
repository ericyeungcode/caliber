package caliber

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func NewHttpClientWithTimeout(dur time.Duration) *http.Client {
	return &http.Client{
		Timeout: dur,
	}
}

func DoHttp(client *http.Client, method string, url string, headers map[string]string, body io.Reader) ([]byte, error) {
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

	return buf, nil
}

func DoHttpGet(client *http.Client, url string, headers map[string]string) ([]byte, error) {
	return DoHttp(client, http.MethodGet, url, headers, nil)
}
