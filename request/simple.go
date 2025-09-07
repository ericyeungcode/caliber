package request

import (
	"net/http"
	"sync/atomic"
	"time"
)

const (
	cliCount = 10
)

var httpClientList []*http.Client
var clientIndex = &atomic.Uint32{}

func init() {
	for i := 0; i < cliCount; i++ {
		httpClientList = append(httpClientList, NewHttpClientWithTimeout(5*time.Second))
	}
}

func GetNextHttpClient() *http.Client {
	return httpClientList[int(clientIndex.Add(1))%cliCount]
}

func Get[T any](url string) (T, error) {
	return HttpRequest[T](GetNextHttpClient(), http.MethodGet, url, nil, "")
}

func Post[T any](url string, headers map[string]string, jsonBodyStr string) (T, error) {
	return HttpRequest[T](GetNextHttpClient(), http.MethodPost, url, headers, jsonBodyStr)
}
