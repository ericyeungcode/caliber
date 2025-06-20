package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ericyeungcode/caliber"
)

type orderBookVo struct {
	Type      string      `json:"type,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Pair      string      `json:"pair"`
	Sequence  int64       `json:"sequence"`
	Bids      [][2]string `json:"bids"` // [[<price>, <qty>]]
	Asks      [][2]string `json:"asks"`
}

func main() {
	httpCli := &http.Client{Timeout: time.Second * 5}

	// return obj (incur clone)
	timeData, err := caliber.HttpRequestAndParse[caliber.ApiResponse[int64]](httpCli, http.MethodGet,
		"https://api.bit.com/spot/v1/system/time", nil, "")
	log.Printf("timeData: %+v, err:%v\n", timeData, err)

	// return pointer
	orderbook, err := caliber.HttpRequestAndParsePtr[caliber.ApiResponse[orderBookVo]](httpCli, http.MethodGet,
		"https://api.bit.com/spot/v1/orderbooks?pair=BTC-USDT", nil, "")
	log.Printf("orderbook: %+v, err:%v\n", orderbook, err)
}
