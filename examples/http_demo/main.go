package main

import (
	"log"

	"github.com/ericyeungcode/caliber/request"
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
	timeData, err := request.Get[*request.ApiResponse[int64]]("https://api.bit.com/spot/v1/system/time")
	log.Printf("timeData: %+v, err:%v\n", timeData, err)

	orderbook, err := request.Get[*request.ApiResponse[orderBookVo]]("https://api.bit.com/spot/v1/orderbooks?pair=BTC-USDT")
	log.Printf("orderbook: %+v, err:%v\n", orderbook, err)

	// Failed CASE:  this will produce error because orderbookStr is not json of string type
	orderbookStr, err := request.Get[string]("https://api.bit.com/spot/v1/orderbooks?pair=BTC-USDT")
	log.Printf("orderbook str1: %+v, err:%v\n", orderbookStr, err)
	if err == nil {
		panic("err should not be nil for GET[string]")
	}

	// make request and get raw string response (OK)
	obRsp, err := request.HttpRequestRaw(request.GetNextHttpClient(), "GET", "https://api.bit.com/spot/v1/orderbooks?pair=BTC-USDT", nil, "")
	log.Printf("orderbook str2: %+v, err:%v\n", request.ResponseToString(obRsp), err)
}
