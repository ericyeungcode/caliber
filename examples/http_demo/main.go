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
}
